// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestParseSynCorpus(t *testing.T) {
	in := "=====|||\nA test\n:error\n=====|||\n\n---- MODULE Test ----\nx\n====\n\n-----|||\n\n=====|||\nB\n=====|||\nok\n-----|||\n(a)\n"
	ts := parseSynCorpus(in)
	if len(ts) != 2 || ts[0].Name != "A test" || !ts[0].has("error") || ts[1].Name != "B" || ts[1].Input != "ok\n" || ts[1].Expect != "(a)\n" {
		t.Fatalf("unexpected parse: %+v", ts)
	}
}

func TestJavaString(t *testing.T) {
	if got := javaString(`a /\\ b \"q\"`); got != `a /\ b "q"` {
		t.Fatal(got)
	}
}

// TestManifest checks the committed manifest when cases/ has been generated.
func TestManifest(t *testing.T) {
	f, err := os.Open("../../manifest.jsonl")
	if err != nil {
		t.Skip("no manifest")
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 1<<20), 1<<26)
	ids := map[string]bool{}
	for sc.Scan() {
		var c Case
		if err := json.Unmarshal(sc.Bytes(), &c); err != nil {
			t.Fatal(err)
		}
		if ids[c.ID] || c.Hash == "" || c.RootFile == "" || len(c.Files) == 0 || (c.Tier == "A") != (c.Expect != nil && c.Tier == "A") {
			t.Fatalf("bad record %s", c.ID)
		}
		if c.Tier == "C" && c.Mut == nil {
			t.Fatalf("tier C without mutant data: %s", c.ID)
		}
		ids[c.ID] = true
		if _, err := os.Stat(filepath.Join("../../cases", c.ID, c.RootFile)); err != nil {
			t.Skip("cases/ not generated")
		}
	}
}
