// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"sort"
	"strings"
)

// Tally is branch counts for a set of methods.
type Tally struct {
	Methods  int
	Branches int
	Specs    int
	Union    int
	// ExclSpecs and ExclUnion are the excludable missed guard arms (no refinement applied).
	ExclSpecs GuardMissed
	ExclUnion GuardMissed
}

// Pct formats a percentage with one decimal, stable across runs.
func Pct(hit, total int) float64 {
	if total == 0 {
		return 0
	}
	return 100 * float64(hit) / float64(total)
}

// ScopeTally splits one scope and precision into handwritten and generated.
type ScopeTally struct {
	Hand, Gen Tally
	// Removed counts outcomes that cuts and refinement took from reached methods.
	RemovedHand, RemovedGen int
	// View says that the jj_ntk ternary was treated as constant; the Ntk counts are
	// what that took out of Gen, in outcomes.
	View                      bool
	NtkTotal, NtkSpecs, NtkUn int
}

func (a *Analysis) methodOf(key string) *Method {
	m, err := a.World.MethodByRef(key)
	if err != nil {
		fatalf("%s is in a report but not in the class files: %v", key, err)
	}
	return m
}

// removedLines gathers the outcomes ruled out in a method, per source line.
func removedLines(w *World, m *Method, o Options) map[int]int {
	var out map[int]int
	add := func(ls []LineOutcome) {
		for _, lo := range ls {
			if out == nil {
				out = map[int]int{}
			}
			out[lo.Line] += lo.Outcomes
		}
	}
	add(CutRemoved(m, o.Cuts))
	if o.Precision == RTAConst {
		add(w.dead(m).removed)
	}
	return out
}

// adjust applies removed outcomes to a method's counters, in JaCoCo's units: a
// line cannot lose more than it counts, and cannot keep more hits than it has
// outcomes left.
func adjust(cov *Coverage, source string, total, hit int, removed map[int]int) (int, int, int) {
	gone := 0
	for line, rem := range removed {
		lc, ok := cov.Lines[source][line]
		if !ok {
			continue
		}
		t := lc.BranchHit + lc.BranchMissed
		if rem > t {
			rem = t
		}
		gone += rem
		if over := lc.BranchHit - (t - rem); over > 0 {
			hit -= over
		}
	}
	return total - gone, hit, gone
}

// TallyScope computes one scope at one precision.
func (a *Analysis) TallyScope(scope string, prec Precision) ScopeTally {
	def := a.Cfg.Scope(scope)
	return a.tallyScopeView(scope, prec, a.NtkAll || def != nil && def.View == ViewNtkConstant)
}

// tallyScopeView computes one scope at one precision, with the ternary view on or
// off. The view only moves numbers: Reached, and so the map, is the same either way.
func (a *Analysis) tallyScopeView(scope string, prec Precision, view bool) ScopeTally {
	res := a.Results[scope][prec]
	st := ScopeTally{View: view}
	for _, k := range sortedKeys(a.Union) {
		m := a.methodOf(k)
		if !res.Reached[m.ID] {
			continue
		}
		u := a.Union.Methods[k]
		s := a.Specs.Methods[k]
		if s == nil {
			fatalf("%s missing from the spec report", k)
		}
		removed := removedLines(a.World, m, res.Opts)
		total, hitU, gone := adjust(a.Union, u.Source, u.BranchTotal, u.BranchHit, removed)
		_, hitS, _ := adjust(a.Specs, s.Source, s.BranchTotal, s.BranchHit, removed)
		t := &st.Hand
		if IsGenerated(m.Class) {
			t = &st.Gen
			st.RemovedGen += gone
		} else {
			st.RemovedHand += gone
		}
		tallyS := map[string]*GuardTally{}
		tallyU := map[string]*GuardTally{}
		gs, err := a.World.MethodGuards(m, a.ProbesSpecs[k], tallyS)
		if err != nil {
			fatalf("%v", err)
		}
		gu, err := a.World.MethodGuards(m, a.ProbesUnion[k], tallyU)
		if err != nil {
			fatalf("%v", err)
		}
		if ns, nu := tallyS[GuardNtk], tallyU[GuardNtk]; view && ns != nil && nu != nil {
			// Per instruction: both outcomes of each ternary leave the denominator and
			// each outcome that was hit leaves the numerator, so the guard numbers must
			// not exclude them a second time.
			total -= nu.Total
			hitS -= ns.Hit
			hitU -= nu.Hit
			gs.Partly -= ns.Excl.Partly
			gs.Never -= ns.Excl.Never
			gu.Partly -= nu.Excl.Partly
			gu.Never -= nu.Excl.Never
			st.NtkTotal += nu.Total
			st.NtkSpecs += ns.Hit
			st.NtkUn += nu.Hit
		}
		t.Methods++
		t.Branches += total
		t.Specs += hitS
		t.Union += hitU
		t.ExclSpecs.Add(gs)
		t.ExclUnion.Add(gu)
	}
	return st
}

// PackageOutside is branches outside a scope in one package.
type PackageOutside struct {
	Package string
	Hand    int
	Gen     int
	Hit     int
}

// OutsideByPackage lists what a scope leaves out, by package.
func (a *Analysis) OutsideByPackage(scope string, prec Precision) []PackageOutside {
	res := a.Results[scope][prec]
	by := map[string]*PackageOutside{}
	for _, k := range sortedKeys(a.Union) {
		m := a.methodOf(k)
		if res.Reached[m.ID] {
			continue
		}
		u := a.Union.Methods[k]
		if u.BranchTotal == 0 {
			continue
		}
		pkg := m.Class.Name[:strings.LastIndexByte(m.Class.Name, '/')]
		p := by[pkg]
		if p == nil {
			p = &PackageOutside{Package: pkg}
			by[pkg] = p
		}
		if IsGenerated(m.Class) {
			p.Gen += u.BranchTotal
		} else {
			p.Hand += u.BranchTotal
		}
		p.Hit += u.BranchHit
	}
	out := make([]PackageOutside, 0, len(by))
	for _, p := range by {
		out = append(out, *p)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Package < out[j].Package })
	return out
}

// ClassOutside is one class's branches outside a scope.
type ClassOutside struct {
	Class  string
	Branch int
	Hit    int
}

// OutsideByClass lists the largest classes a scope leaves out.
func (a *Analysis) OutsideByClass(scope string, prec Precision) []ClassOutside {
	res := a.Results[scope][prec]
	by := map[string]*ClassOutside{}
	for _, k := range sortedKeys(a.Union) {
		m := a.methodOf(k)
		u := a.Union.Methods[k]
		if res.Reached[m.ID] || u.BranchTotal == 0 {
			continue
		}
		p := by[m.Class.Name]
		if p == nil {
			p = &ClassOutside{Class: m.Class.Name}
			by[m.Class.Name] = p
		}
		p.Branch += u.BranchTotal
		p.Hit += u.BranchHit
	}
	out := make([]ClassOutside, 0, len(by))
	for _, p := range by {
		out = append(out, *p)
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Branch != out[j].Branch {
			return out[i].Branch > out[j].Branch
		}
		return out[i].Class < out[j].Class
	})
	return out
}

// ReflectionAPI names the JDK members that can run code the call graph cannot see.
var ReflectionAPI = map[string]map[string]bool{
	"java/lang/Class": {
		"forName": true, "newInstance": true, "getMethod": true, "getMethods": true, "getDeclaredMethod": true,
		"getDeclaredMethods": true, "getConstructor": true, "getConstructors": true, "getDeclaredConstructor": true,
		"getDeclaredConstructors": true, "getField": true, "getFields": true, "getDeclaredField": true, "getDeclaredFields": true,
	},
	"java/lang/ClassLoader":                 {"loadClass": true, "defineClass": true},
	"java/lang/reflect/Method":              {"invoke": true},
	"java/lang/reflect/Constructor":         {"newInstance": true},
	"java/lang/reflect/Field":               {"get": true, "set": true, "setBoolean": true, "setInt": true, "setAccessible": true},
	"java/util/ServiceLoader":               {"load": true, "loadInstalled": true},
	"java/lang/reflect/Proxy":               {"newProxyInstance": true},
	"java/lang/invoke/MethodHandles$Lookup": {"findVirtual": true, "findStatic": true, "findConstructor": true, "findSpecial": true, "unreflect": true},
	"java/io/ObjectInputStream":             {"readObject": true, "readUnshared": true},
	"java/io/ObjectOutputStream":            {"writeObject": true, "writeUnshared": true},
	"java/lang/System":                      {"load": true, "loadLibrary": true},
	"java/lang/Runtime":                     {"load": true, "loadLibrary": true},
}

// ReflectionSite is one reflective call in reached code.
type ReflectionSite struct {
	Caller string
	API    string
	Line   int
	// Scopes names the scopes that reach it, at RTA, in configuration order.
	Scopes []string
}

// ReflectionSites scans reached code, at the CHA upper bound, for reflective calls.
func (a *Analysis) ReflectionSites() []ReflectionSite {
	by := map[string]*ReflectionSite{}
	for i := range a.Cfg.Scopes {
		s := a.Cfg.Scopes[i].Name
		res := a.Results[s][CHA]
		for _, m := range a.World.Methods {
			if !res.Reached[m.ID] {
				continue
			}
			events, err := m.Events()
			if err != nil {
				fatalf("%v", err)
			}
			for j := range events {
				ev := &events[j]
				if ev.Kind != EvInvoke && ev.Kind != EvHandle && ev.Kind != EvLambda {
					continue
				}
				if !ReflectionAPI[ev.Ref.Class][ev.Ref.Name] || res.cutAway(m, ev.Line) {
					continue
				}
				key := m.Key() + " " + ev.Ref.Class + "." + ev.Ref.Name + " " + itoa(ev.Line)
				rs := by[key]
				if rs == nil {
					rs = &ReflectionSite{Caller: m.Key(), API: ev.Ref.Class + "." + ev.Ref.Name, Line: ev.Line}
					by[key] = rs
				}
				rs.Scopes = append(rs.Scopes, s)
			}
		}
	}
	keys := make([]string, 0, len(by))
	for k := range by {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	out := make([]ReflectionSite, 0, len(keys))
	for _, k := range keys {
		out = append(out, *by[k])
	}
	return out
}

// ExplorerSite is a reference into tla2sany/explorer from outside the package.
type ExplorerSite struct {
	Caller string
	Target string
	Line   int
	Kind   string
}

// ExplorerReferences lists every site outside the explorer package that creates,
// calls into or takes a method handle on an explorer class, and returns the
// creators of Explorer itself.
func (a *Analysis) ExplorerReferences() (all []ExplorerSite, creators []ExplorerSite) {
	const pkg = "tla2sany/explorer/"
	for _, m := range a.World.Methods {
		events, err := m.Events()
		if err != nil {
			fatalf("%v", err)
		}
		for i := range events {
			ev := &events[i]
			if !strings.HasPrefix(ev.Ref.Class, pkg) {
				continue
			}
			kind := "call"
			switch ev.Kind {
			case EvNew:
				kind = "new"
			case EvStatic:
				kind = "static"
			case EvLambda, EvHandle:
				kind = "handle"
			}
			site := ExplorerSite{Caller: m.Key(), Target: ev.Ref.Class + "." + ev.Ref.Name, Line: ev.Line, Kind: kind}
			if ev.Kind == EvNew && ev.Ref.Class == pkg+"Explorer" {
				creators = append(creators, site)
			}
			if !strings.HasPrefix(m.Class.Name, pkg) {
				all = append(all, site)
			}
		}
	}
	return all, creators
}

func itoa(n int) string {
	var b [20]byte
	i := len(b)
	neg := n < 0
	if neg {
		n = -n
	}
	for {
		i--
		b[i] = byte('0' + n%10)
		n /= 10
		if n == 0 {
			break
		}
	}
	if neg {
		i--
		b[i] = '-'
	}
	return string(b[i:])
}
