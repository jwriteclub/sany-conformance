// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

// Analysis holds the product world and every scope's result.
type Analysis struct {
	Paths   Paths
	Cfg     *Config
	World   *World
	Results map[string]map[Precision]*Result // scope, precision
	Specs   *Coverage
	Union   *Coverage
	SpecCLI *Coverage
	// ProbesSpecs and ProbesUnion hold JaCoCo's per-instruction arms for the two reports.
	ProbesSpecs ProbeTable
	ProbesUnion ProbeTable
	worlds      map[string]*World // harness worlds
	// NtkAll treats the jj_ntk ternary as constant in every scope; `staticreach view`.
	// Off for the written report, where only a scope naming the view uses it.
	NtkAll bool
}

var precisions = []Precision{CHA, RTA, RTAConst}

// NewAnalysis loads the product and runs every scope at every precision.
func NewAnalysis(p Paths, cfg *Config) (*Analysis, error) {
	w, err := NewWorld([]string{p.Classes()}, p.JDK())
	if err != nil {
		return nil, err
	}
	a := &Analysis{Paths: p, Cfg: cfg, World: w, Results: map[string]map[Precision]*Result{}, worlds: map[string]*World{}}
	for i := range cfg.Scopes {
		s := &cfg.Scopes[i]
		a.Results[s.Name] = map[Precision]*Result{}
		for _, prec := range precisions {
			o, err := scopeOptions(cfg, w, s, prec)
			if err != nil {
				return nil, err
			}
			a.Results[s.Name][prec] = Analyse(w, o)
		}
	}
	if a.Specs, err = ReadCoverage(p.SpecsXML()); err != nil {
		return nil, fmt.Errorf("%w (run: staticreach prepare)", err)
	}
	if a.Union, err = ReadCoverage(p.UnionXML()); err != nil {
		return nil, err
	}
	if a.SpecCLI, err = ReadCoverage(p.abs("data/staticreach/specs-cli.xml.gz")); err != nil {
		return nil, fmt.Errorf("%w (run: staticreach prepare)", err)
	}
	if a.ProbesSpecs, err = ReadProbes(p.abs("data/staticreach/branches-specs.tsv.gz")); err != nil {
		return nil, fmt.Errorf("%w (run: staticreach prepare)", err)
	}
	if a.ProbesUnion, err = ReadProbes(p.abs("data/staticreach/branches-union.tsv.gz")); err != nil {
		return nil, fmt.Errorf("%w (run: staticreach prepare)", err)
	}
	if err := a.ProbesSpecs.CheckAgainst(a.Specs); err != nil {
		return nil, fmt.Errorf("per-instruction data disagrees with the report, rerun prepare: %w", err)
	}
	if err := a.ProbesUnion.CheckAgainst(a.Union); err != nil {
		return nil, fmt.Errorf("per-instruction data disagrees with the report, rerun prepare: %w", err)
	}
	return a, nil
}

// Close releases the worlds.
func (a *Analysis) Close() {
	a.World.Close()
	for _, w := range a.worlds {
		w.Close()
	}
}

// harnessWorld loads the product plus a harness's classes.
func (a *Analysis) harnessWorld(h *HarnessDef) (*World, error) {
	if len(h.Dirs) == 0 {
		return a.World, nil
	}
	if w, ok := a.worlds[h.Name]; ok {
		return w, nil
	}
	dirs := []string{a.Paths.Classes()}
	for _, d := range h.Dirs {
		dirs = append(dirs, a.Paths.abs(d))
	}
	w, err := NewWorld(dirs, a.Paths.JDK())
	if err != nil {
		return nil, err
	}
	a.worlds[h.Name] = w
	return w, nil
}

// harnessRoots picks the classes a harness calls into from outside.
func harnessRoots(w *World, h *HarnessDef) []*Class {
	var out []*Class
	for _, name := range w.AppNames {
		c := w.Classes[name]
		if c.Origin == 0 {
			continue
		}
		for _, r := range h.Roots {
			if r == "*" || r == name {
				out = append(out, c)
				break
			}
		}
	}
	return out
}

// Violation is a method that ran but that the analysis says cannot.
type Violation struct {
	Run       string
	Precision Precision
	Method    string
	Branches  int
	Why       string
}

// DeadLine is a source line ruled out yet executed.
type DeadLine struct {
	Run    string
	Source string
	Line   int
	Detail string
}

// RunResult is one run compared against its scope at each precision.
type RunResult struct {
	Run      RunDef
	Path     string
	Ran      int
	Outside  map[Precision][]Violation
	DeadHits []DeadLine
}

// SoundnessReport is the whole comparison.
type SoundnessReport struct {
	Runs        []*RunResult
	Unexplained int
	Explained   int
	Errors      []string
}

// CheckSoundness compares every configured run with its scope.
func (a *Analysis) CheckSoundness() *SoundnessReport {
	rep := &SoundnessReport{}
	explained := map[string]string{}
	for _, e := range a.Cfg.Explains {
		explained[e.Kind+" "+e.Method] = e.Why
	}
	type key struct {
		scope, harness, extra string
		prec                  Precision
	}
	cache := map[key]*Result{}
	for _, run := range a.Cfg.Runs {
		paths, err := filepath.Glob(a.Paths.abs(run.Glob))
		if err != nil {
			rep.Errors = append(rep.Errors, err.Error())
			continue
		}
		if len(paths) == 0 {
			rep.Errors = append(rep.Errors, "no reports match "+run.Glob)
			continue
		}
		sort.Strings(paths)
		h := a.Cfg.Harness(run.Harness)
		s := a.Cfg.Scope(run.Scope)
		w, err := a.harnessWorld(h)
		if err != nil {
			fatalf("%v", err)
		}
		for _, path := range paths {
			cov, err := ReadCoverage(path)
			if err != nil {
				rep.Errors = append(rep.Errors, err.Error())
				continue
			}
			rel, _ := filepath.Rel(a.Paths.Root, path)
			rr := &RunResult{Run: run, Path: rel, Outside: map[Precision][]Violation{}}
			for _, prec := range precisions {
				k := key{run.Scope, run.Harness, strings.Join(run.ExtraCuts, ","), prec}
				res := cache[k]
				if res == nil {
					o, err := scopeOptions(a.Cfg, w, s, prec)
					if err != nil {
						fatalf("%v", err)
					}
					o.RootClass = harnessRoots(w, h)
					res = Analyse(w, o)
					cache[k] = res
				}
				for _, mk := range sortedKeys(cov) {
					mc := cov.Methods[mk]
					if mc.InsnHit == 0 {
						continue
					}
					if prec == CHA {
						rr.Ran++
					}
					m, err := w.MethodByRef(mk)
					if err != nil {
						rep.Errors = append(rep.Errors, fmt.Sprintf("%s: %v", rel, err))
						continue
					}
					if res.Reached[m.ID] {
						continue
					}
					v := Violation{Run: rel, Precision: prec, Method: mk, Branches: mc.BranchTotal}
					if why, ok := explained[run.Kind+" "+mk]; ok {
						v.Why = why
						rep.Explained++
					} else {
						rep.Unexplained++
					}
					rr.Outside[prec] = append(rr.Outside[prec], v)
				}
			}
			rr.DeadHits = a.deadLineHits(w, s, cov, rel)
			rep.Unexplained += len(rr.DeadHits)
			// A method cut is a claim about every report, whatever its scope.
			for i := range a.Cfg.MethodCut {
				mc := &a.Cfg.MethodCut[i]
				if c := cov.Methods[mc.Method]; c != nil && c.InsnHit > 0 {
					rep.Errors = append(rep.Errors, fmt.Sprintf("%s ran %s (%d instructions): cut %s does not hold", rel, mc.Method, c.InsnHit, mc.Name))
				}
			}
			rep.Runs = append(rep.Runs, rr)
		}
	}
	a.checkLexer(rep)
	a.checkMethodCuts(rep)
	for _, site := range a.ReflectionSites() {
		if a.Cfg.reflectFor(site.Caller) == nil {
			rep.Errors = append(rep.Errors, fmt.Sprintf("unresolved reflective site in reached code: %s line %d: %s", site.Caller, site.Line, site.API))
		}
	}
	rep.Unexplained += len(rep.Errors)
	return rep
}

func sortedKeys(cov *Coverage) []string {
	out := append([]string(nil), cov.Keys...)
	sort.Strings(out)
	return out
}

// deadLineHits tests the intra-method refinement and the cuts against a run: a
// source line whose every instruction is ruled out must not have run, and a line
// that lost a branch outcome cannot have hit more than the rest.
func (a *Analysis) deadLineHits(w *World, s *ScopeDef, cov *Coverage, run string) []DeadLine {
	o, err := scopeOptions(a.Cfg, w, s, RTAConst)
	if err != nil {
		fatalf("%v", err)
	}
	type lineState struct{ dead, alive, removed int }
	files := map[string]map[int]*lineState{}
	var names []string
	state := func(src string, line int) *lineState {
		f := files[src]
		if f == nil {
			f = map[int]*lineState{}
			files[src] = f
			names = append(names, src)
		}
		st := f[line]
		if st == nil {
			st = &lineState{}
			f[line] = st
		}
		return st
	}
	inCut := func(c *Class, line int) bool {
		for i := range o.Cuts {
			if o.Cuts[i].Class == c.Name && line >= o.Cuts[i].From && line <= o.Cuts[i].To {
				return true
			}
		}
		return false
	}
	for _, m := range w.Methods {
		pkg := m.Class.Name
		if i := strings.LastIndexByte(pkg, '/'); i >= 0 {
			pkg = pkg[:i]
		}
		src := pkg + "/" + m.Class.SourceFile
		if _, ok := cov.Lines[src]; !ok {
			continue
		}
		insns, err := m.Decode()
		if err != nil {
			fatalf("%v", err)
		}
		d := w.dead(m)
		for _, in := range insns {
			line := m.LineAt(in.Off)
			st := state(src, line)
			if d.offsets[in.Off] || inCut(m.Class, line) {
				st.dead++
			} else {
				st.alive++
			}
		}
		for _, lo := range d.removed {
			state(src, lo.Line).removed += lo.Outcomes
		}
		for _, lo := range CutRemoved(m, o.Cuts) {
			state(src, lo.Line).removed += lo.Outcomes
		}
	}
	sort.Strings(names)
	var out []DeadLine
	for _, src := range names {
		lines := make([]int, 0, len(files[src]))
		for l := range files[src] {
			lines = append(lines, l)
		}
		sort.Ints(lines)
		for _, l := range lines {
			st := files[src][l]
			jc, ok := cov.Lines[src][l]
			if !ok {
				continue
			}
			if st.alive == 0 && st.dead > 0 && jc.InsnHit > 0 {
				out = append(out, DeadLine{run, src, l, fmt.Sprintf("line ruled out but ran (%d instructions hit)", jc.InsnHit)})
				continue
			}
			total := jc.BranchHit + jc.BranchMissed
			if st.removed > 0 && total > 0 {
				rem := st.removed
				if rem > total {
					rem = total
				}
				if jc.BranchHit > total-rem {
					out = append(out, DeadLine{run, src, l, fmt.Sprintf("%d of %d outcomes ruled out but %d were hit", rem, total, jc.BranchHit)})
				}
			}
		}
	}
	return out
}

// Text renders the report for the terminal.
func (r *SoundnessReport) Text() string {
	var sb strings.Builder
	for _, rr := range r.Runs {
		fmt.Fprintf(&sb, "%-34s kind=%-4s scope=%-16s harness=%-5s ran=%-5d", rr.Path, rr.Run.Kind, rr.Run.Label(), rr.Run.Harness, rr.Ran)
		for _, p := range precisions {
			fmt.Fprintf(&sb, " %s:%d", p, len(rr.Outside[p]))
		}
		fmt.Fprintf(&sb, " deadlines:%d\n", len(rr.DeadHits))
		for _, p := range precisions {
			for _, v := range rr.Outside[p] {
				tag := "UNEXPLAINED"
				if v.Why != "" {
					tag = "explained"
				}
				fmt.Fprintf(&sb, "  %s %s %s %s\n", tag, p, v.Method, v.Why)
			}
		}
		for _, d := range rr.DeadHits {
			fmt.Fprintf(&sb, "  UNEXPLAINED dead line %s:%d: %s\n", d.Source, d.Line, d.Detail)
		}
	}
	for _, e := range r.Errors {
		fmt.Fprintf(&sb, "ERROR %s\n", e)
	}
	fmt.Fprintf(&sb, "unexplained %d, explained %d\n", r.Unexplained, r.Explained)
	return sb.String()
}
