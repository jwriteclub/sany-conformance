// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Diag is one error, reduced to where it is and what it says.
type Diag struct {
	// Module is the module the error is in, when known ("" for parse errors and for
	// errors whose file has no module header prudence could read).
	Module string `json:"module,omitempty"`
	Line   int    `json:"line"`
	Col    int    `json:"col"`
	Msg    string `json:"msg"`
	// Unit is true when the position came from prudence's Diagnostic.Unit: the
	// reporting unit's own module, line and column, as SANY prints them "of module
	// N" for a monolith-extracted unit N (rule 3e). Prudence only; omitted when false,
	// so every other row is byte for byte what it was.
	Unit bool `json:"unit,omitempty"`
}

// Norm is one side's answer for one case, in the form that gets compared.
type Norm struct {
	Class string `json:"class"` // ok, parse-error, semantic-error, crash, incomplete
	Diags []Diag `json:"diags"`
	Note  string `json:"note,omitempty"`
	// Prudence only: its public Completion ("complete", "incomplete", "unknown") and Level.
	Completion string `json:"completion,omitempty"`
	Level      string `json:"level,omitempty"`
	// Prudence only, with -info: its non-error diagnostics (the analysis/incomplete
	// reason among them), as reported. Not used by Compare.
	Info []InfoDiag `json:"info,omitempty"`
}

// InfoDiag is one non-error prudence diagnostic, kept verbatim for diagnosis.
type InfoDiag struct {
	Code     string `json:"code"`
	Phase    string `json:"phase"`
	Severity string `json:"severity"`
	Message  string `json:"message"`
	Path     string `json:"path,omitempty"`
	Line     int    `json:"line"` // prudence's 1-based start line; 0 with an empty span
}

const (
	ClassOK       = "ok"
	ClassParse    = "parse-error"
	ClassSemantic = "semantic-error"
	ClassCrash    = "crash"
	ClassMissing  = "missing-module" // SANY could not find an imported module's file
	ClassCycle    = "module-cycle"   // SANY found a circular dependency among modules
	ClassIncomp   = "incomplete"     // prudence only: it stopped before it could say
	ClassTimeout  = "timeout"        // prudence only: the per-case deadline passed; never a pass
)

var (
	reEncountered = regexp.MustCompile(`Encountered "(?s:.*?)" at line (\d+), column (\d+)`)
	reLexical     = regexp.MustCompile(`Lexical error at line (\d+), column (\d+)`)
	reAnyPos      = regexp.MustCompile(`line (\d+), col(?:umn)? (\d+)`)
	reFairness    = regexp.MustCompile(`Error in fairness expression at (\d+): (\d+)`)
	reSemLoc      = regexp.MustCompile(`^(?:line (-?\d+), col (-?\d+) to line -?\d+, col -?\d+ of module (\S+)|(Unknown location)|In module (.+?))\s*$`)
	reSection     = regexp.MustCompile(`^\*\*\* (Errors|Warnings): \d+`)
	// Lines SANY prints between its error lists; each ends the entry being read.
	reSanyBreak = regexp.MustCompile(`^(Semantic processing of module |Semantic errors|Fatal errors|Parsing file |\*{6} SANY)`)
	reSpace     = regexp.MustCompile(`\s+`)
)

// Collapse is the only text normalisation: runs of whitespace become one space.
func Collapse(s string) string { return strings.TrimSpace(reSpace.ReplaceAllString(s, " ")) }

// NormalizeSany reads what plain `java -cp tla2tools.jar tla2sany.SANY` printed.
//
//   - parse error: one diagnostic per record SANY's parser logged (sanyParseDiags),
//     each at the position it names ("Encountered ... at line L, column C", "Lexical
//     error at line L, column C", ...) with its text, residual stack trace included.
//     An output that is not exactly a sequence of such records is kept as one
//     diagnostic, from after the "Parsing file" line to before "Fatal errors".
//   - circular dependency ("Circular dependency among .tla files"): class module-cycle,
//     no position, just that one sentence (the "*** Errors" tail repeats it).
//   - missing module ("Cannot find source file for module"): class missing-module,
//     no position, just that one sentence (the "*** Errors" tail repeats it).
//   - semantic error: one diagnostic per distinct entry under "*** Errors: N" (see
//     normalizeSemantic), in the module its location line names, at the entry's start
//     "line L, col C", with the entry's text. Entries under "*** Warnings" are dropped.
//   - otherwise ok. A crash is decided by the caller, from the exit status or stderr.
func NormalizeSany(stdout string) Norm {
	if strings.Contains(stdout, "Circular dependency among") {
		return withClass(normalizeSemantic(stdout), ClassCycle, stdout)
	}
	if strings.Contains(stdout, "***Parse Error***") || strings.Contains(stdout, "Lexical error") ||
		strings.Contains(stdout, "Lexical {error") || strings.Contains(stdout, "Error in fairness expression") {
		return Norm{Class: ClassParse, Diags: sanyParseDiags(stdout)}
	}
	if strings.Contains(stdout, "Cannot find source file for module") {
		return withClass(normalizeSemantic(stdout), ClassMissing, stdout)
	}
	return normalizeSemantic(stdout)
}

// withClass gives a module-acquisition failure (cycle, missing module) its class and every
// error SANY printed for it (e.g. 4003 "An unexpected problem ... validate" plus 4220/4222).
// If the output has no "*** Errors" list at all, the one sentence is kept as before.
func withClass(n Norm, class, stdout string) Norm {
	n.Class = class
	if len(n.Diags) == 0 {
		marker := "Cannot find source file"
		if class == ClassCycle {
			marker = "Circular dependency among"
		}
		s := stdout[strings.Index(stdout, marker):]
		if class == ClassCycle {
			s = untilErrors(s)
		} else if j := strings.Index(s, "\n"); j >= 0 {
			s = s[:j]
		}
		n.Diags = []Diag{{Msg: Collapse(s)}}
	}
	return n
}

// sentinel is the begin line SANY prints for a location it never set.
const sentinel = 2147483647

// normalizeSemantic reads SANY's semantic output. Its source (tla2sany/drivers/SANY.java,
// frontEndSemanticAnalysis) shares one Errors object across all modules, and after each
// module "Semantic processing of module X" it prints the whole accumulated list again
// ("Semantic errors:" and "*** Errors: N"), warnings first when there are any. Errors.addMessage
// is idempotent on (code, location, text), so the final list is the set of distinct
// entries. So: every "*** Errors" entry is read, wherever it is; an entry is identified by
// its location line (which names the module the error is in) plus its text; repeats of
// the same identity are dropped. "*** Warnings" entries are dropped. The abort path
// ("Semantic errors detected before the unexpected exception") prints Errors before
// Warnings; both are read the same way, section by section.
func normalizeSemantic(stdout string) Norm {
	var diags []Diag
	seen := map[string]bool{}
	section := ""
	sawErrors := false
	var cur *Diag
	curKey := ""
	flush := func() {
		if cur != nil {
			cur.Msg = Collapse(cur.Msg)
			if k := curKey + "\x00" + cur.Msg; !seen[k] {
				seen[k] = true
				diags = append(diags, *cur)
			}
			cur = nil
		}
	}
	for _, line := range strings.Split(stdout, "\n") {
		if m := reSection.FindStringSubmatch(line); m != nil {
			flush()
			section = m[1]
			if section == "Errors" {
				sawErrors = true
			}
			continue
		}
		if reSanyBreak.MatchString(line) {
			flush()
			section = ""
			continue
		}
		if section != "Errors" {
			continue
		}
		if m := reSemLoc.FindStringSubmatch(line); m != nil {
			flush()
			curKey = strings.TrimSpace(line)
			cur = &Diag{Module: m[3]}
			if m[5] != "" { // "In module M": a location with no coordinates (4003, BUILTINS)
				cur.Module = m[5]
			} else if m[4] == "" {
				cur.Line, _ = strconv.Atoi(m[1])
				cur.Col, _ = strconv.Atoi(m[2])
				if cur.Line == sentinel {
					cur.Line, cur.Col, cur.Module = 0, 0, "" // never set: no position
				}
			}
		} else if cur != nil {
			cur.Msg += " " + line
		}
	}
	flush()
	if sawErrors {
		sortDiags(diags)
		return Norm{Class: ClassSemantic, Diags: diags}
	}
	return Norm{Class: ClassOK}
}

// untilErrors cuts SANY's text before the "*** Errors: N" tail that repeats it.
func untilErrors(s string) string {
	if j := strings.Index(s, "\n*** Errors"); j >= 0 {
		return s[:j]
	}
	return s
}

// parseRecordStarts are the first lines of the records SANY's parser can log. In
// tla2sany/parser at dc6a944, TLAplusParser.parse (TLAplusParser.java:113-116) logs
// every entry of PErrors, in push order, with its own println (SimpleSanyOutput.log),
// and these are all the entries pushed:
//   - TLAplusParser.java:85, a ParseException: msgStackToString (:526) starts with
//     "***Parse Error***" on a line of its own;
//   - :93-96, a TokenMgrError at EOF across lines: "Lexical {error: EOF reached, ...";
//   - :97, any other TokenMgrError: its message, which JavaCC's LexicalError
//     (TokenMgrError.java:98) makes "Lexical error at line L, column C.  Encountered: ...";
//   - :5323, pushed without throwing: "@ used in !.@";
//   - :5536 and :5555, pushed without throwing: "Error in fairness expression at L: C",
//     the second with ": could not link arguments";
//   - OperatorStack.java:447, pushed without throwing: "Couldn't properly parse
//     expression-- incomplete expression at ...".
//
// Each pattern matches a whole line, so no record's own text can start another: the
// lines inside a ParseException record are "Was expecting", "Encountered", "Residual
// stack trace follows:" and "X starting at line L, column C."; the lexical messages
// escape their characters (TokenMgrError.addEscapes).
var parseRecordStarts = []*regexp.Regexp{
	regexp.MustCompile(`^\*\*\*Parse Error\*\*\*$`),
	regexp.MustCompile(`^Lexical \{error: EOF reached, possibly open comment starting around line -?\d+$`),
	regexp.MustCompile(`^Lexical error at line -?\d+, column -?\d+\.  Encountered: `),
	regexp.MustCompile(`^@ used in !\.@$`),
	regexp.MustCompile(`^Error in fairness expression at -?\d+: -?\d+(?:: could not link arguments)?$`),
	regexp.MustCompile(`^Couldn't properly parse expression-- incomplete expression at `),
}

// parseText is the part of SANY's parse output that holds the parser's records: the
// lines before "Fatal errors", without SANY's banner and "Parsing file" lines.
func parseText(out string) []string {
	var kept []string
	for _, l := range strings.Split(out, "\n") {
		if strings.HasPrefix(l, "Fatal errors") {
			break
		}
		if strings.HasPrefix(l, "Parsing file") || strings.HasPrefix(l, "****** SANY") {
			continue
		}
		kept = append(kept, l)
	}
	return kept
}

// parseRecords splits the parse text into SANY's records, one per parseRecordStarts
// line, each running to the next one. It returns false unless the text is exactly such
// a sequence: at least one record, and nothing but blank lines before the first.
func parseRecords(lines []string) ([][]string, bool) {
	var recs [][]string
	for _, l := range lines {
		start := false
		for _, re := range parseRecordStarts {
			if re.MatchString(l) {
				start = true
				break
			}
		}
		switch {
		case start:
			recs = append(recs, []string{l})
		case len(recs) > 0:
			recs[len(recs)-1] = append(recs[len(recs)-1], l)
		case strings.TrimSpace(l) != "":
			return nil, false
		}
	}
	return recs, len(recs) > 0
}

// parseDiag is one parse error: the first position its text names, in this order:
// "Encountered ... at line L, column C", "Lexical error at line L, column C", "Error in
// fairness expression at L: C", any "line L, column C"; none (0:0, rule 5) if it names
// none, as "@ used in !.@" does (SANY's ParseError(String) has no location). The text is
// the record's, minus the "***Parse Error***" line, whitespace collapsed.
func parseDiag(text string) Diag {
	d := Diag{}
	for _, re := range []*regexp.Regexp{reEncountered, reLexical, reFairness, reAnyPos} {
		if m := re.FindStringSubmatch(text); m != nil {
			d.Line, _ = strconv.Atoi(m[len(m)-2])
			d.Col, _ = strconv.Atoi(m[len(m)-1])
			break
		}
	}
	var kept []string
	for _, l := range strings.Split(text, "\n") {
		if !strings.HasPrefix(l, "***Parse Error***") {
			kept = append(kept, l)
		}
	}
	d.Msg = Collapse(strings.Join(kept, " "))
	return d
}

// sanyParseDiags is SANY's parse output as errors: one per record the parser logged
// (parseRecords), in rule-4 order. An output that does not split exactly into such
// records is one error, as before: the position is the first the whole output names,
// the text is all of parseText.
func sanyParseDiags(out string) []Diag {
	recs, ok := parseRecords(parseText(out))
	if !ok {
		d := parseDiag(out)
		d.Msg = parseDiag(strings.Join(parseText(out), "\n")).Msg
		return []Diag{d}
	}
	diags := make([]Diag, 0, len(recs))
	for _, r := range recs {
		diags = append(diags, parseDiag(strings.Join(r, "\n")))
	}
	sortDiags(diags)
	return diags
}

func sortDiags(d []Diag) {
	sort.Slice(d, func(i, j int) bool {
		if d[i].Line != d[j].Line {
			return d[i].Line < d[j].Line
		}
		if d[i].Col != d[j].Col {
			return d[i].Col < d[j].Col
		}
		if d[i].Msg != d[j].Msg {
			return d[i].Msg < d[j].Msg
		}
		return d[i].Module < d[j].Module
	})
}
