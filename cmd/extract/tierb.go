// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func sourceOf(rel string) string {
	parts := strings.Split(filepath.ToSlash(rel), "/")
	repo := parts[0]
	if len(parts) == 2 {
		return repo + "/(root)"
	}
	if repo == "tlaplus" {
		if len(parts) >= 5 && parts[1] == "tlatools" && parts[2] == "org.lamport.tlatools" {
			return "tlaplus/tlatools-" + parts[3]
		}
		return "tlaplus/" + parts[1]
	}
	if repo == "tlx" || repo == "tlapm" {
		if len(parts) > 3 {
			return repo + "/" + parts[1] + "/" + parts[2]
		}
	}
	return repo + "/" + parts[1]
}

// idPath is a readable tail of the path plus a short hash of the full path.
func idPath(rel string) string {
	r := strings.TrimSuffix(filepath.ToSlash(rel), ".tla")
	parts := strings.Split(r, "/")
	if len(parts) > 3 {
		parts = parts[len(parts)-3:]
	}
	return slug(strings.Join(parts, " ")) + "-" + sum([]byte(rel))[:6]
}

// tierB makes one case per .tla file of every source repo, except files that
// tier A already claims. Each case carries its local dependency closure.
func (x *ctx) tierB() []*Case {
	var paths []string
	for repo := range x.commits {
		filepath.Walk(filepath.Join(x.ext, repo), func(p string, fi os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			if fi.IsDir() && fi.Name() == ".git" {
				return filepath.SkipDir
			}
			if !fi.IsDir() && strings.HasSuffix(p, ".tla") {
				rel, _ := filepath.Rel(x.ext, p)
				if !x.claimed[filepath.ToSlash(rel)] {
					paths = append(paths, p)
				}
			}
			return nil
		})
	}
	sort.Strings(paths)
	var out []*Case
	for _, p := range paths {
		c, err := fileCase(x.libs, x.ext, p)
		if err != nil {
			fatal(err)
		}
		c.Tier, c.Source = "B", sourceOf(c.Origin)
		c.ID, c.Name = "B-"+idPath(c.Origin), filepath.Base(p)
		x.setCommit(c)
		out = append(out, c)
	}
	// other parsers' syntax corpora: copies and extensions of tlaplus's corpus
	for _, s := range []struct {
		dir, src, idp, basis string
		skip                 []string
	}{
		{"tree-sitter-tlaplus/test/corpus", "tree-sitter-tlaplus/syntax-corpus", "B-ts", "tree-sitter corpus, asserts tree-sitter not SANY", []string{"pluscal"}},
		{"tlapm/test/parser/syntax_corpus", "tlapm/syntax-corpus", "B-tlapm", "tlapm corpus, asserts tlapm not SANY", nil},
		{"rfcs/language_standard/tests/tlaplus_syntax", "rfcs/syntax-corpus", "B-rfcs", "language-standard corpus, tool-neutral", nil},
	} {
		for _, c := range x.synCorpusCases(s.dir, "B", s.src, s.idp, map[string]bool{}, s.basis, s.skip...) {
			c.Expect.Upstream = "other-parser"
			x.setCommit(c)
			out = append(out, c)
		}
	}
	return out
}
