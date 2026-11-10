// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"strings"
	"testing"
	"time"

	"github.com/jwriteclub/prudence"
)

// unitOf is a direct copy of Diagnostic.Unit: nil stays nil, fields map one to one.
func TestUnitOfReadsField(t *testing.T) {
	if u := unitOf(prudence.Diagnostic{}); u != nil {
		t.Errorf("no Unit: %+v", u)
	}
	got := unitOf(prudence.Diagnostic{Unit: &prudence.DiagnosticUnit{Module: "N", Line: 3, Column: 17}})
	if got == nil || *got != (reportedUnit{Module: "N", Line: 3, Col: 17}) {
		t.Errorf("Unit: %+v", got)
	}
}

// With prudence's own Diagnostic.Unit, D-775b33c0c1e1 matches recorded SANY with rule
// 3d off, and the unit record is the one prudence reported.
func TestUnitCorpusRealField(t *testing.T) {
	defer func(v bool) { no3d = v }(no3d)
	no3d = true
	c := recordedCase(t, e775Set, e775ID)
	pru := RunPrudence(c.Dir, c.Root, "", time.Minute, 0)
	units := 0
	for _, d := range pru.Diags {
		if d.Unit {
			units++
			if d != (Diag{Module: "Inner", Line: 2, Col: 6, Msg: e775Msg, Unit: true}) {
				t.Errorf("unit record %+v", d)
			}
		}
	}
	if r, k, _ := Compare(NormalizeSany(c.Stdout), pru); r != "match" || units != 1 {
		t.Errorf("%s %s, %d unit records\n%s", r, k, units, Show(pru))
	}
	if !strings.Contains(Show(pru), "Inner 2:6 (unit)") {
		t.Errorf("Show: %s", Show(pru))
	}
}
