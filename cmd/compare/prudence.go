// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/jwriteclub/prudence"
)

// checkGrace is how long past the deadline a Check that ignores its context is waited for.
const checkGrace = 30 * time.Second

// recordInfo (-info) keeps prudence's non-error diagnostics in Norm.Info.
var recordInfo bool

// no3d (-no-3d) turns normalisation rule 3d off: nested-module errors stay where rule 3c
// puts them.
var no3d bool

// abandonedChecks counts Check calls still running after deadline plus grace.
var abandonedChecks atomic.Int64

// RunPrudence checks one case with prudence.Check, the library call, nothing else.
// Support modules are read from the case directory (links into libs/ included), which
// is where SANY finds them; standard modules come from prudence itself.
//
// Rendering into Norm:
//   - an error diagnostic from the lexical or syntax phase: class parse-error
//   - an error diagnostic from the semantic phase: class semantic-error
//   - the "***Parse Error***" marker that prudence keeps at the start of SANY's parse
//     message is dropped (the SANY side drops it too); the rest of the message stays
//   - an error whose Code is "semantic/4220" (SANY's cannot-find-source-file) or
//     "semantic/4222" (circular dependency): class missing-module / module-cycle
//   - an error whose Code is "module/...not-found" or "module/...cycle": class
//     missing-module or module-cycle (SANY's classes for the same situations)
//   - no error and Completion == complete: class ok
//   - no error and Completion == incomplete: class incomplete (prudence says it has
//     not decided; this is not an answer and is counted apart from match/mismatch)
//
// Positions (normalisation rule 3): prudence's Column is per byte and its ByteOffset is
// authoritative, so the position is put in SANY's terms by sanyPosition below.
//
// visits is the case's Options.MaxAcquisitionVisits (0: prudence's default).
func RunPrudence(caseDir, root, stage string, timeout time.Duration, visits int) (n Norm) {
	defer func() {
		if r := recover(); r != nil {
			n = Norm{Class: ClassCrash, Note: fmt.Sprintf("prudence panic: %v", r)}
		}
	}()
	text, err := os.ReadFile(filepath.Join(caseDir, root))
	if err != nil {
		return Norm{Class: ClassCrash, Note: err.Error()}
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	resolve := func(ctx context.Context, name string) (prudence.Source, error) {
		data, err := os.ReadFile(filepath.Join(caseDir, name+".tla"))
		if errors.Is(err, os.ErrNotExist) {
			return prudence.Source{}, prudence.ErrModuleNotFound
		}
		if err != nil {
			return prudence.Source{}, err
		}
		return prudence.Source{Path: name + ".tla", Text: data, Origin: prudence.OriginWorkspace}, nil
	}
	req := prudence.Request{
		Stage:   prudence.AnalysisLevel(stage), // "" is prudence's default: semantic and level
		Origin:  prudence.OriginWorkspace,
		Path:    root,
		Text:    text,
		Resolve: resolve,
	}
	setAcquisitionVisits(&req, visits)
	// Check runs in its own goroutine so that a Check that ignores its context cannot
	// hold this worker: after the deadline plus a grace period the case is a timeout and
	// the worker moves on (the goroutine is abandoned and counted in abandonedChecks).
	type answer struct {
		res   prudence.Result
		err   error
		panic any
	}
	done := make(chan answer, 1)
	go func() {
		var a answer
		defer func() {
			if r := recover(); r != nil {
				a.panic = r
			}
			done <- a
		}()
		a.res, a.err = prudence.Check(ctx, req)
	}()
	var a answer
	select {
	case a = <-done:
	case <-time.After(timeout + checkGrace):
		abandonedChecks.Add(1)
		return Norm{Class: ClassTimeout, Note: fmt.Sprintf("no answer %s after the %s deadline; Check abandoned", checkGrace, timeout)}
	}
	if a.panic != nil {
		return Norm{Class: ClassCrash, Note: fmt.Sprintf("prudence panic: %v", a.panic)}
	}
	res, err := a.res, a.err
	if ctx.Err() != nil {
		// The deadline passed: whatever Check returned is not an answer.
		note := fmt.Sprintf("deadline %s passed", timeout)
		if err != nil {
			note += "; Check returned: " + err.Error()
		}
		return Norm{Class: ClassTimeout, Note: note}
	}
	if err != nil {
		return Norm{Class: ClassCrash, Note: "Check returned error: " + err.Error()}
	}
	comp := map[prudence.Completion]string{prudence.CompletionComplete: "complete", prudence.CompletionIncomplete: "incomplete"}[res.Completion]
	if comp == "" {
		comp = "unknown"
	}
	var info []InfoDiag
	defer func() { n.Completion, n.Level, n.Info = comp, string(res.Level), info }()
	var diags []Diag
	var offs []int // per diag: root-file byte offset, or -1 (not a root-file semantic diagnostic)
	class := ""
	for _, d := range res.Diagnostics {
		if d.Severity != prudence.SeverityError {
			if recordInfo {
				line := 0
				if d.Span != (prudence.Span{}) {
					line = d.Span.Start.Line + 1
				}
				info = append(info, InfoDiag{Code: d.Code, Phase: string(d.Phase), Severity: severityName(d.Severity), Message: d.Message, Path: d.Path, Line: line})
			}
			continue
		}
		switch d.Code {
		case "semantic/4220":
			class = ClassMissing
		case "semantic/4222":
			class = ClassCycle
		}
		if strings.HasPrefix(d.Code, "module/") {
			// Source diagnostics for module acquisition. They carry whatever location
			// prudence gives; Compare ignores positions where SANY has none.
			switch {
			case strings.Contains(d.Code, "cycl"):
				class = ClassCycle
			case strings.Contains(d.Code, "not-found"):
				class = ClassMissing
			}
		}
		u := unitOf(d)
		if u != nil && !u.valid() {
			return Norm{Class: ClassCrash, Note: fmt.Sprintf("invalid Diagnostic.Unit %+v on %s %q", *u, d.Code, d.Message)}
		}
		dg, off := projectDiag(caseDir, root, text, d, u)
		diags = append(diags, dg)
		offs = append(offs, off)
		if class == ClassCycle || class == ClassMissing {
			continue
		}
		if d.Phase == prudence.LevelLexical || d.Phase == prudence.LevelSyntax {
			class = ClassParse
		} else if class == "" {
			class = ClassSemantic
		}
	}
	if class != "" {
		note := ""
		if class == ClassSemantic && !no3d {
			note = splitNestedUnits(text, diags, offs)
		}
		sortDiags(diags)
		return Norm{Class: class, Diags: diags, Note: note}
	}
	if res.Completion == prudence.CompletionComplete {
		return Norm{Class: ClassOK}
	}
	return Norm{Class: ClassIncomp, Note: fmt.Sprintf("level %s", res.Level)}
}

// reportedUnit is prudence's Diagnostic.Unit as the comparison reads it: the module
// that reported the diagnostic and SANY's 1-based start line and column (UTF-16
// units, tab to the next multiple of 8) counted in that module's extracted text.
type reportedUnit struct {
	Module    string
	Line, Col int
}

// valid: a reported unit names a module and a real start. Anything else breaks the
// API's contract, and the case is scored as a prudence error, never compared.
func (u reportedUnit) valid() bool { return u.Module != "" && u.Line >= 1 && u.Col >= 1 }

// projectDiag puts one error diagnostic in SANY's terms. It returns the root-file
// byte offset rule 3d works on, or -1.
//
// Rule 3e: when prudence reports the unit (u != nil; Diagnostic.Unit), its module,
// line and column are SANY's own coordinates for that unit, and they are used as they
// are. The root mapping (rules 3a-3c) is not applied and rule 3d never sees the error
// (offset -1). Path and Span are not consulted. Otherwise rules 3a-3c apply as before.
func projectDiag(caseDir, root string, text []byte, d prudence.Diagnostic, u *reportedUnit) (Diag, int) {
	msg := Collapse(strings.TrimPrefix(d.Message, "***Parse Error***"))
	if u != nil {
		return Diag{Module: u.Module, Line: u.Line, Col: u.Col, Msg: msg, Unit: true}, -1
	}
	line, col := sanyPosition(sourceOf(caseDir, root, text, d.Path), d)
	mod := ""
	off := -1
	if !(d.Phase == prudence.LevelLexical || d.Phase == prudence.LevelSyntax) {
		mod, line = moduleOf(caseDir, root, text, d.Path, line)
		if d.Path == "" || d.Path == root {
			off = d.Span.Start.ByteOffset
		}
	}
	return Diag{Module: mod, Line: line, Col: col, Msg: msg}, off
}

var (
	rePosEncountered = regexp.MustCompile(`Encountered "(?s:.*?)" at line (\d+), column (\d+)`)
	rePosAny         = regexp.MustCompile(`line (\d+), col(?:umn)? (\d+)`)
)

// sanyPosition returns a diagnostic's start as SANY would print it.
//
// (a) For a lexical or syntax error, if the message already holds a SANY-style position ("Encountered ... at line L,
// column C", else the first "line L, col C"), that is the position.
// (b) Otherwise the line is prudence's line, and the column is rebuilt from the source
// line and ByteOffset with SANY's own rule. JavaCC's SimpleCharStream (tabSize = 8)
// counts one column per UTF-16 unit and moves a tab to the next multiple of eight:
// after a character the column value is v+1; after a tab, v is rounded up to the next
// multiple of 8. A token's column is the value before it, plus one.
func sanyPosition(text []byte, d prudence.Diagnostic) (int, int) {
	syntax := d.Phase == prudence.LevelLexical || d.Phase == prudence.LevelSyntax
	for _, re := range []*regexp.Regexp{rePosEncountered, reFairness, rePosAny} {
		if !syntax {
			break // a semantic message names other places ("conflicts with the one at line 31...")
		}
		if m := re.FindStringSubmatch(d.Message); m != nil {
			l, _ := strconv.Atoi(m[1])
			c, _ := strconv.Atoi(m[2])
			return l, c
		}
	}
	off := d.Span.Start.ByteOffset
	if off > len(text) {
		off = len(text)
	}
	start := off
	for start > 0 && text[start-1] != '\n' && text[start-1] != '\r' {
		start--
	}
	v := 0
	for _, r := range string(text[start:off]) {
		switch {
		case r == '\t':
			v += 8 - v%8
		case r > 0xFFFF:
			v += 2
		default:
			v++
		}
	}
	return d.Span.Start.Line + 1, v + 1
}

var (
	// reModHeader is a module header as SANY's lexer reads it on one line
	// (javacc/tla+.jj, pinned dc6a944): BEGIN_MODULE, `"----" ("-")* (" ")* "MODULE"`
	// (:948; spaces only, so the name may follow MODULE directly, as in
	// `---- MODULEOFt ----`), the SPEC-mode skip of spaces and tabs (:995), one
	// identifier-shaped token, the skip again, and SEPARATOR, `"----" ("-")*` (:1022).
	// moduleName decides whether the token is one BeginModule accepts (:1670-1684).
	reModHeader = regexp.MustCompile("-{4,} *MODULE[ \t]*([A-Za-z0-9_]+|@|\u2115|\u2124|\u211d)[ \t]*-{4,}")
	reModEnd    = regexp.MustCompile(`={4,}`) // SANY's end-of-module token is ====+ anywhere on a line
	reHasLetter = regexp.MustCompile(`[A-Za-z]`)
)

// moduleName reports whether BeginModule (tla+.jj:1670-1684) takes s as the module
// name: one token that is an IDENTIFIER or a keyword isFieldNameToken turns into one.
//   - IDENTIFIER (:1325) is CASE0 | CASE1 | CASE2 | CASE3 | CASE6 | CASEN | "@" |
//     NUMBER_SET (:961-975). Over [A-Za-z0-9_] these are exactly the words with a
//     letter, except that a word starting "WF_" or "SF_" is not one token: the WF and
//     SF tokens (:1055, :1084) match those three characters, longer than any
//     identifier there (CASE2 excludes F after W or S, CASE3 allows no _ after WF or
//     SF), so the name would be cut after them. A word with no letter is a number or
//     "_" followed by one.
//   - Every keyword spelled [A-Za-z0-9_]+ (ACTION ... WITH, :1027-1085, and ENABLED,
//     UNCHANGED, SUBSET, UNION, DOMAIN, :1150-1154) has a letter, and all of them are in
//     isFieldNameToken's ranges (:444-456), including MODULE, "WF_" and "SF_"; so is
//     "_" (US, :1090).
func moduleName(s string) bool {
	switch s {
	case "_", "WF_", "SF_", "@", "\u2115", "\u2124", "\u211d":
		return true
	}
	return reHasLetter.MatchString(s) && !strings.HasPrefix(s, "WF_") && !strings.HasPrefix(s, "SF_")
}

// moduleHeader returns the name a module header on line l opens, if there is one.
func moduleHeader(l string) (string, bool) {
	m := reModHeader.FindStringSubmatch(l)
	if m == nil || !moduleName(m[1]) {
		return "", false
	}
	return m[1], true
}

// splitLines splits as JavaCC counts lines: \n, \r\n and a lone \r each end a line.
func splitLines(text []byte) []string {
	s := strings.ReplaceAll(string(text), "\r\n", "\n")
	return strings.Split(strings.ReplaceAll(s, "\r", "\n"), "\n")
}

// codeOf cuts a line at a `\*` line comment: SANY's lexer skips the rest of the line
// there, so a ==== (or a module header) after it is not a token. Block comments are not
// tracked.
func codeOf(l string) string {
	if i := strings.Index(l, `\*`); i >= 0 {
		return l[:i]
	}
	return l
}

type modSpan struct {
	name       string
	start, end int // 1-based lines of the header and of the closing ====
	base       int // lines SANY numbers from: 0, or start-1 for a later top-level module
}

// fileModules lists the modules a file declares, with SANY's line base for each.
// Observed in SANY's output (e.g. mutants/C-examples-braf-...-bad-arity: an error at
// file line 451 in the second top-level module Common, which starts at line 386, is
// printed as "line 66 ... of module Common"): the first top-level module of a file is
// numbered from the file start, a later top-level module from its own header line, and
// a module nested in another is numbered like its top-level ancestor.
func fileModules(text []byte) []modSpan {
	var all []modSpan
	var open []int // indexes into all
	top := 0
	for i, l := range splitLines(text) {
		raw := strings.TrimRight(l, "\r")
		l = codeOf(raw)
		if len(open) == 0 {
			l = raw // between modules SANY skips arbitrary text up to the next header, comments included
		}
		ln := i + 1
		if hname, ok := moduleHeader(l); ok {
			base := 0
			name := hname
			if len(open) == 0 {
				if top > 0 {
					base = ln - 1
				}
				top++
			} else {
				base = all[open[0]].base
				name = all[open[0]].name // SANY names a nested module's errors for its top-level module
			}
			all = append(all, modSpan{name: name, start: ln, end: 1 << 30, base: base})
			open = append(open, len(all)-1)
		} else if reModEnd.MatchString(l) && len(open) > 0 {
			all[open[len(open)-1]].end = ln
			open = open[:len(open)-1]
		}
	}
	return all
}

// moduleOf names the module a semantic diagnostic is in, and puts its (1-based) line
// in SANY's numbering. path "" means the root file. If the file cannot be read or the
// line is in no module, the module is "" and the line is unchanged.
func moduleOf(caseDir, root string, rootText []byte, path string, line int) (string, int) {
	text := rootText
	if path != "" && path != root {
		b, err := os.ReadFile(filepath.Join(caseDir, path))
		if err != nil {
			return "", line
		}
		text = b
	}
	best := -1
	mods := fileModules(text)
	for i, m := range mods {
		if m.start <= line && line <= m.end && (best < 0 || m.start >= mods[best].start) {
			best = i
		}
	}
	if best < 0 {
		return "", line
	}
	return mods[best].name, line - mods[best].base
}

// sourceOf is the text of the file a diagnostic belongs to: the root's, or, for a
// diagnostic in a dependency (its Path names another file in the case directory), that
// file's. If the file cannot be read the root text is the fallback.
func sourceOf(caseDir, root string, rootText []byte, path string) []byte {
	if path == "" || path == root {
		return rootText
	}
	if b, err := os.ReadFile(filepath.Join(caseDir, path)); err == nil {
		return b
	}
	return rootText
}

func severityName(s prudence.Severity) string {
	switch s {
	case prudence.SeverityError:
		return "error"
	case prudence.SeverityWarning:
		return "warning"
	case prudence.SeverityInfo:
		return "info"
	case prudence.SeverityHint:
		return "hint"
	}
	return fmt.Sprintf("severity-%d", int(s))
}

// unit is one module header found in a file, with its real name.
type unit struct {
	name       string
	start, end int // 1-based header line and closing ==== line (end is 1<<30 if never closed)
	parent     int // index of the enclosing unit, or -1 for a top-level module
	depth      int
}

func scanUnits(text []byte) []unit {
	var all []unit
	var open []int
	for i, l := range splitLines(text) {
		raw := strings.TrimRight(l, "\r")
		l = codeOf(raw)
		if len(open) == 0 {
			l = raw // between modules SANY skips arbitrary text up to the next header, comments included
		}
		ln := i + 1
		if hname, ok := moduleHeader(l); ok {
			u := unit{name: hname, start: ln, end: 1 << 30, parent: -1}
			if len(open) > 0 {
				u.parent = open[len(open)-1]
				u.depth = all[u.parent].depth + 1
			}
			all = append(all, u)
			open = append(open, len(all)-1)
		} else if reModEnd.MatchString(l) && len(open) > 0 {
			all[open[len(open)-1]].end = ln
			open = open[:len(open)-1]
		}
	}
	return all
}

// splitNestedUnits is normalisation rule 3d. When a nested module N is reached by name
// (EXTENDS or INSTANCE) the monolith extractor makes it a unit of its own: SANY then
// prints an error inside N twice, once in the enclosing module (rule 3c numbering) and
// once "of module N" with lines counted from N's header. Prudence reports both at the
// same root-file bytes. Where that is unambiguous, the second of two identical
// diagnostics (same byte offset, same text) in N is renumbered into N's own unit:
// module N, line - (header line - 1). Unambiguous means: N is nested directly in a
// top-level module, N's name is declared once in the file, N and its parent are closed
// by ====, and exactly two identical diagnostics lie at that offset. Counts, columns,
// byte spans and module identity of every other diagnostic are untouched. Any
// diagnostic pair inside a nested module that does not fit (a duplicated or unterminated
// module, a deeper nesting, three or more copies) is left as it is, and the returned
// note names why, so the case stays a visible mismatch if the numbers differ.
func splitNestedUnits(text []byte, diags []Diag, offs []int) string {
	units := scanUnits(text)
	names := map[string]int{}
	for _, u := range units {
		names[u.name]++
	}
	lineOf := func(off int) int {
		if off > len(text) {
			off = len(text)
		}
		return 1 + len(splitLines(text[:off])) - 1
	}
	type key struct {
		off int
		msg string
	}
	groups := map[key][]int{}
	var order []key
	inUnit := map[key]int{}
	dupTop := map[string]bool{}
	for i, off := range offs {
		if off < 0 {
			continue
		}
		ln := lineOf(off)
		best := -1
		for j, u := range units {
			if u.start <= ln && ln <= u.end && (best < 0 || u.start >= units[best].start) {
				best = j
			}
		}
		if best >= 0 && units[best].parent < 0 && names[units[best].name] > 1 {
			dupTop[units[best].name] = true // a repeated top-level module: SANY's own choice among the copies is not known here
		}
		if best < 0 || units[best].parent < 0 {
			continue
		}
		k := key{off, diags[i].Msg}
		if _, ok := groups[k]; !ok {
			order = append(order, k)
			inUnit[k] = best
		}
		groups[k] = append(groups[k], i)
	}
	var reasons []string
	for n := range dupTop {
		reasons = append(reasons, "module "+n+" is declared more than once")
	}
	for _, k := range order {
		u := units[inUnit[k]]
		idx := groups[k]
		if len(idx) < 2 {
			continue
		}
		switch {
		case names[u.name] > 1:
			reasons = append(reasons, "module "+u.name+" is declared more than once")
		case u.end == 1<<30 || units[u.parent].end == 1<<30:
			reasons = append(reasons, "module "+u.name+" or its parent is not closed")
		case u.depth != 1:
			reasons = append(reasons, "module "+u.name+" is nested more than one level deep")
		case len(idx) > 2:
			reasons = append(reasons, fmt.Sprintf("%d identical diagnostics in module %s", len(idx), u.name))
		default:
			diags[idx[1]].Module = u.name
			diags[idx[1]].Line = lineOf(k.off) - (u.start - 1)
		}
	}
	if len(reasons) == 0 {
		return ""
	}
	seen := map[string]bool{}
	var out []string
	for _, r := range reasons {
		if !seen[r] {
			seen[r] = true
			out = append(out, r)
		}
	}
	return "nested-module attribution ambiguous (rule 3d): " + strings.Join(out, "; ")
}
