// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//go:embed scopes.conf
var defaultScopes string

// ScopeDef is one named set of entries and cuts.
type ScopeDef struct {
	Name    string
	Entries []string // entry set names
	Cuts    []string // cut names
	View    string   // reporting view, ViewNtkConstant or empty; never touches the map
	Why     string
}

// HarnessDef names extra classes loaded for a soundness comparison, with roots.
type HarnessDef struct {
	Name  string
	Dirs  []string
	Roots []string // class names or prefixes; "*" means every harness class
	Why   string
}

// RunDef holds a coverage report to its scope and harness.
type RunDef struct {
	Glob    string
	Scope   string
	Harness string
	Kind    string
	// ExtraCuts are cuts held against this report on top of its scope's.
	ExtraCuts []string
}

// Explanation accounts for a method outside its run's scope.
type Explanation struct {
	Kind   string
	Method string
	Why    string
}

// CutDef is one cut with its documentation.
type CutDef struct {
	Cut
	Flag string
}

// MethodCutDef removes one method from a scope: nothing can reach it. A line cut
// removes call sites; this removes the edge into a method whose only route is an
// assumption, such as a JDK callback.
type MethodCutDef struct {
	Name   string
	Method string // class.name:descriptor
	Flag   string
	Why    string
}

// ViewNtkConstant reports with the jj_ntk ternary treated as constant.
const ViewNtkConstant = "ntk-constant"

// Config is the parsed scopes file.
type Config struct {
	EntrySets map[string][]string
	Cuts      []CutDef
	MethodCut []MethodCutDef
	Scopes    []ScopeDef
	Harnesses []HarnessDef
	Runs      []RunDef
	Explains  []Explanation
	Reflects  []ReflectDef
	Text      string
}

// ReflectDef resolves one reflective call site by hand: the method that holds it
// initialises this class when it runs.
type ReflectDef struct {
	Caller string
	Target string
	Why    string
}

// reflectFor finds the resolution of a reflective site in a method.
func (c *Config) reflectFor(caller string) *ReflectDef {
	for i := range c.Reflects {
		if c.Reflects[i].Caller == caller {
			return &c.Reflects[i]
		}
	}
	return nil
}

func splitList(s string) []string {
	if s == "" {
		return nil
	}
	return strings.Split(s, ",")
}

// keyValues reads key=value words; why= takes the rest of the line.
func keyValues(words []string) map[string]string {
	out := map[string]string{}
	for i, w := range words {
		k, v, ok := strings.Cut(w, "=")
		if !ok {
			continue
		}
		if k == "why" {
			out[k] = strings.Join(append([]string{v}, words[i+1:]...), " ")
			break
		}
		out[k] = v
	}
	return out
}

// ParseConfig reads the scopes file format described at its top.
func ParseConfig(text string) (*Config, error) {
	c := &Config{EntrySets: map[string][]string{}, Text: text}
	for n, line := range strings.Split(text, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || line[0] == '#' {
			continue
		}
		w := strings.Fields(line)
		fail := func(msg string) error { return fmt.Errorf("scopes line %d: %s: %s", n+1, msg, line) }
		switch w[0] {
		case "entries":
			if len(w) != 3 {
				return nil, fail("want: entries <set> <method>")
			}
			c.EntrySets[w[1]] = append(c.EntrySets[w[1]], w[2])
		case "cut":
			if len(w) < 5 {
				return nil, fail("want: cut <name> <class> <from> <to>")
			}
			from, err1 := strconv.Atoi(w[3])
			to, err2 := strconv.Atoi(w[4])
			if err1 != nil || err2 != nil || from > to {
				return nil, fail("bad line range")
			}
			kv := keyValues(w[5:])
			c.Cuts = append(c.Cuts, CutDef{Cut: Cut{Name: w[1], Class: w[2], From: from, To: to, Why: kv["why"]}, Flag: kv["flag"]})
		case "cutmethod":
			if len(w) < 3 {
				return nil, fail("want: cutmethod <name> <class.name:desc> flag=... why=...")
			}
			kv := keyValues(w[3:])
			c.MethodCut = append(c.MethodCut, MethodCutDef{Name: w[1], Method: w[2], Flag: kv["flag"], Why: kv["why"]})
		case "scope":
			if len(w) < 2 {
				return nil, fail("want: scope <name> ...")
			}
			kv := keyValues(w[2:])
			if v := kv["view"]; v != "" && v != ViewNtkConstant {
				return nil, fail("unknown view")
			}
			c.Scopes = append(c.Scopes, ScopeDef{Name: w[1], Entries: splitList(kv["entries"]), Cuts: splitList(kv["cuts"]), View: kv["view"], Why: kv["why"]})
		case "harness":
			if len(w) < 2 {
				return nil, fail("want: harness <name> ...")
			}
			kv := keyValues(w[2:])
			c.Harnesses = append(c.Harnesses, HarnessDef{Name: w[1], Dirs: splitList(kv["dirs"]), Roots: splitList(kv["roots"]), Why: kv["why"]})
		case "run":
			if len(w) < 2 {
				return nil, fail("want: run <glob> ...")
			}
			kv := keyValues(w[2:])
			c.Runs = append(c.Runs, RunDef{Glob: w[1], Scope: kv["scope"], Harness: kv["harness"], Kind: kv["kind"], ExtraCuts: splitList(kv["extracuts"])})
		case "explain":
			if len(w) < 3 {
				return nil, fail("want: explain <kind> <method> why=...")
			}
			kv := keyValues(w[3:])
			c.Explains = append(c.Explains, Explanation{Kind: w[1], Method: w[2], Why: kv["why"]})
		case "reflect":
			if len(w) < 3 {
				return nil, fail("want: reflect <method> <class> why=...")
			}
			kv := keyValues(w[3:])
			c.Reflects = append(c.Reflects, ReflectDef{Caller: w[1], Target: w[2], Why: kv["why"]})
		default:
			return nil, fail("unknown directive")
		}
	}
	return c, c.validate()
}

// requiredScopes are named by the report's comparisons. A custom file must keep
// them; it can change what they contain, and add others.
var requiredScopes = []string{"cli", "cli-noexp", "cli+xml-noexp", "cli+xml-nodebug", "parse-noexp", "port"}

func (c *Config) validate() error {
	for _, name := range requiredScopes {
		if c.Scope(name) == nil {
			return fmt.Errorf("scopes file must define scope %s", name)
		}
	}
	for _, s := range c.Scopes {
		for _, e := range s.Entries {
			if len(c.EntrySets[e]) == 0 {
				return fmt.Errorf("scope %s: no entry set %s", s.Name, e)
			}
		}
		for _, name := range s.Cuts {
			if !c.hasCut(name) {
				return fmt.Errorf("scope %s: no cut %s", s.Name, name)
			}
		}
	}
	for _, r := range c.Runs {
		if c.Scope(r.Scope) == nil {
			return fmt.Errorf("run %s: no scope %s", r.Glob, r.Scope)
		}
		if c.Harness(r.Harness) == nil {
			return fmt.Errorf("run %s: no harness %s", r.Glob, r.Harness)
		}
		for _, name := range r.ExtraCuts {
			if !c.hasCut(name) {
				return fmt.Errorf("run %s: no cut %s", r.Glob, name)
			}
		}
	}
	return nil
}

// CutsNamed lists the cuts with a name.
func (c *Config) CutsNamed(name string) []Cut {
	var out []Cut
	for _, d := range c.Cuts {
		if d.Name == name {
			out = append(out, d.Cut)
		}
	}
	return out
}

// hasCut says that a name is a line cut or a method cut.
func (c *Config) hasCut(name string) bool {
	if len(c.CutsNamed(name)) > 0 {
		return true
	}
	for i := range c.MethodCut {
		if c.MethodCut[i].Name == name {
			return true
		}
	}
	return false
}

// Scope finds a scope by name.
func (c *Config) Scope(name string) *ScopeDef {
	for i := range c.Scopes {
		if c.Scopes[i].Name == name {
			return &c.Scopes[i]
		}
	}
	return nil
}

// Harness finds a harness by name.
func (c *Config) Harness(name string) *HarnessDef {
	for i := range c.Harnesses {
		if c.Harnesses[i].Name == name {
			return &c.Harnesses[i]
		}
	}
	return nil
}

// loadConfig reads a scopes file, or the embedded one when path is empty.
func loadConfig(path string) (*Config, error) {
	text := defaultScopes
	if path != "" {
		b, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		text = string(b)
	}
	return ParseConfig(text)
}

// scopeOptions turns a scope into analysis options for a world.
func scopeOptions(c *Config, w *World, s *ScopeDef, p Precision) (Options, error) {
	return scopeOptionsExtra(c, w, s, p, nil)
}

// scopeOptionsExtra adds further cuts by name, for a run held to more than its scope.
func scopeOptionsExtra(c *Config, w *World, s *ScopeDef, p Precision, extra []string) (Options, error) {
	o := Options{Precision: p, Reflect: map[string]string{}}
	for _, r := range c.Reflects {
		o.Reflect[r.Caller] = r.Target
	}
	for _, set := range s.Entries {
		for _, ref := range c.EntrySets[set] {
			m, err := w.MethodByRef(ref)
			if err != nil {
				return o, fmt.Errorf("scope %s: %w", s.Name, err)
			}
			o.Entries = append(o.Entries, m)
		}
	}
	for _, name := range append(append([]string(nil), s.Cuts...), extra...) {
		o.Cuts = append(o.Cuts, c.CutsNamed(name)...)
		for i := range c.MethodCut {
			if c.MethodCut[i].Name != name {
				continue
			}
			m, err := w.MethodByRef(c.MethodCut[i].Method)
			if err != nil {
				return o, fmt.Errorf("cutmethod %s: %w", name, err)
			}
			if o.CutMethods == nil {
				o.CutMethods = map[int]bool{}
			}
			o.CutMethods[m.ID] = true
		}
	}
	return o, nil
}
