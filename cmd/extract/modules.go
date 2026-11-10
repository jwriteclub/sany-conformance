// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

var (
	reModule   = regexp.MustCompile(`-{4,}\s*MODULE\s+([A-Za-z0-9_]+)\s*-{4,}`)
	reExtends  = regexp.MustCompile(`\bEXTENDS\s+([A-Za-z0-9_]+(?:\s*,\s*[A-Za-z0-9_]+)*)`)
	reInstance = regexp.MustCompile(`\bINSTANCE\s+([A-Za-z0-9_]+)`)
	rePluscal  = regexp.MustCompile(`--(?:fair\s+)?algorithm\b`)
)

// stripComments blanks TLA+ comments ((* nested *) and \* to end of line) and
// string contents, preserving offsets and newlines.
func stripComments(s string) string {
	b := []byte(s)
	n := len(b)
	depth := 0
	for i := 0; i < n; i++ {
		switch {
		case depth > 0:
			if b[i] == '(' && i+1 < n && b[i+1] == '*' {
				depth++
				b[i], b[i+1] = ' ', ' '
				i++
			} else if b[i] == '*' && i+1 < n && b[i+1] == ')' {
				depth--
				b[i], b[i+1] = ' ', ' '
				i++
			} else if b[i] != '\n' {
				b[i] = ' '
			}
		case b[i] == '(' && i+1 < n && b[i+1] == '*':
			depth = 1
			b[i], b[i+1] = ' ', ' '
			i++
		case b[i] == '\\' && i+1 < n && b[i+1] == '*':
			for i < n && b[i] != '\n' {
				b[i] = ' '
				i++
			}
		case b[i] == '"':
			i++
			for i < n && b[i] != '"' && b[i] != '\n' {
				if b[i] == '\\' && i+1 < n {
					b[i] = ' '
					i++
				}
				b[i] = ' '
				i++
			}
		}
	}
	return string(b)
}

// moduleInfo returns the declared module names in order and the referenced
// module names (EXTENDS and INSTANCE).
func moduleInfo(src string) (declared []string, refs []string) {
	code := stripComments(src)
	for _, m := range reModule.FindAllStringSubmatch(code, -1) {
		declared = append(declared, m[1])
	}
	seen := map[string]bool{}
	add := func(n string) {
		if !seen[n] {
			seen[n] = true
			refs = append(refs, n)
		}
	}
	for _, m := range reExtends.FindAllStringSubmatch(code, -1) {
		for _, n := range strings.Split(m[1], ",") {
			add(strings.TrimSpace(n))
		}
	}
	for _, m := range reInstance.FindAllStringSubmatch(code, -1) {
		add(m[1])
	}
	return
}

// Libs maps a module basename to the library that provides it.
type Libs struct {
	idx  map[string]string
	dirs map[string]string // library name -> directory
}

func buildLibs(ext string) *Libs {
	l := &Libs{idx: map[string]string{}, dirs: map[string]string{}}
	for _, d := range []struct{ dir, name string }{
		{"tlapm/library", "tlapm-library"},
		{"CommunityModules/modules", "CommunityModules"},
		{"tlaplus/tlatools/org.lamport.tlatools/src/tla2sany/StandardModules", "std"},
	} { // later entries win, so std takes priority
		l.dirs[d.name] = filepath.Join(ext, d.dir)
		ents, _ := os.ReadDir(filepath.Join(ext, d.dir))
		for _, e := range ents {
			if strings.HasSuffix(e.Name(), ".tla") {
				l.idx[strings.TrimSuffix(e.Name(), ".tla")] = d.name
			}
		}
	}
	return l
}

// Closure collects the local files (same directory as root) a root file needs,
// and the dependencies that must come from a library path.
type Closure struct {
	Local      map[string][]byte
	Ext        []Dep
	Unresolved []string
	Declared   []string
}

func closure(dir, rootName string, libs *Libs) (*Closure, error) {
	c := &Closure{Local: map[string][]byte{}}
	extSeen := map[string]bool{}
	queue := []string{rootName}
	for len(queue) > 0 {
		f := queue[0]
		queue = queue[1:]
		if _, ok := c.Local[f]; ok {
			continue
		}
		b, err := os.ReadFile(filepath.Join(dir, f))
		if err != nil {
			return nil, err
		}
		c.Local[f] = b
		declared, refs := moduleInfo(string(b))
		if f == rootName {
			c.Declared = declared
		}
		own := map[string]bool{}
		for _, d := range declared {
			own[d] = true
		}
		for _, r := range refs {
			if own[r] {
				continue
			}
			if _, err := os.Stat(filepath.Join(dir, r+".tla")); err == nil {
				queue = append(queue, r+".tla")
			} else if lib, ok := libs.idx[r]; ok {
				if !extSeen[r] {
					extSeen[r] = true
					c.Ext = append(c.Ext, Dep{Name: r, Lib: lib})
				}
			} else if !extSeen[r] {
				extSeen[r] = true
				c.Unresolved = append(c.Unresolved, r)
			}
		}
	}
	libs.followLibraryImports(c, extSeen)
	sort.Slice(c.Ext, func(i, j int) bool { return c.Ext[i].Name < c.Ext[j].Name })
	sort.Strings(c.Unresolved)
	return c, nil
}

// followLibraryImports adds what library modules themselves import, so a case's
// library list is complete: tlapm's FiniteSetTheorems extends Functions, which
// only CommunityModules provides. Standard modules are skipped, they are built in.
func (l *Libs) followLibraryImports(c *Closure, extSeen map[string]bool) {
	for i := 0; i < len(c.Ext); i++ {
		dep := c.Ext[i]
		if dep.Lib == "std" {
			continue
		}
		b, err := os.ReadFile(filepath.Join(l.dirs[dep.Lib], dep.Name+".tla"))
		if err != nil {
			continue
		}
		_, refs := moduleInfo(string(b))
		for _, r := range refs {
			lib, ok := l.idx[r]
			if ok && lib != "std" && !extSeen[r] {
				extSeen[r] = true
				c.Ext = append(c.Ext, Dep{Name: r, Lib: lib})
			}
		}
	}
}

// fileCase builds a case from a .tla file on disk; relPath is relative to the
// source repo root and used only for Origin.
func fileCase(libs *Libs, repoRoot, abs string) (*Case, error) {
	dir, base := filepath.Dir(abs), filepath.Base(abs)
	cl, err := closure(dir, base, libs)
	if err != nil {
		return nil, err
	}
	c := &Case{RootFile: base, content: cl.Local, ExtDeps: cl.Ext, Unresolved: cl.Unresolved}
	rel, _ := filepath.Rel(repoRoot, abs)
	c.Origin = rel
	if len(cl.Declared) == 0 {
		c.Flags = append(c.Flags, "no_module_header")
	} else {
		c.Root = cl.Declared[0]
		if c.Root != strings.TrimSuffix(base, ".tla") {
			c.Flags = append(c.Flags, "module_name_mismatch")
		}
	}
	if rePluscal.Match(cl.Local[base]) {
		c.Flags = append(c.Flags, "has_pluscal")
	}
	if len(cl.Local) > 1 {
		c.Flags = append(c.Flags, "multi_file")
	}
	return c, nil
}
