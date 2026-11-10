// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"
	"path/filepath"
	"testing"
)

// Rule 3c': a header is exactly what tla+.jj's BEGIN_MODULE, name token and SEPARATOR
// accept on one line.
func TestModuleHeaderAccepts(t *testing.T) {
	for line, want := range map[string]string{
		"---- MODULE M ----":                    "M",
		"---- MODULEOFt ----":                   "OFt", // D-cabb48f9aaa0: BEGIN_MODULE ends at MODULE
		"---- MODULEMODULE ----":                "MODULE",
		"---- MODULE MODULE ----":               "MODULE", // a keyword isFieldNameToken converts
		"----MODULE M----":                      "M",
		"-----   MODULE M---------":             "M",
		"---- MODULE\tM\t----":                  "M", // SPEC mode skips tabs after MODULE
		"---- MODULE _x ----":                   "_x",
		"---- MODULE 1a ----":                   "1a",
		"---- MODULE _1a ----":                  "_1a",
		"---- MODULE _ ----":                    "_", // US
		"---- MODULE WF_ ----":                  "WF_",
		"---- MODULE SF_ ----":                  "SF_",
		"---- MODULE WF ----":                   "WF",
		"---- MODULE WFx_1 ----":                "WFx_1",
		"---- MODULE W_x ----":                  "W_x",
		"---- MODULE S ----":                    "S",
		"---- MODULE ENABLED ----":              "ENABLED",
		"---- MODULE @ ----":                    "@",
		"---- MODULE ℕ ----":                    "ℕ",
		"x ---- MODULE M ---- y":                "M",
		"(* c *) ---- MODULE Ab_9 ---- \\* end": "Ab_9",
	} {
		got, ok := moduleHeader(line)
		if !ok || got != want {
			t.Errorf("moduleHeader(%q) = %q, %v; want %q", line, got, ok, want)
		}
	}
}

func TestModuleHeaderRejects(t *testing.T) {
	for _, line := range []string{
		"----\tMODULE M ----",   // BEGIN_MODULE allows spaces only before MODULE
		"---- MODULE M ----",    // nor any other space
		"--- MODULE M ----",     // BEGIN_MODULE needs four dashes
		"---- MODULE M ---",     // so does SEPARATOR
		"---- MODULE\fM ----",   // SPEC mode skips space, tab, CR, LF only
		"---- MODULE 123 ----",  // a number
		"---- MODULE _1 ----",   // US then a number
		"---- MODULE __ ----",   // two US tokens
		"---- MODULE WF_x ----", // WF_ then x
		"---- MODULE SF_x ----", // SF_ then x
		"---- MODULE M N ----",  // two names
		"---- MODULE ----",      // no name
		"---- MODULE M-x ----",  // M, then "-" is not SEPARATOR
		"---- module M ----",
		"---- MODULE ℕx ----", // NUMBER_SET then x
		"---- MODULE Mé ----", // not an identifier character
	} {
		if got, ok := moduleHeader(line); ok {
			t.Errorf("moduleHeader(%q) = %q, true; want no header", line, got)
		}
	}
}

// The corpus case D-cabb48f9aaa0: `---- MODULEOFt ----` opens OFt, nested in the root,
// so the root ends at line 14, not at OFt's ==== on line 10.
func TestModuleHeaderCorpusMODULEOFt(t *testing.T) {
	text, err := os.ReadFile(filepath.Join("..", "..", "sets", "ours-mechanical", "fuzz-jazzer", "cases", "D-cabb48f9aaa0", "TSFSpecInsEOFt.tla"))
	if err != nil {
		t.Fatal(err)
	}
	units := scanUnits(text)
	want := []unit{
		{name: "TSFSpecInsEOFt", start: 1, end: 14, parent: -1, depth: 0},
		{name: "OFt", start: 2, end: 10, parent: 0, depth: 1},
		{name: "Ring", start: 3, end: 7, parent: 1, depth: 2},
	}
	if len(units) != len(want) {
		t.Fatalf("scanUnits = %+v", units)
	}
	for i := range want {
		if units[i] != want[i] {
			t.Errorf("unit %d = %+v, want %+v", i, units[i], want[i])
		}
	}
	for line, want := range map[int]int{2: 2, 5: 5, 9: 9, 11: 11, 13: 13, 14: 14} {
		mod, got := moduleOf("", "", text, "", line)
		if mod != "TSFSpecInsEOFt" || got != want {
			t.Errorf("moduleOf(line %d) = %q %d, want TSFSpecInsEOFt %d", line, mod, got, want)
		}
	}
}

// A header the grammar rejects opens nothing: the tab before MODULE leaves the second
// module unopened, so its ==== closes the root.
func TestModuleHeaderRejectedOpensNothing(t *testing.T) {
	text := []byte("---- MODULE Root ----\n----\tMODULE Inner ----\nx == 1\n====\ny == 2\n====\n")
	mods := fileModules(text)
	if len(mods) != 1 || mods[0].name != "Root" || mods[0].end != 4 {
		t.Fatalf("fileModules = %+v", mods)
	}
	if mod, line := moduleOf("", "", text, "", 5); mod != "" || line != 5 {
		t.Errorf("moduleOf(5) = %q %d, want no module", mod, line)
	}
}
