// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const tlatools = "tlaplus/tlatools/org.lamport.tlatools"

type ctx struct {
	ext     string
	commits map[string]string
	libs    *Libs
	claimed map[string]bool // external-relative paths already used by tier A
}

func (x *ctx) read(rel string) string {
	b, err := os.ReadFile(filepath.Join(x.ext, rel))
	if err != nil {
		fatal(err)
	}
	return string(b)
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, "extract:", err)
	os.Exit(1)
}

func slug(s string) string {
	s = strings.ToLower(s)
	var sb strings.Builder
	dash := false
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			sb.WriteRune(r)
			dash = false
		} else if !dash && sb.Len() > 0 {
			sb.WriteByte('-')
			dash = true
		}
	}
	out := strings.Trim(sb.String(), "-")
	if len(out) > 40 {
		out = strings.Trim(out[:40], "-")
	}
	return out
}

// snippetCase builds a single-file case from inline module text.
func snippetCase(text string) *Case {
	declared, _ := moduleInfo(text)
	c := &Case{}
	mod := "Test"
	if len(declared) > 0 {
		mod = declared[0]
		c.Root = mod
		if mod != "Test" {
			c.Flags = append(c.Flags, "module_name_not_Test")
		}
	} else {
		c.Flags = append(c.Flags, "no_module_header")
	}
	c.RootFile = mod + ".tla"
	c.content = map[string][]byte{c.RootFile: []byte(text)}
	return c
}

// synCorpusCases reads every .txt corpus file under dir (relative to ext).
func (x *ctx) synCorpusCases(dir, tier, source, idp string, knownFail map[string]bool, basis string, skipDirs ...string) []*Case {
	var files []string
	filepath.Walk(filepath.Join(x.ext, dir), func(p string, fi os.FileInfo, err error) error {
		if err != nil || fi.IsDir() || !strings.HasSuffix(p, ".txt") {
			return nil
		}
		for _, s := range skipDirs {
			if strings.Contains(p, "/"+s+"/") {
				return nil
			}
		}
		files = append(files, p)
		return nil
	})
	sort.Strings(files)
	var out []*Case
	for _, f := range files {
		rel, _ := filepath.Rel(filepath.Join(x.ext, dir), f)
		b, _ := os.ReadFile(f)
		tests := parseSynCorpus(string(b))
		for i, t := range tests {
			c := snippetCase(t.Input)
			c.ID = fmt.Sprintf("%s-%s-%03d", idp, slug(strings.TrimSuffix(rel, ".txt")), i+1)
			c.Tier, c.Source = tier, source
			c.Origin = dir + "/" + rel + " :: " + t.Name
			c.Name = t.Name
			if t.Expect != "" && !t.has("error") {
				c.aux = map[string][]byte{"expected.sexp": []byte(t.Expect)}
			}
			if basis != "" {
				e := &Expect{Stage: "syntax", Result: "accept", Basis: basis, Upstream: "ok"}
				if t.has("error") {
					e.Result = "reject"
				}
				if t.has("skip") {
					e.Upstream = "skipped"
				}
				if knownFail[strings.ToLower(t.Name)] {
					e.Upstream = "known_failure"
					e.Note = "listed in TlaPlusSyntaxCorpusTests expectFailures: SANY is expected to deviate from the corpus here"
				}
				c.Expect = e
			}
			if strings.Contains(rel, "unicode") {
				c.Flags = append(c.Flags, "unicode_syntax")
			}
			if strings.Contains(rel, "nonstandard") {
				c.Flags = append(c.Flags, "nonstandard")
			}
			out = append(out, c)
		}
	}
	return out
}

var reQuoted = regexp.MustCompile(`"((?:[^"\\]|\\.)*)"`)

// knownFailures extracts the expectFailures(...) names from TlaPlusSyntaxCorpusTests.java.
func (x *ctx) knownFailures() map[string]bool {
	src := x.read(tlatools + "/test/tla2sany/parser/TlaPlusSyntaxCorpusTests.java")
	i := strings.Index(src, "expectFailures(")
	j := strings.Index(src[i:], ");")
	m := map[string]bool{}
	for _, q := range reQuoted.FindAllStringSubmatch(src[i:i+j], -1) {
		m[strings.ToLower(q[1])] = true
	}
	return m
}

func (x *ctx) setCommit(c *Case) {
	c.Commit = x.commits[strings.SplitN(c.Source, "/", 2)[0]]
}

var reSemName = regexp.MustCompile(`^([WE])(\d+)(?:_(\S+))?_Test\.tla$`)
var reAssert = regexp.MustCompile(`\b(?:RefersTo|IsLevel)\(`)

func (x *ctx) tierA() []*Case {
	var out []*Case
	known := x.knownFailures()
	// A1: unified syntax corpus (SANY syntax stage)
	for _, c := range x.synCorpusCases(tlatools+"/test/tla2sany/corpus", "A", "tlaplus/syntax-corpus", "A-syn", known, "SyntaxCorpusRunner: tree-sitter corpus test, ':error' tag means SANY must reject, otherwise the expected S-expression must match") {
		out = append(out, c)
	}
	// A2: semantic corpus (accept; embedded RefersTo/IsLevel assertions)
	semDir := tlatools + "/test/tla2sany/semantic/corpus"
	ents, _ := os.ReadDir(filepath.Join(x.ext, semDir))
	for _, e := range ents {
		n := e.Name()
		if !strings.HasSuffix(n, ".tla") || n == "Semantics.tla" {
			continue
		}
		c, err := fileCase(x.libs, x.ext, filepath.Join(x.ext, semDir, n))
		if err != nil {
			fatal(err)
		}
		c.ID, c.Tier, c.Source = "A-sem-"+slug(strings.TrimSuffix(n, ".tla")), "A", "tlaplus/semantic-corpus"
		c.Name = n
		c.Expect = &Expect{Stage: "semantic", Result: "accept", Basis: "SemanticCorpusTests: no parse/semantic errors, no warnings, and every RefersTo/IsLevel assertion in the module holds",
			Upstream: "ok", Asserts: len(reAssert.FindAllString(string(c.content[n]), -1))}
		if n == "NegativeOpTest.tla" {
			c.Expect.Upstream = "skipped"
			c.Expect.Note = "Assume.assumeFalse in SemanticCorpusTests, tlaplus/tlaplus#1130"
		}
		x.claim(semDir + "/" + n)
		out = append(out, c)
	}
	// A3: semantic error corpus (expected code from file name)
	errDir := tlatools + "/test/tla2sany/semantic/error_corpus"
	ents, _ = os.ReadDir(filepath.Join(x.ext, errDir))
	for _, e := range ents {
		n := e.Name()
		m := reSemName.FindStringSubmatch(n)
		if m == nil {
			continue
		}
		c, err := fileCase(x.libs, x.ext, filepath.Join(x.ext, errDir, n))
		if err != nil {
			fatal(err)
		}
		code, _ := strconv.Atoi(m[2])
		c.ID, c.Tier, c.Source = "A-err-"+slug(strings.TrimSuffix(n, "_Test.tla")), "A", "tlaplus/semantic-error-corpus"
		c.Name = n
		e := &Expect{Stage: "semantic", Code: code, Severity: m[1], Upstream: "ok",
			Basis: "SemanticErrorCorpusTests: error code in file name; log.isFailure() must equal (severity is E) and the log must contain the code"}
		if m[1] == "E" {
			e.Result = "reject"
		} else {
			e.Result = "accept"
			e.Note = "warning code: SANY must succeed and report the warning"
		}
		c.Expect = e
		x.claim(errDir + "/" + n)
		out = append(out, c)
	}
	// A4: TestLevelChecking table
	out = append(out, x.levelCases()...)
	// A5: ParseErrorTests
	{
		c := snippetCase("---- MODULE Test ----\nx = 0\n====")
		c.ID, c.Tier, c.Source = "A-perr-001", "A", "tlaplus/inline-java"
		c.Origin = tlatools + "/test/tla2sany/parser/ParseErrorTests.java"
		c.Name = "x = 0"
		c.Expect = &Expect{Stage: "syntax", Result: "reject", Upstream: "ok", Basis: "ParseErrorTests: ParseException and output contains the message",
			Messages: []string{"Was expecting \"==== or more Module body\"\nEncountered \"x\" at line"}}
		out = append(out, c)
	}
	// A6: curated fixtures of test-model used by tla2sany tests
	tm := tlatools + "/test-model/"
	type fx struct {
		path, result, stage, basis, up, note string
		msgs                                 []string
	}
	fxs := []fx{
		{"sany/Github723.tla", "reject", "semantic", "Github723Test: SANYmain output contains message", "ok", "", []string{"An operator must be substituted for symbol 'C', and it must have arity 1."}},
		{"sany/RecursiveDefDeclMismatch.tla", "reject", "semantic", "RecursiveDefDeclMismatchTest: SANYmain output contains messages", "ok", "", []string{"Definition of CountDown has different arity than its RECURSIVE declaration.", "The operator CountDown requires 2 arguments."}},
		{"sany/NestedModuleTopLevelInstance.tla", "accept", "semantic", "NestedModuleInstanceTest: semantic errors isSuccess", "ok", "", nil},
		{"sany/NestedModuleLetInstance.tla", "accept", "semantic", "NestedModuleInstanceTest: semantic errors isSuccess", "ignored", "@Ignore upstream: SANY reports a spurious MODULE_REDEFINED, so SANY is expected to deviate", nil},
		{"sany/StringEscapes.tla", "accept", "semantic", "TestXMLExporterStringEscapes: XMLExporter exit code OK", "ok", "", nil},
		{"sany/StringFormFeedEscape.tla", "accept", "semantic", "TestXMLExporterStringEscapes: SANY parses it, XMLExporter must exit XML_UNREPRESENTABLE_CHARACTER and name U+000C", "ok", "", nil},
		{"Github429.tla", "accept", "semantic", "Github429Test: frontEndParse + frontEndSemanticAnalysis throw no exception", "ok", "weak: only 'no exception' is asserted, not the absence of reported errors", nil},
		{"DieHard.tla", "accept", "semantic", "TestXMLExporterModule: XMLExporter exit code 0 with empty stderr", "ok", "", nil},
		{"CaseOtherXml.tla", "accept", "semantic", "TestXMLExporterModule: XMLExporter exit code 0", "ok", "", nil},
		{"TLACommentStyles.tla", "accept", "semantic", "TestXMLExporterModule: XMLExporter exit code 0", "ok", "", nil},
		{"Echo/Relation.tla", "accept", "semantic", "TestXMLExporterModule: XMLExporter exit code 0", "ok", "", nil},
		{"Decimal.tla", "accept", "semantic", "TestDecimalXMLExport: parseSpec succeeds", "ok", "", nil},
	}
	for _, f := range fxs {
		c, err := fileCase(x.libs, x.ext, filepath.Join(x.ext, tm+f.path))
		if err != nil {
			fatal(err)
		}
		c.ID, c.Tier, c.Source = "A-fix-"+slug(strings.TrimSuffix(f.path, ".tla")), "A", "tlaplus/test-model-fixture"
		c.Name = f.path
		c.Expect = &Expect{Stage: f.stage, Result: f.result, Basis: f.basis, Upstream: f.up, Note: f.note, Messages: f.msgs}
		x.claim(tm + f.path)
		out = append(out, c)
	}
	for _, c := range out {
		x.setCommit(c)
		if c.Origin == "" {
			c.Origin = c.Name
		}
	}
	return out
}

func (x *ctx) claim(rel string) { x.claimed[rel] = true }

// javaString decodes a Java string literal body.
func javaString(s string) string {
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] != '\\' || i+1 >= len(s) {
			sb.WriteByte(s[i])
			continue
		}
		i++
		switch s[i] {
		case 'n':
			sb.WriteByte('\n')
		case 't':
			sb.WriteByte('\t')
		case '"':
			sb.WriteByte('"')
		case '\\':
			sb.WriteByte('\\')
		default:
			sb.WriteByte('\\')
			sb.WriteByte(s[i])
		}
	}
	return sb.String()
}

var reLevel = regexp.MustCompile(`(?m)^\s*new LevelCheckingTestCase\("((?:[^"\\]|\\.)*)",\s*(true|false)\)`)

func (x *ctx) levelCases() []*Case {
	path := tlatools + "/test/tla2sany/semantic/TestLevelChecking.java"
	src := x.read(path)
	var out []*Case
	for i, m := range reLevel.FindAllStringSubmatch(src, -1) {
		expr := javaString(m[1])
		c := snippetCase("---- MODULE Test ----\nCONSTANT c\nVARIABLE x\nop == " + expr + "\n====")
		c.ID, c.Tier, c.Source = fmt.Sprintf("A-lvl-%03d", i+1), "A", "tlaplus/inline-java"
		c.Origin, c.Name = path, expr
		res := "reject"
		if m[2] == "true" {
			res = "accept"
		}
		c.Expect = &Expect{Stage: "level", Result: res, Upstream: "ok",
			Basis: "TestLevelChecking: parse and semantics succeed, checkLevel result equals the table value"}
		out = append(out, c)
	}
	return out
}
