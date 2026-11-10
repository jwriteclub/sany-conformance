// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"regexp"
	"strings"
)

// SynTest is one test of a tree-sitter style corpus file (the format read by
// SyntaxCorpusFileParser.java: ===|||, name, ===|||, input, ---|||, S-expression).
type SynTest struct {
	Name   string
	Attrs  []string // from ":error" / ":skip" lines
	Input  string
	Expect string
}

var (
	reHeaderLine = regexp.MustCompile(`^===+\|\|\|\r?$`)
	reSepLine    = regexp.MustCompile(`^---+\|\|\|\r?$`)
)

// parseSynCorpus mirrors the regex logic of SyntaxCorpusFileParser.getCorpusTests
// with a line scanner (Go's RE2 has no look-ahead).
func parseSynCorpus(content string) []SynTest {
	lines := strings.SplitAfter(content, "\n")
	var tests []SynTest
	i := 0
	for i < len(lines) {
		if !reHeaderLine.MatchString(strings.TrimSuffix(lines[i], "\n")) {
			i++
			continue
		}
		// header: name lines until the next ===||| line
		j := i + 1
		var nameLines []string
		for j < len(lines) && !reHeaderLine.MatchString(strings.TrimSuffix(lines[j], "\n")) {
			nameLines = append(nameLines, strings.TrimRight(lines[j], "\r\n"))
			j++
		}
		if j >= len(lines) {
			break
		}
		j++ // past closing header line
		var inp strings.Builder
		for j < len(lines) && !reSepLine.MatchString(strings.TrimSuffix(lines[j], "\n")) {
			inp.WriteString(lines[j])
			j++
		}
		if j >= len(lines) {
			break
		}
		j++ // past separator
		var exp strings.Builder
		for j < len(lines) && !reHeaderLine.MatchString(strings.TrimSuffix(lines[j], "\n")) {
			exp.WriteString(lines[j])
			j++
		}
		t := SynTest{Input: inp.String(), Expect: exp.String()}
		var nm []string
		for _, l := range nameLines {
			if strings.HasPrefix(l, ":") {
				t.Attrs = append(t.Attrs, strings.ToLower(strings.TrimSpace(l[1:])))
			} else if len(t.Attrs) == 0 {
				nm = append(nm, l)
			}
		}
		t.Name = strings.TrimSpace(strings.Join(nm, "\n"))
		tests = append(tests, t)
		i = j
	}
	return tests
}

func (t SynTest) has(a string) bool {
	for _, x := range t.Attrs {
		if x == a {
			return true
		}
	}
	return false
}
