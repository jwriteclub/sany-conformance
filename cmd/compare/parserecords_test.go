// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

// recordedCase reads one case's recorded SANY output from the repository's manifest.
func recordedCase(t *testing.T, set, id string) *Case {
	t.Helper()
	ids := filepath.Join(t.TempDir(), "ids.txt")
	if err := os.WriteFile(ids, []byte(id+"\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	cases, _, err := loadCases(filepath.Join("..", ".."), set, 0, false, ids)
	if err != nil || len(cases) != 1 {
		t.Fatalf("loadCases %s/%s: %v, %d cases", set, id, err, len(cases))
	}
	return cases[0]
}

const (
	fairnessExc = `Was expecting "==== or more Module body" Encountered "<EOF>" at line 2, column 30 and token ")" Residual stack trace follows: Module definition starting at line 1, column 1.`
	atBangExc   = `Was expecting "= or ," Encountered "<EOF>" at line 2, column 19 and token "@" Residual stack trace follows: Except Spec starting at line 2, column 16. Some [] Form starting at line 2, column 6. ExtendableExpr starting at line 2, column 6. Expression starting at line 2, column 6. Definition starting at line 2, column 1.`
)

func parseNorm(d ...Diag) Norm { return Norm{Class: ClassParse, Diags: d} }

// The two corpus cases whose SANY output holds two parser records
// (TLAplusParser.java:113-116): each record is its own error.
func TestParseRecordsCorpus(t *testing.T) {
	for _, tc := range []struct {
		set, id string
		want    []Diag
		// prudence's two records as the 306282b public run recorded them
		pru []Diag
	}{
		{"fuzz-jazzer", "D-71c3f9889b63",
			[]Diag{{Line: 2, Col: 8, Msg: "Error in fairness expression at 2: 8"}, {Line: 2, Col: 30, Msg: fairnessExc}},
			[]Diag{{Line: 2, Col: 8, Msg: "Error in fairness expression at 2: 8"}, {Line: 2, Col: 30, Msg: fairnessExc}}},
		{"model-one-shot", "E-ed3076cde6ee",
			[]Diag{{Line: 0, Col: 0, Msg: "@ used in !.@"}, {Line: 2, Col: 19, Msg: atBangExc}},
			[]Diag{{Line: 2, Col: 18, Msg: "@ used in !.@"}, {Line: 2, Col: 19, Msg: atBangExc}}},
	} {
		c := recordedCase(t, tc.set, tc.id)
		got := NormalizeSany(c.Stdout)
		if got.Class != ClassParse || !reflect.DeepEqual(got.Diags, tc.want) {
			t.Errorf("%s: NormalizeSany = %+v, want %+v", tc.id, got, tc.want)
		}
		if r, k, _ := Compare(got, parseNorm(tc.pru...)); r != "match" {
			t.Errorf("%s: split vs prudence's split = %s %s, want match", tc.id, r, k)
		}
		// Negative controls: each must stay a mismatch.
		merged := Diag{Line: tc.want[1].Line, Col: tc.want[1].Col, Msg: tc.want[0].Msg + " " + tc.want[1].Msg}
		for name, pru := range map[string][]Diag{
			"prudence merged":       {merged},
			"prudence first only":   {tc.pru[0]},
			"prudence second only":  {tc.pru[1]},
			"texts swapped":         {{Line: tc.pru[0].Line, Col: tc.pru[0].Col, Msg: tc.pru[1].Msg}, {Line: tc.pru[1].Line, Col: tc.pru[1].Col, Msg: tc.pru[0].Msg}},
			"first text changed":    {{Line: tc.pru[0].Line, Col: tc.pru[0].Col, Msg: tc.pru[0].Msg + "!"}, tc.pru[1]},
			"second record moved":   {tc.pru[0], {Line: tc.pru[1].Line, Col: tc.pru[1].Col + 1, Msg: tc.pru[1].Msg}},
			"prudence record added": {tc.pru[0], tc.pru[1], tc.pru[1]},
		} {
			if r, k, _ := Compare(got, parseNorm(pru...)); r != "mismatch" {
				t.Errorf("%s, %s: %s %s, want mismatch", tc.id, name, r, k)
			}
		}
		// SANY merged (the old one-error projection) against prudence's split records.
		if r, k, _ := Compare(parseNorm(merged), parseNorm(tc.pru...)); r != "mismatch" || k != "position" {
			t.Errorf("%s: merged SANY vs split prudence = %s %s, want mismatch position", tc.id, r, k)
		}
	}
}

// Every parser record kind on its own is one error, as before the split.
func TestParseRecordsSingle(t *testing.T) {
	const pre = "\n****** SANY2 Version 2.2 created 08 July 2020\n\nParsing file $CASE/A.tla\n"
	const tail = "\n\nFatal errors while parsing TLA+ spec in file A.tla\n\nIn module A\n\nCould not parse module A from file A.tla\n*** Errors: 1\n"
	for out, want := range map[string]Diag{
		pre + "***Parse Error***\nEncountered \"x\" at line 3, column 4 and token \"y\" \n\nResidual stack trace follows:\nDefinition starting at line 3, column 1.\n" + tail: {Line: 3, Col: 4, Msg: `Encountered "x" at line 3, column 4 and token "y" Residual stack trace follows: Definition starting at line 3, column 1.`},
		pre + "Lexical error at line 5, column 7.  Encountered: \"$\" (36), after : \"\"\n" + tail:                                                                            {Line: 5, Col: 7, Msg: `Lexical error at line 5, column 7. Encountered: "$" (36), after : ""`},
		pre + "Lexical {error: EOF reached, possibly open comment starting around line 9\n" + tail:                                                                            {Msg: "Lexical {error: EOF reached, possibly open comment starting around line 9"},
		pre + "Error in fairness expression at 4: 2\n\n" + tail:                                                                                                               {Line: 4, Col: 2, Msg: "Error in fairness expression at 4: 2"},
	} {
		got := NormalizeSany(out)
		if got.Class != ClassParse || len(got.Diags) != 1 || got.Diags[0] != want {
			t.Errorf("NormalizeSany(%q) = %+v, want %+v", out, got, want)
		}
	}
}

// Records are sorted like every other error list (rule 4), whatever order SANY logged
// them in, and each record's position is its own.
func TestParseRecordsOrderAndOwnPosition(t *testing.T) {
	out := "Parsing file $CASE/A.tla\nError in fairness expression at 7: 3\n\n@ used in !.@\n***Parse Error***\nEncountered \"<EOF>\" at line 2, column 9 and token \"@\" \n\nResidual stack trace follows:\nDefinition starting at line 2, column 1.\n\n\nFatal errors while parsing TLA+ spec in file A.tla\n"
	want := []Diag{
		{Msg: "@ used in !.@"},
		{Line: 2, Col: 9, Msg: `Encountered "<EOF>" at line 2, column 9 and token "@" Residual stack trace follows: Definition starting at line 2, column 1.`},
		{Line: 7, Col: 3, Msg: "Error in fairness expression at 7: 3"},
	}
	if got := NormalizeSany(out).Diags; !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v\nwant %+v", got, want)
	}
	pru := parseNorm(Diag{Line: 1, Col: 1, Msg: "@ used in !.@"}, want[1], want[2])
	if r, k, _ := Compare(NormalizeSany(out), pru); r != "match" {
		t.Errorf("Compare = %s %s, want match", r, k)
	}
}

// A record start must be a whole line: a start text inside a line continues the record
// before it.
func TestParseRecordsStartIsWholeLine(t *testing.T) {
	out := "Parsing file $CASE/A.tla\n@ used in !.@\n x Error in fairness expression at 4: 2\n***Parse Error***\nEncountered \"a\" at line 2, column 3 and token \"b\" \n\nFatal errors\n"
	want := []Diag{
		{Line: 2, Col: 3, Msg: `Encountered "a" at line 2, column 3 and token "b"`},
		{Line: 4, Col: 2, Msg: "@ used in !.@ x Error in fairness expression at 4: 2"},
	}
	if got := NormalizeSany(out).Diags; !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v\nwant %+v", got, want)
	}
	// Blank lines alone are no record: the caller falls back to one error.
	if recs, ok := parseRecords([]string{"", "  ", "\t"}); ok || len(recs) != 0 {
		t.Errorf("blank text: %v %v, want no records", recs, ok)
	}
}

// Output that is not exactly a sequence of parser records keeps the old one-error
// projection: text before the first record, a record-like line that is not a whole
// start line, and no record at all.
func TestParseRecordsFallback(t *testing.T) {
	for out, want := range map[string]Diag{
		"Parsing file $CASE/A.tla\nsomething else\n@ used in !.@\n***Parse Error***\nEncountered \"a\" at line 2, column 3 and token \"b\" \n\nFatal errors\n": {Line: 2, Col: 3, Msg: `something else @ used in !.@ Encountered "a" at line 2, column 3 and token "b"`},
		"Parsing file $CASE/A.tla\n@ used in !.@ twice\n***Parse Error***\nEncountered \"a\" at line 2, column 3 and token \"b\" \n\nFatal errors\n":           {Line: 2, Col: 3, Msg: `@ used in !.@ twice Encountered "a" at line 2, column 3 and token "b"`},
		"Parsing file $CASE/A.tla\n x Error in fairness expression at 4: 2\n\nFatal errors\n":                                                                  {Line: 4, Col: 2, Msg: "x Error in fairness expression at 4: 2"},
	} {
		got := NormalizeSany(out)
		if len(got.Diags) != 1 || got.Diags[0] != want {
			t.Errorf("NormalizeSany(%q) = %+v, want %+v", out, got.Diags, want)
		}
	}
}
