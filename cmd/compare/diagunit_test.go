// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/jwriteclub/prudence"
)

// unitRoot is authored for the projection alone (no SANY output is claimed for it):
// an error at y inside nested module Inner, which a reported unit would place at
// Inner's line 2. The tab puts y at SANY column 14.
const unitRoot = "---- MODULE Root ----\n---- MODULE Inner ----\n\tx == y\n====\nI == INSTANCE Inner\n====\n"

const unitMsg = "Unknown operator: `y'."

func unitDiag() prudence.Diagnostic {
	off := strings.Index(unitRoot, "y\n")
	return prudence.Diagnostic{Phase: prudence.LevelSemantic, Severity: prudence.SeverityError, Code: "semantic/4202",
		Message: unitMsg, Path: "Root.tla",
		Span: prudence.Span{Start: prudence.Position{ByteOffset: off, Line: 2, Column: 6}, End: prudence.Position{ByteOffset: off + 1, Line: 2, Column: 7}}}
}

// Rule 3e: a reported unit's coordinates are used as they are; without one, rules 3a-3c.
func TestProjectDiagUnit(t *testing.T) {
	d := unitDiag()
	got, off := projectDiag("", "Root.tla", []byte(unitRoot), d, nil)
	if want := (Diag{Module: "Root", Line: 3, Col: 14, Msg: unitMsg}); got != want || off != d.Span.Start.ByteOffset {
		t.Errorf("no unit: %+v %d, want %+v %d", got, off, want, d.Span.Start.ByteOffset)
	}
	got, off = projectDiag("", "Root.tla", []byte(unitRoot), d, &reportedUnit{Module: "Inner", Line: 2, Col: 14})
	if want := (Diag{Module: "Inner", Line: 2, Col: 14, Msg: unitMsg, Unit: true}); got != want || off != -1 {
		t.Errorf("unit: %+v %d, want %+v -1", got, off, want)
	}
	// Authoritative: not recomputed from the Span or the root text.
	got, _ = projectDiag("", "Root.tla", []byte(unitRoot), d, &reportedUnit{Module: "Elsewhere", Line: 40, Col: 99})
	if want := (Diag{Module: "Elsewhere", Line: 40, Col: 99, Msg: unitMsg, Unit: true}); got != want {
		t.Errorf("unit taken as reported: %+v, want %+v", got, want)
	}
}

// fakeUnits stands in for prudence's Diagnostic.Unit: the n-th diagnostic with
// message msg gets units[n], every other diagnostic none.
func fakeUnits(t *testing.T, msg string, units ...*reportedUnit) {
	t.Helper()
	old := unitOf
	n := 0
	unitOf = func(d prudence.Diagnostic) *reportedUnit {
		if d.Message != msg || n >= len(units) {
			return nil
		}
		n++
		return units[n-1]
	}
	t.Cleanup(func() { unitOf = old })
}

// The corpus case D-775b33c0c1e1 (group E of NEXT-DIAGNOSTIC-DIFFERENCES-HANDOFF.md):
// SANY extracts the nested module Inner (file lines 5-8) as a unit of its own and
// prints the error at y twice: "line 6, col 6 ... of module Ie_to_omitted_modunits_test"
// and "line 2, col 6 ... of module Inner". Prudence (the pinned commit) reports both
// copies at the same root bytes. Real prudence.Check runs; the stand-in supplies the
// Unit prudence would add to the extracted copy: Inner, 2, 6.
const (
	e775Set = "fuzz-jazzer"
	e775ID  = "D-775b33c0c1e1"
	e775Msg = "Unknown operator: `y'."
)

func e775(t *testing.T, units ...*reportedUnit) (Norm, Norm) {
	t.Helper()
	c := recordedCase(t, e775Set, e775ID)
	fakeUnits(t, e775Msg, units...)
	return NormalizeSany(c.Stdout), RunPrudence(c.Dir, c.Root, "", time.Minute, 0)
}

// Rule 3e against SANY's recorded output: the unit copy matches SANY's "of module
// Inner" record; a wrong unit, N-relative line or column, a missing Unit, or a Unit on
// the root copy is a position mismatch; an invalid Unit is a prudence error.
func TestUnitCorpusControls(t *testing.T) {
	defer func(v bool) { no3d = v }(no3d)
	good := &reportedUnit{"Inner", 2, 6}
	for _, off := range []bool{true, false} {
		no3d = off
		_, pru := e775(t)
		if n := strings.Count(Show(pru), e775Msg); pru.Class != ClassSemantic || n != 2 {
			t.Fatalf("no3d=%v: this test needs prudence to report %q twice:\n%s", off, e775Msg, Show(pru))
		}
		for name, units := range map[string][]*reportedUnit{
			"unit on the second copy": {nil, good},
			"unit on the first copy":  {good, nil},
		} {
			sany, pru := e775(t, units...)
			if r, k, _ := Compare(sany, pru); r != "match" || pru.Note != "" {
				t.Errorf("no3d=%v, %s: %s %s, note %q\n%s", off, name, r, k, pru.Note, Show(pru))
			}
		}
		for name, units := range map[string][]*reportedUnit{
			"wrong unit":                  {nil, {"Ie_to_omitted_modunits_test", 2, 6}},
			"other unit":                  {nil, {"Outer", 2, 6}},
			"wrong N-relative line":       {nil, {"Inner", 3, 6}},
			"unit line in root numbering": {nil, {"Inner", 6, 6}},
			"wrong column":                {nil, {"Inner", 2, 7}},
			"unit on both copies":         {good, good},
		} {
			sany, pru := e775(t, units...)
			if r, k, _ := Compare(sany, pru); r != "mismatch" || k != "position" {
				t.Errorf("no3d=%v, %s: %s %s, want mismatch position\n%s", off, name, r, k, Show(pru))
			}
		}
		for _, bad := range []*reportedUnit{{"", 2, 6}, {"Inner", 0, 6}, {"Inner", 2, 0}, {"Inner", -1, 6}} {
			sany, pru := e775(t, nil, bad)
			if r, _, _ := Compare(sany, pru); pru.Class != ClassCrash || r != "error" || !strings.Contains(pru.Note, "invalid Diagnostic.Unit") {
				t.Errorf("no3d=%v, invalid unit %+v: %s, %+v, want a prudence error", off, *bad, r, pru)
			}
		}
	}
	// Without Unit: with -no-3d the case stays the visible mismatch it is today; with
	// rule 3d on, 3d still pairs the two copies as before.
	no3d = true
	if sany, pru := e775(t); func() bool { r, k, _ := Compare(sany, pru); return r != "mismatch" || k != "position" }() {
		t.Errorf("no unit, -no-3d: want mismatch position\n%s", Show(pru))
	}
	no3d = false
	if sany, pru := e775(t); func() bool { r, _, _ := Compare(sany, pru); return r != "match" }() {
		t.Errorf("no unit, rule 3d on: want match\n%s", Show(pru))
	}
}

// -no-3d on a recorded row undoes rule 3d, but never touches a unit record.
func TestUndoNestedUnitsSkipsUnit(t *testing.T) {
	text := []byte(unitRoot)
	unitRec := []Diag{{Module: "Root", Line: 3, Col: 14, Msg: unitMsg}, {Module: "Inner", Line: 2, Col: 14, Msg: unitMsg, Unit: true}}
	if moved, orphans := undoNestedUnits(text, unitRec); moved != 0 || orphans != 0 || unitRec[1].Module != "Inner" || unitRec[1].Line != 2 {
		t.Errorf("unit record: moved %d orphans %d, %+v", moved, orphans, unitRec)
	}
	// A unit record is nobody's twin either.
	twinless := []Diag{{Module: "Root", Line: 3, Col: 14, Msg: unitMsg, Unit: true}, {Module: "Inner", Line: 2, Col: 14, Msg: unitMsg}}
	if moved, orphans := undoNestedUnits(text, twinless); moved != 0 || orphans != 1 {
		t.Errorf("unit record as twin: moved %d orphans %d, %+v", moved, orphans, twinless)
	}
	moved3d := []Diag{{Module: "Root", Line: 3, Col: 14, Msg: unitMsg}, {Module: "Inner", Line: 2, Col: 14, Msg: unitMsg}}
	if moved, _ := undoNestedUnits(text, moved3d); moved != 1 || moved3d[1].Module != "Root" || moved3d[1].Line != 3 {
		t.Errorf("rule-3d record: moved %d, %+v", moved, moved3d)
	}
}

// "unit" appears in a row only on a unit record.
func TestDiagUnitJSON(t *testing.T) {
	b, _ := json.Marshal(Diag{Module: "M", Line: 1, Col: 2, Msg: "x"})
	if string(b) != `{"module":"M","line":1,"col":2,"msg":"x"}` {
		t.Errorf("no unit: %s", b)
	}
	b, _ = json.Marshal(Diag{Module: "M", Line: 1, Col: 2, Msg: "x", Unit: true})
	if string(b) != `{"module":"M","line":1,"col":2,"msg":"x","unit":true}` {
		t.Errorf("unit: %s", b)
	}
	var d Diag
	if err := json.Unmarshal(b, &d); err != nil || !d.Unit {
		t.Errorf("round trip: %+v %v", d, err)
	}
}

// -rescore -sany-from-manifest derives the SANY side again from the manifest; without
// the flag the recorded projection is kept.
func TestRescoreSanyFromManifest(t *testing.T) {
	repo := t.TempDir()
	set := filepath.Join(repo, "sets", "g", "fake")
	stdout := "Parsing file $CASE/A.tla\n@ used in !.@\n***Parse Error***\nEncountered \"<EOF>\" at line 2, column 9 and token \"@\" \n\nFatal errors while parsing TLA+ spec in file A.tla\n"
	if err := os.MkdirAll(filepath.Join(set, "cases", "P"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(set, "cases", "P", "A.tla"), []byte("---- MODULE A ----\n====\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	row, _ := json.Marshal(map[string]any{"id": "P", "root_file": "A.tla", "expected": map[string]any{"status": "ok", "stdout": stdout}})
	if err := os.WriteFile(filepath.Join(set, "manifest.jsonl"), append([]byte("{\"environment\": \"fake\"}\n"), append(row, '\n')...), 0o644); err != nil {
		t.Fatal(err)
	}
	exc := `Encountered "<EOF>" at line 2, column 9 and token "@"`
	pru := Norm{Class: ClassParse, Diags: []Diag{{Line: 2, Col: 8, Msg: "@ used in !.@"}, {Line: 2, Col: 9, Msg: exc}}}
	old := Norm{Class: ClassParse, Diags: []Diag{{Line: 2, Col: 9, Msg: "@ used in !.@ " + exc}}}
	rec, _ := json.Marshal(map[string]any{"set": "fake", "id": "P", "sany": old, "prudence": pru})
	in := filepath.Join(repo, "in.jsonl")
	if err := os.WriteFile(in, append(rec, '\n'), 0o644); err != nil {
		t.Fatal(err)
	}
	defer func(v bool) { sanyFromManifest = v }(sanyFromManifest)
	for _, tc := range []struct {
		flag   bool
		result string
		sany   int
	}{{false, `"result":"mismatch"`, 1}, {true, `"result":"match"`, 2}} {
		sanyFromManifest = tc.flag
		out := t.TempDir()
		if err := rescoreResults(repo, in, out, "", false); err != nil {
			t.Fatal(err)
		}
		got, _ := os.ReadFile(filepath.Join(out, "results.jsonl"))
		var r struct{ Sany Norm }
		if err := json.Unmarshal(got, &r); err != nil {
			t.Fatal(err)
		}
		if !bytes.Contains(got, []byte(tc.result)) || len(r.Sany.Diags) != tc.sany {
			t.Errorf("flag %v: %s", tc.flag, got)
		}
	}
}
