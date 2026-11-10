// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// rescoredFrom names the results file a -rescore run read, for the report.
var rescoredFrom string

// sanyFromManifest (-sany-from-manifest, with -rescore) derives the SANY side again
// from the manifest's recorded output with the current NormalizeSany, instead of
// taking the row's recorded projection. Only for a results file from a recorded-SANY
// run: a -java run's SANY side is not in the manifest.
var sanyFromManifest bool

const note3d = "nested-module attribution ambiguous (rule 3d)"

// rescoreResults scores an existing results.jsonl again, with the rules as set now,
// and writes results.jsonl and report.md to out. Neither side is run: the SANY and
// prudence answers are the recorded ones. With -no-3d, rule 3d is undone first
// (undoNestedUnits). Every row must be a case that loadCases selects, and every
// selected case must have a row.
func rescoreResults(repo, in, out, stage string, deep bool) error {
	rescoredFrom = in
	all, sets, err := loadCases(repo, "", 0, deep, "")
	if err != nil {
		return err
	}
	byKey := map[string]*Case{}
	for _, c := range all {
		byKey[c.Set+"\x00"+c.ID] = c
	}
	f, err := os.Open(in)
	if err != nil {
		return err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 1<<20), 1<<30)
	seen := map[string]bool{}
	moved, orphans := 0, 0
	var resany []string
	for sc.Scan() {
		var row struct {
			Set      string `json:"set"`
			ID       string `json:"id"`
			Sany     Norm   `json:"sany"`
			Prudence Norm   `json:"prudence"`
			// The override the recorded run used, if any; the table is not consulted,
			// since a rescore does not run prudence.
			Acquisition *AcquisitionRecord `json:"acquisition_visits"`
		}
		if err := json.Unmarshal(sc.Bytes(), &row); err != nil {
			return fmt.Errorf("%s: %v", in, err)
		}
		k := row.Set + "\x00" + row.ID
		c := byKey[k]
		if c == nil || seen[k] {
			return fmt.Errorf("%s: row %s/%s is not a selected case, or is repeated", in, row.Set, row.ID)
		}
		seen[k] = true
		c.Sany, c.Prudence, c.Acquisition = row.Sany, row.Prudence, row.Acquisition
		if sanyFromManifest {
			n := NormalizeSany(c.Stdout)
			if c.Status == "crash" {
				n = Norm{Class: ClassCrash, Note: "recorded crash"}
			}
			was, _ := json.Marshal(c.Sany)
			now, _ := json.Marshal(n)
			if string(was) != string(now) {
				resany = append(resany, c.Set+"/"+c.ID)
			}
			c.Sany = n
		}
		if no3d && c.Prudence.Class == ClassSemantic {
			text, err := os.ReadFile(filepath.Join(c.Dir, c.Root))
			if err != nil {
				return err
			}
			m, o := undoNestedUnits(text, c.Prudence.Diags)
			moved += m
			orphans += o
			sortDiags(c.Prudence.Diags)
			if strings.HasPrefix(c.Prudence.Note, note3d) {
				c.Prudence.Note = ""
			}
		}
		c.Result, c.Kind, c.Policy = Compare(c.Sany, c.Prudence)
		c.FalseComplete = c.Prudence.Completion == "complete" && c.Prudence.Class == ClassOK && c.Sany.Class != ClassOK && c.Sany.Class != ClassCrash
	}
	if err := sc.Err(); err != nil {
		return err
	}
	if len(seen) != len(all) {
		return fmt.Errorf("%s has %d rows, %d cases selected", in, len(seen), len(all))
	}
	if sanyFromManifest {
		fmt.Fprintf(os.Stderr, "compare: SANY side derived again from the manifests; %d rows changed\n", len(resany))
		for _, k := range resany {
			fmt.Fprintf(os.Stderr, "compare: SANY side changed: %s\n", k)
		}
	}
	if no3d {
		fmt.Fprintf(os.Stderr, "compare: rule 3d undone on %d errors; %d errors carry a nested module name with no rule-3c twin and were left as they are\n", moved, orphans)
	}
	return writeOutputs(out, all, sets, false, "", stage)
}

// undoNestedUnits is the inverse of splitNestedUnits, for a recorded prudence answer
// (which keeps no byte offsets). Rule 3c never names a nested module, so a root-file
// error that names nested module N got there by rule 3d, which moved the second of two
// identical errors into N's unit. That needs N nested directly in a top-level module,
// declared once, with N and its parent closed. Such an error is put back at its file
// line (line + N's header line - 1) in rule-3c numbering, but only if its twin is
// there: an error with the same text, column, module and line that is not itself moved.
// An error that names N with no twin is left alone and counted as an orphan. An error
// whose position prudence reported as its unit's (Diag.Unit, rule 3e) was never moved
// by rule 3d: it is neither undone nor anyone's twin.
func undoNestedUnits(text []byte, diags []Diag) (moved, orphans int) {
	units := scanUnits(text)
	names := map[string]int{}
	for _, u := range units {
		names[u.name]++
	}
	nested := map[string]unit{}
	for _, u := range units {
		if u.parent >= 0 && u.depth == 1 && names[u.name] == 1 && u.end != 1<<30 && units[u.parent].end != 1<<30 {
			nested[u.name] = u
		}
	}
	twin := map[int]bool{}
	for i := range diags {
		u, ok := nested[diags[i].Module]
		if !ok || diags[i].Unit {
			continue
		}
		mod, line := moduleOf("", "", text, "", diags[i].Line+u.start-1)
		found := false
		for j := range diags {
			if j == i || twin[j] || diags[j].Unit {
				continue
			}
			if _, isNested := nested[diags[j].Module]; isNested {
				continue
			}
			if diags[j].Msg == diags[i].Msg && diags[j].Col == diags[i].Col && diags[j].Module == mod && diags[j].Line == line {
				twin[j] = true
				found = true
				break
			}
		}
		if !found {
			orphans++
			continue
		}
		diags[i].Module, diags[i].Line = mod, line
		moved++
	}
	return moved, orphans
}
