// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/jwriteclub/prudence"
)

var fakeTable = []AcquisitionOverride{{Set: "fake", ID: "F-1", Root: "A.tla", Visits: 3000000, Reason: "fake case"}}

func TestAcquisitionOverrideApplies(t *testing.T) {
	hit := &Case{Set: "fake", ID: "F-1", Root: "A.tla"}
	other := &Case{Set: "fake", ID: "F-2", Root: "A.tla"}
	if err := applyAcquisitionOverrides([]*Case{hit, other}, fakeTable); err != nil {
		t.Fatal(err)
	}
	if hit.Acquisition == nil || hit.Acquisition.MaxAcquisitionVisits != 3000000 || !hit.Acquisition.Applied {
		t.Errorf("override not recorded: %+v", hit.Acquisition)
	}
	if acquisitionVisits(hit) != 3000000 || acquisitionVisits(other) != 0 || other.Acquisition != nil {
		t.Errorf("value passed: %d / %d", acquisitionVisits(hit), acquisitionVisits(other))
	}
	// An entry whose case is not selected is skipped.
	if err := applyAcquisitionOverrides([]*Case{other}, fakeTable); err != nil {
		t.Error(err)
	}
}

func TestAcquisitionOverrideRejectsBadTable(t *testing.T) {
	c := func() []*Case { return []*Case{{Set: "fake", ID: "F-1", Root: "A.tla"}} }
	bad := map[string][]AcquisitionOverride{
		"root mismatch": {{Set: "fake", ID: "F-1", Root: "B.tla", Visits: 1, Reason: "r"}},
		"duplicate":     {fakeTable[0], fakeTable[0]},
		"zero value":    {{Set: "fake", ID: "F-1", Root: "A.tla", Visits: 0, Reason: "r"}},
		"negative":      {{Set: "fake", ID: "F-1", Root: "A.tla", Visits: -1, Reason: "r"}},
		"no reason":     {{Set: "fake", ID: "F-1", Root: "A.tla", Visits: 1}},
		"no set":        {{ID: "F-1", Root: "A.tla", Visits: 1, Reason: "r"}},
	}
	for name, table := range bad {
		if err := applyAcquisitionOverrides(c(), table); err == nil {
			t.Errorf("%s: accepted", name)
		}
	}
}

// The real table matches the manifest it names. This reads the manifest only; no case
// runs.
func TestAcquisitionTableMatchesManifest(t *testing.T) {
	if len(acquisitionOverrides) != 1 {
		t.Fatalf("table has %d entries; this test names the one expected", len(acquisitionOverrides))
	}
	o := acquisitionOverrides[0]
	if o.Set != "tlaplus-test-models" || o.ID != "B-test-model-bug156-te-9a905e" || o.Root != "TE.tla" || o.Visits != 3000000 {
		t.Errorf("TE entry changed: %+v", o)
	}
	cases, _, err := loadCases(filepath.Join("..", ".."), "tlaplus-test-models", 0, false, "")
	if err != nil {
		t.Fatal(err)
	}
	if err := applyAcquisitionOverrides(cases, acquisitionOverrides); err != nil {
		t.Fatal(err)
	}
	n := 0
	for _, c := range cases {
		if c.Acquisition != nil {
			n++
			if c.ID != o.ID {
				t.Errorf("override on %s", c.ID)
			}
		}
	}
	if n != 1 {
		t.Errorf("%d cases carry an override, want 1", n)
	}
}

// The plumbing reaches a real Check on a fake one-module case (never TE).
func TestAcquisitionOverrideFakeCaseRuns(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "A.tla"), []byte("---- MODULE A ----\nx == 1\n====\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	for _, visits := range []int{0, 3000000} {
		n := RunPrudence(dir, "A.tla", "", 60*time.Second, visits)
		if n.Class != ClassOK || n.Completion != "complete" {
			t.Errorf("visits %d: %s", visits, Show(n))
		}
	}
}

// A tiny override really reaches Check: the fake module completes at prudence's
// default budget and stops with CodeAcquisitionLimit at 5 visits, and the row says
// applied: true.
func TestAcquisitionOverrideChangesFakeCheck(t *testing.T) {
	recordInfo = true
	defer func() { recordInfo = false }()
	repo := t.TempDir()
	if err := os.WriteFile(filepath.Join(repo, "A.tla"), []byte("---- MODULE A ----\nx == 1\ny == <<x, TRUE>>\n====\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	sany := Norm{Class: ClassOK}
	run := func(table []AcquisitionOverride) *Case {
		c := &Case{Set: "fake", ID: "F-1", Root: "A.tla", Dir: repo, Sany: sany}
		if err := applyAcquisitionOverrides([]*Case{c}, table); err != nil {
			t.Fatal(err)
		}
		c.Prudence = RunPrudence(c.Dir, c.Root, "", 60*time.Second, acquisitionVisits(c))
		c.Result, c.Kind, c.Policy = Compare(c.Sany, c.Prudence)
		return c
	}
	limited := func(n Norm) bool {
		for _, d := range n.Info {
			if d.Code == prudence.CodeAcquisitionLimit {
				return true
			}
		}
		return false
	}
	def := run(nil)
	if def.Prudence.Class != ClassOK || def.Prudence.Completion != "complete" || limited(def.Prudence) || def.Result != "match" || def.Acquisition != nil {
		t.Fatalf("default budget: %s, result %s", Show(def.Prudence), def.Result)
	}
	small := run([]AcquisitionOverride{{Set: "fake", ID: "F-1", Root: "A.tla", Visits: 5, Reason: "fake: tiny budget"}})
	if small.Prudence.Class != ClassIncomp || small.Prudence.Completion != "incomplete" || !limited(small.Prudence) || small.Result != "incomplete" {
		t.Fatalf("5 visits: %s, info %+v, result %s", Show(small.Prudence), small.Prudence.Info, small.Result)
	}
	b, _ := json.Marshal(caseRecord(small))
	if !strings.Contains(string(b), `"acquisition_visits":{"max_acquisition_visits":5,"applied":true,"reason":"fake: tiny budget"}`) {
		t.Errorf("row: %s", b)
	}
}

func TestAcquisitionOverrideRowAndReport(t *testing.T) {
	ok := sem(at(1, 1, "x"))
	hit := &Case{Set: "fake", ID: "F-1", Root: "A.tla", Sany: ok, Prudence: ok}
	plain := &Case{Set: "fake", ID: "F-2", Root: "A.tla", Sany: ok, Prudence: ok}
	if err := applyAcquisitionOverrides([]*Case{hit, plain}, fakeTable); err != nil {
		t.Fatal(err)
	}
	for _, c := range []*Case{hit, plain} {
		c.Result, c.Kind, c.Policy = Compare(c.Sany, c.Prudence)
	}
	if hit.Result != "match" {
		t.Fatalf("override changed the outcome: %s", hit.Result)
	}
	b, _ := json.Marshal(caseRecord(hit))
	if !strings.Contains(string(b), `"acquisition_visits":{"max_acquisition_visits":3000000,"applied":`) {
		t.Errorf("row lacks the value: %s", b)
	}
	if b, _ := json.Marshal(caseRecord(plain)); strings.Contains(string(b), "acquisition") {
		t.Errorf("plain row carries an override: %s", b)
	}
	out := t.TempDir()
	if err := writeOutputs(out, []*Case{hit, plain}, []string{"fake"}, false, "", ""); err != nil {
		t.Fatal(err)
	}
	r, _ := os.ReadFile(filepath.Join(out, "report.md"))
	for _, want := range []string{"| **all** | 2 | 2 | 0 | 0 |", "**Acquisition-visit overrides**: 1 case(s)", "`fake/F-1`: `Options.MaxAcquisitionVisits` = 3000000", "scored match"} {
		if !strings.Contains(string(r), want) {
			t.Errorf("report lacks %q:\n%s", want, r)
		}
	}
	if !strings.Contains(string(r), "= 3000000 (applied; scored match)") {
		t.Errorf("applied state not reported:\n%s", r)
	}
}

// -rescore keeps the override the recorded run used and adds none.
func TestAcquisitionOverrideRescoreKeepsRecord(t *testing.T) {
	repo := t.TempDir()
	set := filepath.Join(repo, "sets", "g", "fake")
	manifest := `{"environment": "fake"}` + "\n"
	for _, id := range []string{"F-1", "F-2"} {
		if err := os.MkdirAll(filepath.Join(set, "cases", id), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(set, "cases", id, "A.tla"), []byte("---- MODULE A ----\n====\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		manifest += `{"id": "` + id + `", "root_file": "A.tla", "expected": {"status": "ok", "stdout": ""}}` + "\n"
	}
	if err := os.WriteFile(filepath.Join(set, "manifest.jsonl"), []byte(manifest), 0o644); err != nil {
		t.Fatal(err)
	}
	okNorm := Norm{Class: ClassOK, Completion: "complete"}
	rec := &AcquisitionRecord{MaxAcquisitionVisits: 3000000, Applied: true, Reason: "fake case"}
	in := filepath.Join(repo, "in.jsonl")
	var lines []string
	for _, c := range []*Case{{Set: "fake", ID: "F-1", Sany: okNorm, Prudence: okNorm, Result: "match", Acquisition: rec},
		{Set: "fake", ID: "F-2", Sany: okNorm, Prudence: okNorm, Result: "match"}} {
		b, _ := json.Marshal(caseRecord(c))
		lines = append(lines, string(b))
	}
	if err := os.WriteFile(in, []byte(strings.Join(lines, "\n")+"\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	out := t.TempDir()
	defer func() { rescoredFrom = "" }()
	if err := rescoreResults(repo, in, out, "", false); err != nil {
		t.Fatal(err)
	}
	got, _ := os.ReadFile(filepath.Join(out, "results.jsonl"))
	rows := strings.Split(strings.TrimSpace(string(got)), "\n")
	if len(rows) != 2 || !strings.Contains(rows[0], `"acquisition_visits":{"max_acquisition_visits":3000000,"applied":true`) || strings.Contains(rows[1], "acquisition") {
		t.Errorf("rescored rows:\n%s", got)
	}
}
