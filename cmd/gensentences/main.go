// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

// Command gensentences regenerates the grammar sentences from the pinned tla+.jj: one
// short sentence per choice point, in the manner of Purdom, ignoring semantic
// lookahead and actions. Each goes to <out>/S-<hash>/<Module>.tla. Deterministic.
//
//	gensentences -jj external/tlaplus/tlatools/org.lamport.tlatools/javacc/tla+.jj -out DIR
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"sany-conformance/internal/grammar"
)

var reModule = regexp.MustCompile(`-{4,}[ \t\r\n]*MODULE[ \t\r\n]+([A-Za-z0-9_]{1,64})`)

func main() {
	jj := flag.String("jj", "", "tla+.jj")
	out := flag.String("out", "", "output directory")
	flag.Parse()
	if err := run(*jj, *out); err != nil {
		fmt.Fprintln(os.Stderr, "gensentences:", err)
		os.Exit(2)
	}
}

func run(jj, out string) error {
	g, err := grammar.ParseGrammar(jj)
	if err != nil {
		return err
	}
	table, err := grammar.TokenTable(jj)
	if err != nil {
		return err
	}
	sents, points := g.Sentences("CompilationUnit", table)
	seen := map[string]bool{}
	for _, s := range sents {
		text := []byte(grammar.WrapSentence(s))
		name := "Fuzz"
		if m := reModule.FindSubmatch(text); m != nil {
			name = string(m[1])
		}
		root := name + ".tla"
		h := sha256.New()
		fmt.Fprintf(h, "root %s\n", root)
		fmt.Fprintf(h, "%s %d\n", root, len(text))
		h.Write(text)
		sum := hex.EncodeToString(h.Sum(nil))
		if seen[sum] {
			continue
		}
		seen[sum] = true
		d := filepath.Join(out, "S-"+sum[:12])
		if err := os.MkdirAll(d, 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(d, root), text, 0o644); err != nil {
			return err
		}
	}
	fmt.Printf("%d sentences from %d choice points, %d distinct\n", len(sents), points, len(seen))
	return nil
}
