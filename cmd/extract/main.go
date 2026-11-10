// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

// Command extract builds cases/ and manifest.jsonl from external/.
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func readLock(path string) map[string]string {
	b, err := os.ReadFile(path)
	if err != nil {
		fatal(err)
	}
	m := map[string]string{}
	for _, l := range strings.Split(string(b), "\n") {
		f := strings.Fields(l)
		if len(f) == 3 && !strings.HasPrefix(l, "#") {
			m[f[0]] = f[2]
		}
	}
	return m
}

func main() {
	ext := flag.String("external", "external", "fetched sources")
	out := flag.String("cases", "cases", "generated case directory")
	manifest := flag.String("manifest", "manifest.jsonl", "manifest path")
	lock := flag.String("lock", "sources.lock", "lock file")
	seed := flag.Int64("seed", 20261103, "mutation seed")
	cap := flag.Int("mutant-cap", 3000, "maximum tier C cases")
	flag.Parse()

	x := &ctx{ext: *ext, commits: readLock(*lock), claimed: map[string]bool{}}
	x.libs = buildLibs(*ext)
	all := x.tierA()
	all = append(all, x.tierB()...)

	// de-duplicate by content hash; tier A wins, then source order
	for _, c := range all {
		c.finalize()
	}
	canon := map[string]*Case{}
	for _, tier := range []string{"A", "B"} {
		for _, c := range all {
			if c.Tier != tier {
				continue
			}
			if p, ok := canon[c.Hash]; ok {
				c.DupOf = p.ID
			} else {
				canon[c.Hash] = c
			}
		}
	}
	all = append(all, x.mutants(all, canon, *seed, *cap)...)

	sort.SliceStable(all, func(i, j int) bool {
		if all[i].Tier != all[j].Tier {
			return all[i].Tier < all[j].Tier
		}
		return all[i].ID < all[j].ID
	})
	ids := map[string]bool{}
	for _, c := range all {
		if ids[c.ID] {
			fatal(fmt.Errorf("duplicate id %s", c.ID))
		}
		ids[c.ID] = true
	}

	if err := os.RemoveAll(*out); err != nil {
		fatal(err)
	}
	mf, err := os.Create(*manifest)
	if err != nil {
		fatal(err)
	}
	w := bufio.NewWriter(mf)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	for _, c := range all {
		dir := filepath.Join(*out, c.ID)
		for p, b := range c.content {
			write(filepath.Join(dir, p), b)
		}
		for p, b := range c.aux {
			write(filepath.Join(dir, p), b)
		}
		if err := enc.Encode(c); err != nil {
			fatal(err)
		}
	}
	w.Flush()
	mf.Close()
	fmt.Printf("wrote %d cases\n", len(all))
}

func write(path string, b []byte) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		fatal(err)
	}
	if err := os.WriteFile(path, b, 0o644); err != nil {
		fatal(err)
	}
}
