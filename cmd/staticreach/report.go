// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"crypto/sha256"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// prototypeViolations are the methods the Python prototype found outside its
// scope, kept so that the report can say what happened to each.
var prototypeViolations = []string{
	"tla2sany/semantic/BuiltInOperators$BuiltInOperator.<init>",
	"tla2sany/semantic/BuiltInOperators.<clinit>",
	"tla2sany/semantic/BuiltInOperators.mk",
	"tla2sany/semantic/ErrorCode$ErrorLevel.<clinit>",
	"tla2sany/semantic/ErrorCode.getStandardValue",
	"tla2sany/semantic/Errors$ErrorDetails.getLocation",
	"tla2sany/semantic/Errors.getMessages",
	"tla2sany/parser/JunctionListContext$JunctionListType.<clinit>",
	"tla2sany/drivers/SanySettings.defaultSettings",
	"tla2sany/modanalyzer/SpecObj.getName",
	"tla2sany/modanalyzer/SpecObj.getParseErrors",
	"tla2sany/parser/SyntaxTreeNode.getHeirs",
	"tla2sany/xml/XMLExportable.export",
	"tla2sany/xml/XMLExportable.lambda$export$0",
}

const tsvHeader = "# Copyright 2026 Christopher O'Connell\n# SPDX-License-Identifier: Apache-2.0\n"

// WriteOutputs writes the map, the branch table and the report.
func (a *Analysis) WriteOutputs(dir string, rep *SoundnessReport) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	if err := writeFile(filepath.Join(dir, "candidate-map.tsv"), a.mapTSV()); err != nil {
		return err
	}
	if err := writeFile(filepath.Join(dir, "candidate-map-branches.tsv"), a.branchTSV()); err != nil {
		return err
	}
	md, err := a.reportMD(rep)
	if err != nil {
		return err
	}
	if err := writeFile(filepath.Join(dir, "ntk-sites.tsv"), a.ntkSitesTSV()); err != nil {
		return err
	}
	if err := writeFile(filepath.Join(dir, "port-cut-arms.tsv"), a.portCutArmsTSV()); err != nil {
		return err
	}
	return writeFile(filepath.Join(dir, "candidate-map.md"), md)
}

func writeFile(path, text string) error { return os.WriteFile(path, []byte(text), 0o644) }

// membership is R when RTA reaches the method, C when only CHA does, - otherwise.
func (a *Analysis) membership(scope string, m *Method) string {
	switch {
	case a.Results[scope][RTA].Reached[m.ID]:
		return "R"
	case a.Results[scope][CHA].Reached[m.ID]:
		return "C"
	}
	return "-"
}

func (a *Analysis) inAnyScope(m *Method) bool {
	for i := range a.Cfg.Scopes {
		if a.membership(a.Cfg.Scopes[i].Name, m) != "-" {
			return true
		}
	}
	return false
}

// bestChain is the shortest chain in the first scope that reaches the method.
func (a *Analysis) bestChain(m *Method) (string, int, []*Method) {
	for _, p := range []Precision{RTA, CHA} {
		for i := range a.Cfg.Scopes {
			s := a.Cfg.Scopes[i].Name
			r := a.Results[s][p]
			if r.Reached[m.ID] {
				return s, r.Depth(m), r.Chain(m)
			}
		}
	}
	return "", -1, nil
}

// short renders a method with its simple class name, adding the descriptor only
// when the class has an overload that the name alone would blur.
func short(m *Method) string {
	name := m.Class.Name
	if i := strings.LastIndexByte(name, '/'); i >= 0 {
		name = name[i+1:]
	}
	n := 0
	for _, o := range m.Class.Methods {
		if o.Name == m.Name {
			n++
		}
	}
	if n > 1 {
		return name + "." + m.Name + m.Desc
	}
	return name + "." + m.Name
}

func chainText(chain []*Method) string {
	parts := make([]string, len(chain))
	for i, m := range chain {
		parts[i] = short(m)
	}
	return strings.Join(parts, " > ")
}

func (a *Analysis) mapTSV() string {
	var sb strings.Builder
	sb.WriteString(tsvHeader)
	sb.WriteString("class\tname\tdescriptor\tkind\tfirst_line")
	for i := range a.Cfg.Scopes {
		sb.WriteString("\t" + a.Cfg.Scopes[i].Name)
	}
	sb.WriteString("\tchain_scope\tdepth\tchain\tbranches\thit_specs\thit_union\tguard_missed_specs_partly\tguard_missed_specs_never\tguard_missed_union_partly\tguard_missed_union_never\n")
	ms := append([]*Method(nil), a.World.Methods...)
	sort.Slice(ms, func(i, j int) bool { return ms[i].Key() < ms[j].Key() })
	for _, m := range ms {
		inSany := strings.HasPrefix(m.Class.Name, "tla2sany/")
		if !inSany && !a.inAnyScope(m) {
			continue
		}
		kind := "other"
		if inSany {
			kind = "handwritten"
			if IsGenerated(m.Class) {
				kind = "generated"
			}
		}
		k := m.Key()
		u := a.Union.Methods[k]
		s := a.Specs.Methods[k]
		first := "-"
		if u != nil && u.FirstLine > 0 {
			first = strconv.Itoa(u.FirstLine)
		}
		fmt.Fprintf(&sb, "%s\t%s\t%s\t%s\t%s", m.Class.Name, m.Name, m.Desc, kind, first)
		for i := range a.Cfg.Scopes {
			sb.WriteString("\t" + a.membership(a.Cfg.Scopes[i].Name, m))
		}
		scope, depth, chain := a.bestChain(m)
		if scope == "" {
			sb.WriteString("\t-\t-\t-")
		} else {
			fmt.Fprintf(&sb, "\t%s\t%d\t%s", scope, depth, chainText(chain))
		}
		if u == nil || s == nil {
			sb.WriteString("\t-\t-\t-\t-\t-\t-\t-\n")
			continue
		}
		gs, err := a.World.MethodGuards(m, a.ProbesSpecs[k], map[string]*GuardTally{})
		if err != nil {
			fatalf("%v", err)
		}
		gu, err := a.World.MethodGuards(m, a.ProbesUnion[k], map[string]*GuardTally{})
		if err != nil {
			fatalf("%v", err)
		}
		fmt.Fprintf(&sb, "\t%d\t%d\t%d\t%d\t%d\t%d\t%d\n", u.BranchTotal, s.BranchHit, u.BranchHit, gs.Partly, gs.Never, gu.Partly, gu.Never)
	}
	return sb.String()
}

// armList names the missed arms. A switch's non-default arms are only counted:
// JaCoCo's numbering of them depends on probe placement.
func armList(s *BranchShape, r *ProbeRow) string {
	var parts []string
	if s.Op == opTableswi || s.Op == opLookupswi {
		other := r.Total - r.HitCount()
		if !r.Has(0) {
			other--
			if s.Plain {
				parts = append(parts, "default")
			} else {
				other++
			}
		}
		if other > 0 {
			parts = append(parts, strconv.Itoa(other)+"-other")
		}
	} else {
		for arm := 0; arm < r.Total; arm++ {
			if !r.Has(arm) {
				parts = append(parts, s.ArmName(arm))
			}
		}
	}
	if len(parts) == 0 {
		return "-"
	}
	return strings.Join(parts, ",")
}

// branchTSV is the per-instruction table: which arm was missed, by shape.
func (a *Analysis) branchTSV() string {
	var sb strings.Builder
	sb.WriteString(tsvHeader)
	sb.WriteString("class\tname\tdescriptor\tline\toffset\tkind")
	for i := range a.Cfg.Scopes {
		sb.WriteString("\t" + a.Cfg.Scopes[i].Name)
	}
	sb.WriteString("\tshape\tarms\tmissed_specs\tmissed_union\n")
	for _, k := range sortedKeys(a.Union) {
		specs := a.ProbesSpecs[k]
		union := a.ProbesUnion[k]
		if len(union) == 0 {
			continue
		}
		m := a.methodOf(k)
		// Fills in each shape's Plain flag before its arms are named.
		if _, err := a.World.MethodGuards(m, union, map[string]*GuardTally{}); err != nil {
			fatalf("%v", err)
		}
		shapes, err := a.World.ClassifyMethod(m)
		if err != nil {
			fatalf("%v", err)
		}
		kind := "handwritten"
		if IsGenerated(m.Class) {
			kind = "generated"
		}
		for i, r := range union {
			s := shapes[r.Ordinal]
			fmt.Fprintf(&sb, "%s\t%s\t%s\t%d\t%d\t%s", m.Class.Name, m.Name, m.Desc, r.Line, s.Off, kind)
			for j := range a.Cfg.Scopes {
				sb.WriteString("\t" + a.membership(a.Cfg.Scopes[j].Name, m))
			}
			fmt.Fprintf(&sb, "\t%s\t%d\t%s\t%s\n", s.Class, r.Total, armList(s, specs[i]), armList(s, r))
		}
	}
	return sb.String()
}

func p1(hit, total int) string { return fmt.Sprintf("%.1f%%", Pct(hit, total)) }

func (a *Analysis) reportMD(rep *SoundnessReport) (string, error) {
	var b strings.Builder
	cfg := a.Cfg
	primary := cfg.Scopes[0].Name
	tallies := map[string]map[Precision]ScopeTally{}
	for i := range cfg.Scopes {
		n := cfg.Scopes[i].Name
		tallies[n] = map[Precision]ScopeTally{}
		for _, p := range precisions {
			tallies[n][p] = a.TallyScope(n, p)
		}
	}
	pt := tallies[primary]
	allHand, allGen := a.allBranches()

	b.WriteString("<!-- Copyright 2026 Christopher O'Connell -->\n<!-- SPDX-License-Identifier: Apache-2.0 -->\n<!-- Generated by cmd/staticreach. Do not edit. -->\n\n")
	b.WriteString("# Candidate coverage map for SANY\n\n")
	rta := pt[RTA]
	cha := pt[CHA]
	fmt.Fprintf(&b, "Primary scope `%s` (the SANY command line, Explorer cut, XMLExporter excluded): **%d of %d handwritten branches are candidates under RTA (%d under CHA); specs alone hit %s of them, the union %s.** Generated code: %d of %d candidates, specs %s, union %s. For comparison, against all %d handwritten branches the union is %s.\n\n",
		primary, rta.Hand.Branches, allHand, cha.Hand.Branches, p1(rta.Hand.Specs, rta.Hand.Branches), p1(rta.Hand.Union, rta.Hand.Branches),
		rta.Gen.Branches, allGen, p1(rta.Gen.Specs, rta.Gen.Branches), p1(rta.Gen.Union, rta.Gen.Branches), allHand, p1(a.handHit(a.Union), allHand))
	b.WriteString("**What this is.** Static call-graph reachability over the pinned class files: an upper bound on what can run. A branch in this map has not been ruled out; nothing here claims it can be taken, and whether it can is undecidable in general. A branch outside the map cannot run from these entry points, given the assumptions under \"Limits\". No model, no network, no clock: the same input gives the same bytes.\n\n")
	b.WriteString("**Rerun, from the repository root:** `go run ./cmd/staticreach -jdk $JAVA_HOME`, after `scripts/setup reference` and `scripts/maintainer/unpack-classes`.\n\n")
	b.WriteString("It reads the JaCoCo reports committed in `data/staticreach`, starts no JVM and runs no SANY code, writes this report plus `candidate-map.tsv` and `candidate-map-branches.tsv` under `.work/staticreach`, then runs the soundness check and exits 1 on anything unexplained. The check here covers 2 report rows, both of the plain command; the measurement it was first run on covered 14, with the measuring wrapper and upstream's test classes on the path, neither of which is committed here. Scopes, cuts and run-to-scope assignments are in `cmd/staticreach/scopes.conf`; `-scopes` reads another file.\n\n")
	b.WriteString("## Numbers by scope and precision\n\nBranches are JaCoCo's. \"specs\" is corpus A+B+C plus kept tiers D and E; \"union\" adds upstream's tests (`data/staticreach/union.xml.gz`). Percentages are of the branches in scope.\n\n")
	b.WriteString("| scope | precision | handwritten branches | specs | union | generated branches | specs | union |\n|---|---|---:|---:|---:|---:|---:|---:|\n")
	for i := range cfg.Scopes {
		n := cfg.Scopes[i].Name
		for _, p := range precisions {
			t := tallies[n][p]
			label := "`" + n + "`"
			if i == 0 {
				label = "**`" + n + "`**"
			}
			if t.View {
				label += " (ternary constant)"
			}
			fmt.Fprintf(&b, "| %s | %s | %d | %s | %s | %d | %s | %s |\n", label, p, t.Hand.Branches, p1(t.Hand.Specs, t.Hand.Branches), p1(t.Hand.Union, t.Hand.Branches), t.Gen.Branches, p1(t.Gen.Specs, t.Gen.Branches), p1(t.Gen.Union, t.Gen.Branches))
		}
	}
	b.WriteString("\n")
	fmt.Fprintf(&b, "Definitions. CHA dispatches a virtual call on an application type to every concrete subtype. RTA dispatches only to classes instantiated in reachable code. JDK callbacks count under both, and only for classes instantiated in reachable code: an unrestricted `Object.toString` would otherwise pull in all %d classes and make the Explorer cut meaningless. RTA+const adds the one intra-method refinement below.\n\n", len(a.World.AppNames))
	b.WriteString("Scopes:\n\n")
	for i := range cfg.Scopes {
		fmt.Fprintf(&b, "- `%s`: %s\n", cfg.Scopes[i].Name, cfg.Scopes[i].Why)
	}
	b.WriteString("\n")

	if err := a.writePort(&b, rep); err != nil {
		return "", err
	}
	a.writeExporter(&b, tallies)
	a.writeGuards(&b, tallies)
	a.writeExplorer(&b)
	a.writeFlags(&b)
	a.writeRefinement(&b, tallies)
	a.writeSoundness(&b, rep)
	a.writeCallbacks(&b)
	a.writeReflection(&b)
	a.writeOutside(&b)
	if err := a.writeLexer(&b); err != nil {
		return "", err
	}
	a.writeLimits(&b)
	return b.String(), nil
}

func (a *Analysis) allBranches() (int, int) {
	h, g := 0, 0
	for _, k := range sortedKeys(a.Union) {
		if IsGenerated(a.methodOf(k).Class) {
			g += a.Union.Methods[k].BranchTotal
		} else {
			h += a.Union.Methods[k].BranchTotal
		}
	}
	return h, g
}

func (a *Analysis) handHit(cov *Coverage) int {
	n := 0
	for _, k := range sortedKeys(cov) {
		if !IsGenerated(a.methodOf(k).Class) {
			n += cov.Methods[k].BranchHit
		}
	}
	return n
}

func (a *Analysis) writeExporter(b *strings.Builder, tallies map[string]map[Precision]ScopeTally) {
	primary := a.Cfg.Scopes[0].Name
	res := a.Results[primary][RTA]
	onlyHand, onlyGen, cliAhead := 0, 0, 0
	for _, k := range sortedKeys(a.Specs) {
		m := a.methodOf(k)
		if !res.Reached[m.ID] {
			continue
		}
		s := a.Specs.Methods[k]
		c := a.SpecCLI.Methods[k]
		if c == nil {
			continue
		}
		d := s.BranchHit - c.BranchHit
		switch {
		case d < 0:
			cliAhead -= d
		case IsGenerated(m.Class):
			onlyGen += d
		default:
			onlyHand += d
		}
	}
	t := tallies[primary][RTA]
	x := tallies["cli+xml-noexp"][RTA]
	b.WriteString("## The exporter is secondary\n\n")
	fmt.Fprintf(b, "`XMLExporter` is not part of parsing a spec, so the primary scope leaves it out. It stays as the `cli+xml` rows because the L3 differential oracle uses it as a test harness tool, which is why its runs are in the coverage reports. Adding it grows the handwritten denominator from %d to %d under RTA, and the specs' hits inside the scope from %d to %d.\n\n", t.Hand.Branches, x.Hand.Branches, t.Hand.Specs, x.Hand.Specs)
	fmt.Fprintf(b, "Hits in the primary scope that came only from exporter runs: **at most %d of %d handwritten and %d of %d generated (RTA)**. That is an upper bound, method by method: specs minus the plain command-line passes (`cli-ABC`, `cli-D`, `cli-E`). The exec files merge the probe JVM and the exporter JVM per case and the per-JVM files are gone, so an exporter hit cannot be told from a `SanyProbe` hit. `cli-D` and `cli-E` were last run before tiers D and E stopped growing, so cases kept since then count here too. Hits the other way, CLI ahead of specs: %d.\n\n", onlyHand, t.Hand.Specs, onlyGen, t.Gen.Specs, cliAhead)
}

// lineLevelGuards recomputes the guard numbers the way the prototype did: per
// line, a line counts as a guard when every branch on it has a guard shape, and
// all of the line's missed outcomes go.
func (a *Analysis) lineLevelGuards(scope string, cov *Coverage, probes ProbeTable) (int, int, int, int) {
	res := a.Results[scope][RTA]
	guard := map[string]bool{GuardNull: true, GuardInstanceof: true, GuardAssert: true, GuardThrow: true}
	type lineInfo struct {
		src  string
		line int
		all  bool
	}
	seen := map[string]*lineInfo{}
	total, hit, partly, never := 0, 0, 0, 0
	for _, k := range sortedKeys(cov) {
		m := a.methodOf(k)
		if !res.Reached[m.ID] || IsGenerated(m.Class) {
			continue
		}
		shapes, err := a.World.ClassifyMethod(m)
		if err != nil {
			fatalf("%v", err)
		}
		mc := cov.Methods[k]
		total += mc.BranchTotal
		hit += mc.BranchHit
		for _, r := range probes[k] {
			key := mc.Source + ":" + strconv.Itoa(r.Line)
			li := seen[key]
			if li == nil {
				li = &lineInfo{mc.Source, r.Line, true}
				seen[key] = li
			}
			if !guard[shapes[r.Ordinal].Class] {
				li.all = false
			}
		}
	}
	for _, li := range seen {
		if !li.all {
			continue
		}
		lc := cov.Lines[li.src][li.line]
		if lc.BranchHit > 0 {
			partly += lc.BranchMissed
		} else {
			never += lc.BranchMissed
		}
	}
	return total, hit, partly, never
}

func triple(hit, total int, g GuardMissed) string {
	return fmt.Sprintf("%.1f / %.1f / %.1f", Pct(hit, total), Pct(hit, total-g.Partly), Pct(hit, total-g.Partly-g.Never))
}

func (a *Analysis) writeGuards(b *strings.Builder, tallies map[string]map[Precision]ScopeTally) {
	primary := a.Cfg.Scopes[0].Name
	b.WriteString("## Guard shapes: a separate number\n\n")
	b.WriteString("**These are shapes, not proof.** `if (x != null)` can be real logic. Guard numbers are a reading aid: they never alter the map, the scopes or the soundness check, and the tables above do not use them.\n\n")
	b.WriteString("Attribution is per branch instruction and per arm. JaCoCo's public API stops at per-line counts, so `cmd/staticreach/java/.../BranchProbes.java` runs JaCoCo's own probe flow and filters and reads each instruction's covered arms before they are summed. The tool fails if those arms do not add up to JaCoCo's own per-method counters in the XML. So `candidate-map-branches.tsv` names the missed arm of every branch.\n\n")
	b.WriteString("Each cell is raw / excluding excludable missed arms on instructions that ran (another arm hit) / excluding also those on instructions that never ran. Under RTA, not refined. Only the denominator moves.\n\n")
	b.WriteString("| scope | handwritten, specs | handwritten, union | generated, specs | generated, union |\n|---|---|---|---|---|\n")
	for i := range a.Cfg.Scopes {
		n := a.Cfg.Scopes[i].Name
		t := tallies[n][RTA]
		fmt.Fprintf(b, "| `%s` | %s | %s | %s | %s |\n", n, triple(t.Hand.Specs, t.Hand.Branches, t.Hand.ExclSpecs), triple(t.Hand.Union, t.Hand.Branches, t.Hand.ExclUnion), triple(t.Gen.Specs, t.Gen.Branches, t.Gen.ExclSpecs), triple(t.Gen.Union, t.Gen.Branches, t.Gen.ExclUnion))
	}
	b.WriteString("\nShapes, first match wins:\n\n")
	b.WriteString("- `assert`: `$assertionsDisabled` just before the branch, or an arm that builds an `AssertionError` and throws. Only the throwing arm is excludable. JaCoCo already filters the `$assertionsDisabled` guard itself.\n")
	b.WriteString("- `scaffold-class` (generated): any branch in `ParseException`, `SimpleCharStream`, `Token`, `TokenMgrError`, `TLAplusParserConstants`.\n")
	b.WriteString("- `jj_ntk-cache` (generated): `(jj_ntk==-1)?jj_ntk():jj_ntk`, seen as `getfield jj_ntk; iconst_m1; if_icmp`.\n")
	b.WriteString("- `jj_scan_token` (generated): the branch straight after a call to it.\n")
	b.WriteString("- `jj_3R` (generated): any other branch in a `jj_3R_*` or `jj_3_*` lookahead method.\n")
	b.WriteString("- `null`: `ifnull`, `ifnonnull`. Either missed arm is excludable.\n")
	b.WriteString("- `instanceof`: `instanceof` then a branch. Either missed arm is excludable.\n")
	b.WriteString("- `throw-arm`: one arm reaches `athrow` with no other branch on the way. Only that arm is excludable: if the other arm is the missed one, the code has only ever thrown, which is not a guard left alone.\n")
	b.WriteString("- `switch-default`: a switch's default arm. Only the default is excludable.\n\n")

	res := a.Results[primary][RTA]
	type row struct {
		hand, gen map[string]*GuardTally
	}
	build := func(probes ProbeTable, cov *Coverage) row {
		r := row{map[string]*GuardTally{}, map[string]*GuardTally{}}
		for _, k := range sortedKeys(cov) {
			m := a.methodOf(k)
			if !res.Reached[m.ID] {
				continue
			}
			target := r.hand
			if IsGenerated(m.Class) {
				target = r.gen
			}
			if _, err := a.World.MethodGuards(m, probes[k], target); err != nil {
				fatalf("%v", err)
			}
		}
		return r
	}
	spec := build(a.ProbesSpecs, a.Specs)
	uni := build(a.ProbesUnion, a.Union)
	fmt.Fprintf(b, "Breakdown for `%s`, RTA. Outcomes is all outcomes on instructions of that shape; the other columns are excludable missed arms.\n\n", primary)
	b.WriteString("| code | shape | outcomes | specs, ran | specs, never ran | union, ran | union, never ran |\n|---|---|---:|---:|---:|---:|---:|\n")
	for _, side := range []struct {
		name string
		s, u map[string]*GuardTally
	}{{"handwritten", spec.hand, uni.hand}, {"generated", spec.gen, uni.gen}} {
		for _, c := range append(append([]string(nil), guardOrder...), GuardLogic) {
			st := side.s[c]
			ut := side.u[c]
			if st == nil && ut == nil {
				continue
			}
			if st == nil {
				st = &GuardTally{}
			}
			if ut == nil {
				ut = &GuardTally{}
			}
			fmt.Fprintf(b, "| %s | %s | %d | %d | %d | %d | %d |\n", side.name, c, st.Total, st.Excl.Partly, st.Excl.Never, ut.Excl.Partly, ut.Excl.Never)
		}
	}
	t := tallies[primary][RTA]
	b.WriteString("\n### The prototype's figures\n\n")
	b.WriteString("Prototype (`guards.py`): handwritten specs 73.2% / 75.6% / 81.3%, union 80.6% / 83.8% / 87.0%; generated specs 80.9% to 81.3%. These are corrected, for two reasons:\n\n")
	fmt.Fprintf(b, "1. **Scope.** The prototype's denominator was 4,227: it matched methods and JDK callbacks by bare name, and cut the Explorer by package prefix only. Exact signatures and the real edge cut give %d in the primary scope, so the raw figures move before any guard is counted.\n", t.Hand.Branches)
	b.WriteString("2. **Attribution.** The prototype classified lines and removed every missed outcome on a guard line. Arms are finer: the non-throwing arm of a throw guard stays. Its per-line rule, rerun on this tool's scopes, gives:\n\n")
	b.WriteString("| scope | specs | union |\n|---|---|---|\n")
	for _, n := range []string{primary, "cli"} {
		ts, hs, ps, ns := a.lineLevelGuards(n, a.Specs, a.ProbesSpecs)
		tu, hu, pu, nu := a.lineLevelGuards(n, a.Union, a.ProbesUnion)
		fmt.Fprintf(b, "| `%s`, handwritten, per line | %.1f / %.1f / %.1f | %.1f / %.1f / %.1f |\n", n, Pct(hs, ts), Pct(hs, ts-ps), Pct(hs, ts-ps-ns), Pct(hu, tu), Pct(hu, tu-pu), Pct(hu, tu-pu-nu))
	}
	b.WriteString("\nOn `cli` the per-line rule lands within a point of the prototype for specs; its union figures remain 2 to 4 points higher. That gap was not reconciled method by method.\n")
	fmt.Fprintf(b, "\nGenerated code moves once the generated shapes count: specs %s to %s, not 81.3%%. The prototype had no `jj_*` classes, so it saw only null checks there.\n\n", p1(t.Gen.Specs, t.Gen.Branches), p1(t.Gen.Specs, t.Gen.Branches-t.Gen.ExclSpecs.Partly-t.Gen.ExclSpecs.Never))
}

func isDepthToString(m *Method) bool {
	return m.Name == "toString" && strings.HasPrefix(m.Desc, "(ILtla2sany/semantic/Errors;)")
}

func isExploreMethod(m *Method) bool {
	return m.Name == "walkGraph" || m.Name == "walkGaph" || isDepthToString(m) || m.Name == "levelDataToString"
}

// writeDepthToString measures why depth-taking toString survives: its route runs
// through a toString() only the JDK callback assumption reaches. A measurement,
// not a cut: dropping the assumption could remove code that really runs.
func (a *Analysis) writeDepthToString(b *strings.Builder) {
	const holder = "tla2sany/semantic/SymbolTable"
	const sig = "toString:()Ljava/lang/String;"
	c := a.World.Classes[holder]
	if c == nil || c.MethodBy[sig] == nil {
		return
	}
	target := c.MethodBy[sig]
	callers := 0
	for _, m := range a.World.Methods {
		events, err := m.Events()
		if err != nil {
			fatalf("%v", err)
		}
		for i := range events {
			ev := &events[i]
			if ev.Kind == EvInvoke && ev.Ref.Class == holder && ev.Ref.Name == "toString" && ev.Ref.Desc == target.Desc {
				callers++
			}
			for _, t := range ev.Args {
				if t == holder {
					callers++
				}
			}
		}
	}
	lo, hi := 0, 0
	for _, e := range target.Lines {
		if lo == 0 || e.Line < lo {
			lo = e.Line
		}
		if e.Line > hi {
			hi = e.Line
		}
	}
	o, err := scopeOptions(a.Cfg, a.World, a.Cfg.Scope("cli-noexp"), RTA)
	if err != nil {
		fatalf("%v", err)
	}
	o.Cuts = append(o.Cuts, Cut{Name: "whatif", Class: holder, From: lo, To: hi})
	res := Analyse(a.World, o)
	base := a.Results["cli-noexp"][RTA]
	left, leftBr, gone, goneBr := 0, 0, 0, 0
	for _, m := range a.World.Methods {
		if !isDepthToString(m) || !base.Reached[m.ID] {
			continue
		}
		br := 0
		if u := a.Union.Methods[m.Key()]; u != nil {
			br = u.BranchTotal
		}
		if res.Reached[m.ID] {
			left++
			leftBr += br
		} else {
			gone++
			goneBr += br
		}
	}
	_, all := a.tallyResult(res)
	_, was := a.tallyResult(base)
	ran := "No report ran it."
	for _, cov := range []*Coverage{a.Specs, a.Union} {
		if mc := cov.Methods[target.Key()]; mc != nil && mc.InsnHit > 0 {
			ran = "`" + filepath.Base(cov.Path) + "` ran it."
		}
	}
	fmt.Fprintf(b, "`walkGraph`, `walkGaph` and `levelDataToString` drop out. **Depth-taking `toString` does not.** Its route is `SymbolTable.toString()`, which calls `Context.getContextEntryStringVector`: `SymbolTable` is instantiated, so its `toString()` counts as a JDK callback. The bytecode has %d direct calls or string concatenations on that type, so only the assumption that the JDK may stringify any instantiated object keeps it. %s In `cli-noexp` this is a what-if, not applied: without that one method's calls %d of these `toString` methods (%d branches) drop and %d (%d branches) stay, and the whole scope falls by %d branches. Proving that no `SymbolTable` reaches a JDK call needs points-to analysis, so the raw scopes keep them. Chris chose to cut it by hand in `port`, as the named entry `symboltable-tostring` in `scopes.conf`, which he can veto.\n\n", callers, ran, gone, goneBr, left, leftBr, was-all)
}

func (a *Analysis) writeExplorer(b *strings.Builder) {
	all, creators := a.ExplorerReferences()
	b.WriteString("## Cutting the Explorer\n\n")
	fmt.Fprintf(b, "Bytecode scan of all %d classes: `new tla2sany/explorer/Explorer` occurs %d time(s):", len(a.World.AppNames), len(creators))
	for _, c := range creators {
		fmt.Fprintf(b, " `%s`, line %d", shortKey(c.Caller), c.Line)
	}
	b.WriteString(". That is `drivers/SANY.java`'s `if (doDebugging)` block, entered only with `-d` or `-D`. The cut removes the body, lines 760-764, and the guard's true outcome; the guard itself on line 758 stays, because it runs. No Java source changes.\n\n")
	if len(creators) != 1 {
		b.WriteString("**WARNING: not exactly one creator.**\n\n")
	}
	counts := map[string]int{}
	for _, s := range all {
		short := s.Target[strings.LastIndexByte(s.Target, '/')+1:]
		if s.Kind == "new" {
			short = "new " + strings.TrimSuffix(short, ".")
		}
		counts[short]++
	}
	names := make([]string, 0, len(counts))
	for n := range counts {
		names = append(names, n)
	}
	sort.Strings(names)
	b.WriteString("Other references into `tla2sany/explorer/` from outside the package, by target: ")
	parts := make([]string, 0, len(names))
	for _, n := range names {
		parts = append(parts, fmt.Sprintf("`%s` %d", n, counts[n]))
	}
	b.WriteString(strings.Join(parts, ", ") + ". Apart from the `Explorer` sites at lines 760 and 762, all of them use `ExplorerVisitor`, a plain visitor type, which can stay reachable without any `Explorer`.\n\n")

	b.WriteString("Do the methods semantic nodes implement for `ExploreNode` drop out once `Explorer` is gone? Methods (branches):\n\n")
	b.WriteString("| methods | with Explorer, RTA | Explorer cut, RTA | with Explorer, CHA | Explorer cut, CHA |\n|---|---:|---:|---:|---:|\n")
	count := func(scope string, p Precision, match func(*Method) bool) (int, int) {
		r := a.Results[scope][p]
		n, br := 0, 0
		for _, m := range a.World.Methods {
			if r.Reached[m.ID] && match(m) {
				n++
				if u := a.Union.Methods[m.Key()]; u != nil {
					br += u.BranchTotal
				}
			}
		}
		return n, br
	}
	rows := []struct {
		name  string
		match func(*Method) bool
	}{
		{"`walkGraph`, `walkGaph`", func(m *Method) bool { return m.Name == "walkGraph" || m.Name == "walkGaph" }},
		{"`toString(int, Errors)`", isDepthToString},
		{"`levelDataToString`", func(m *Method) bool { return m.Name == "levelDataToString" }},
	}
	for _, row := range rows {
		n1, b1 := count("cli", RTA, row.match)
		n2, b2 := count("cli-noexp", RTA, row.match)
		n3, b3 := count("cli", CHA, row.match)
		n4, b4 := count("cli-noexp", CHA, row.match)
		fmt.Fprintf(b, "| %s | %d (%d) | %d (%d) | %d (%d) | %d (%d) |\n", row.name, n1, b1, n2, b2, n3, b3, n4, b4)
	}
	b.WriteString("\n")
	_, b1 := count("cli", RTA, isExploreMethod)
	_, b2 := count("cli-noexp", RTA, isExploreMethod)

	a.writeDepthToString(b)

	r := a.Results["cli-noexp"][RTA]

	rWith := a.Results["cli"][RTA]
	type cc struct {
		class string
		n, br int
	}
	by := map[string]*cc{}
	total := 0
	for _, k := range sortedKeys(a.Union) {
		m := a.methodOf(k)
		if rWith.Reached[m.ID] && !r.Reached[m.ID] {
			c := by[m.Class.Name]
			if c == nil {
				c = &cc{class: m.Class.Name}
				by[m.Class.Name] = c
			}
			c.n++
			c.br += a.Union.Methods[k].BranchTotal
			total += a.Union.Methods[k].BranchTotal
		}
	}
	list := make([]*cc, 0, len(by))
	for _, c := range by {
		list = append(list, c)
	}
	sort.Slice(list, func(i, j int) bool {
		if list[i].br != list[j].br {
			return list[i].br > list[j].br
		}
		return list[i].class < list[j].class
	})
	fmt.Fprintf(b, "The cut removes %d branches under RTA in all, %d of them in those explore methods. Largest classes (methods, branches):", total, b1-b2)
	for i, c := range list {
		if i == 12 {
			break
		}
		fmt.Fprintf(b, " `%s` %d, %d;", c.class[strings.LastIndexByte(c.class, '/')+1:], c.n, c.br)
	}
	b.WriteString("\n\n")
}

func shortKey(k string) string {
	if i := strings.Index(k, ":"); i >= 0 {
		k = k[:i]
	}
	return k
}

func (a *Analysis) writeFlags(b *strings.Builder) {
	b.WriteString("## Debug and dump flags: the cuts to veto\n\n")
	b.WriteString("`cli-nodebug` cuts every switch that only enables debugging or dumping. Each row is an edge removed; delete a row in `scopes.conf` to veto it. Branches is what that flag alone removes on top of the Explorer cut, under RTA.\n\n")
	b.WriteString("| flag | cut | why | branches |\n|---|---|---|---:|\n")
	s := a.Cfg.Scope("cli-noexp")
	base := a.TallyScope("cli-noexp", RTA)
	seen := map[string]bool{}
	for _, c := range a.Cfg.Cuts {
		o, err := scopeOptions(a.Cfg, a.World, s, RTA)
		if err != nil {
			fatalf("%v", err)
		}
		delta := "(the Explorer cut itself)"
		if c.Name != "explorer" {
			if !seen[c.Name] {
				seen[c.Name] = true
				o.Cuts = append(o.Cuts, a.Cfg.CutsNamed(c.Name)...)
				res := Analyse(a.World, o)
				_, br := a.tallyResult(res)
				delta = strconv.Itoa(base.Hand.Branches + base.Gen.Branches - br)
			} else {
				delta = "(same flag, above)"
			}
		}
		fmt.Fprintf(b, "| `%s` | `%s` lines %d-%d | %s | %s |\n", c.Flag, c.Class, c.From, c.To, c.Why, delta)
	}
	b.WriteString("\nConsidered and not cut, because they change behaviour rather than only print: `-S`, `-L`, `-G`, `-error-codes`, `-suppressMessages`, `-messagesAsErrors`, `-help`, `-Dtla2sany.drivers.SANY.errors2stderr`; for the exporter `-o`, `-t`, `-r`, `-u`, `-I`, `-help`. `-stat` is accepted and ignored at `drivers/SANY.java`, so it has no edge to cut.\n\n")
}

// tallyResult sums a one-off result the way TallyScope does, without guards.
func (a *Analysis) tallyResult(res *Result) (int, int) {
	methods, br := 0, 0
	for _, k := range sortedKeys(a.Union) {
		m := a.methodOf(k)
		if !res.Reached[m.ID] {
			continue
		}
		u := a.Union.Methods[k]
		total, _, _ := adjust(a.Union, u.Source, u.BranchTotal, u.BranchHit, removedLines(a.World, m, res.Opts))
		methods++
		br += total
	}
	return methods, br
}

func (a *Analysis) writeRefinement(b *strings.Builder, tallies map[string]map[Precision]ScopeTally) {
	primary := a.Cfg.Scopes[0].Name
	cf := a.World.Constants()
	b.WriteString("## Intra-method refinement\n\n")
	b.WriteString("One rule, in `refine.go`: a branch on a field or a static method whose value the closed world fixes is folded, and code only that branch reaches is dead. A field qualifies when it is never written, written only with its default, or a static written once with a constant in its own initialiser before that initialiser can run other code. javac has already folded `static final` compile-time constants: there is no branch left to find.\n\n")
	res := a.Results[primary][RTAConst]
	var ms []*Method
	for _, m := range a.World.Methods {
		if res.Reached[m.ID] {
			ms = append(ms, m)
		}
	}
	facts := a.World.ConstantFacts(ms)
	fmt.Fprintf(b, "%d fields qualify in the whole build, and very few guard a branch. In `%s`, by what licensed the removal (outcomes, before capping at what JaCoCo counts on the line):\n\n", len(cf.values), primary)
	b.WriteString("| fact | why | outcomes |\n|---|---|---:|\n")
	for _, f := range facts {
		name := f.Fact
		why := cf.why[f.Fact]
		if name == "" {
			name = "(branches in code the folds made dead)"
		}
		if why == "" && f.Fact != "" {
			why = "method returns a constant"
		}
		fmt.Fprintf(b, "| `%s` | %s | %d |\n", name, why, f.Outcomes)
	}
	b.WriteString("\n| scope | branches, RTA | RTA+const | removed |\n|---|---:|---:|---:|\n")
	for i := range a.Cfg.Scopes {
		n := a.Cfg.Scopes[i].Name
		r := tallies[n][RTA]
		c := tallies[n][RTAConst]
		fmt.Fprintf(b, "| `%s` | %d | %d | %d |\n", n, r.Hand.Branches+r.Gen.Branches, c.Hand.Branches+c.Gen.Branches, r.Hand.Branches+r.Gen.Branches-c.Hand.Branches-c.Gen.Branches)
	}
	b.WriteString("\nThe honest result: SANY barely has such branches, so sound reasoning at this level buys almost nothing. Removing more would need value tracking (what can this field hold at this point), which is where cheap and sound stop. The soundness check tests each fold against every run: a line whose instructions are all ruled out must not have run, and a line that lost an outcome cannot have hit more than the rest.\n\n")
	fmt.Fprintf(b, "A cut also removes the guard outcome that leads into it: %d handwritten outcome(s) in `%s`, `if (doDebugging)`'s true arm.\n\n", tallies[primary][RTA].RemovedHand, primary)
}

func (a *Analysis) writeSoundness(b *strings.Builder, rep *SoundnessReport) {
	b.WriteString("## Soundness\n\n")
	b.WriteString("`staticreach check`, and every full run, compares each coverage report with its scope and exits 1 on anything unexplained: every method JaCoCo saw run must be reachable, at CHA, RTA and RTA+const. A spec run's harness is an extra root, recorded in the table; upstream's test classes are all roots, since JUnit calls them by reflection. Lines ruled out by a cut or a fold must not have run either, nor unreachable lexer arms, and every reflective site in reached code must be resolved.\n\n")
	fmt.Fprintf(b, "**Result: %d unexplained, %d explained.**\n\n", rep.Unexplained, rep.Explained)
	b.WriteString("| report | kind | scope | extra roots | methods that ran | outside CHA | outside RTA | outside RTA+const | ruled-out lines hit |\n|---|---|---|---|---:|---:|---:|---:|---:|\n")
	for _, rr := range rep.Runs {
		fmt.Fprintf(b, "| `%s` | %s | `%s` | %s | %d | %d | %d | %d | %d |\n", rr.Path, rr.Run.Kind, rr.Run.Label(), rr.Run.Harness, rr.Ran, len(rr.Outside[CHA]), len(rr.Outside[RTA]), len(rr.Outside[RTAConst]), len(rr.DeadHits))
	}
	for _, e := range rep.Errors {
		fmt.Fprintf(b, "\nERROR: %s\n", e)
	}
	b.WriteString("\nThe plain command's coverage report stays entirely inside these scopes with no extra roots. The original measurement also checked reports taken under the measuring wrapper and upstream's test classes, with those classes as extra roots; neither is committed here, so those rows, and the table of what ran only through them, are not reproduced by this tool.\n\n")
}

func (a *Analysis) writeCallbacks(b *strings.Builder) {
	primary := a.Cfg.Scopes[0].Name
	r := a.Results[primary][RTA]
	byPkg := map[string]int{}
	for _, m := range a.World.Methods {
		if !r.Reached[m.ID] {
			continue
		}
		k := m.Class.Name
		if i := strings.IndexByte(k, '/'); i >= 0 {
			k = k[:i]
		}
		byPkg[k]++
	}
	names := make([]string, 0, len(byPkg))
	for n := range byPkg {
		names = append(names, n)
	}
	sort.Strings(names)
	b.WriteString("## How the graph is built\n\n")
	b.WriteString("A small class-file parser in `classfile.go` and `bytecode.go`; no text parsing of `javap`. Methods are name plus descriptor throughout.\n\n")
	b.WriteString("- **Dispatch** follows JVMS 5.4.6 selection, so a call lands on a superclass's method or an interface default when that is what a class inherits.\n")
	b.WriteString("- **JDK callbacks** come from the hierarchy, read from the JDK's jmods: a method is a callback when a JDK supertype of an instantiated class declares its exact signature and selection picks it. No list of names. `toString(int, Errors)` is not one.\n")
	b.WriteString("- **Class initialisation** is modelled: `new`, `getstatic`, `putstatic`, `invokestatic`, plus superclasses and interfaces with default methods, per JVMS 5.5.\n")
	b.WriteString("- **`invokedynamic`** reads the real bootstrap arguments. A lambda or method reference reaches its implementation method and instantiates a modelled class for its interface; string concatenation is a `toString()` on each argument type. Any other bootstrap stops the tool with an error.\n\n")
	fmt.Fprintf(b, "`%s` under RTA reaches, by top package, in methods:", primary)
	for i, n := range names {
		if i > 0 {
			b.WriteString(",")
		}
		fmt.Fprintf(b, " `%s` %d", n, byPkg[n])
	}
	b.WriteString(". Only `tla2sany` has JaCoCo counters; the rest are in `candidate-map.tsv` with blank branch columns.\n\n")
}

func (a *Analysis) writeReflection(b *strings.Builder) {
	sites := a.ReflectionSites()
	b.WriteString("## Reflection\n\n")
	b.WriteString("Scan of reached code at the CHA upper bound, in every scope, for `Class.forName`, `newInstance`, `getMethod` and friends, `Method.invoke`, `ServiceLoader`, `Proxy`, `MethodHandles.Lookup`, object streams and `System.loadLibrary`. An unresolved site fails the check.\n\n")
	b.WriteString("| caller | line | API | resolved to | scopes |\n|---|---:|---|---|---:|\n")
	for _, s := range sites {
		to := "**unresolved**"
		if r := a.Cfg.reflectFor(s.Caller); r != nil {
			to = "`" + r.Target + "`: " + r.Why
		}
		fmt.Fprintf(b, "| `%s` | %d | `%s` | %s | %d of %d |\n", shortKey(s.Caller), s.Line, s.API, to, len(s.Scopes), len(a.Cfg.Scopes))
	}
	b.WriteString("\nTLC's `TLAClass` and `SpecProcessor` reflection are not reached from any scope: they sit above SANY. `tla2sany` itself has none.\n\n")
}

func (a *Analysis) writeOutside(b *strings.Builder) {
	b.WriteString("## What is outside each scope\n\nBranches outside the scope, handwritten+generated, by package, under RTA, with what the union hit anyway in parentheses. A hit outside means a harness or an upstream test called the method directly.\n\n")
	pkgs := map[string]bool{}
	data := map[string]map[string]PackageOutside{}
	for i := range a.Cfg.Scopes {
		n := a.Cfg.Scopes[i].Name
		data[n] = map[string]PackageOutside{}
		for _, p := range a.OutsideByPackage(n, RTA) {
			pkgs[p.Package] = true
			data[n][p.Package] = p
		}
	}
	names := make([]string, 0, len(pkgs))
	for p := range pkgs {
		names = append(names, p)
	}
	sort.Strings(names)
	b.WriteString("| package |")
	for i := range a.Cfg.Scopes {
		fmt.Fprintf(b, " `%s` |", a.Cfg.Scopes[i].Name)
	}
	b.WriteString("\n|---|")
	for range a.Cfg.Scopes {
		b.WriteString("---:|")
	}
	b.WriteString("\n")
	for _, p := range names {
		fmt.Fprintf(b, "| `%s` |", strings.TrimPrefix(p, "tla2sany/"))
		for i := range a.Cfg.Scopes {
			x := data[a.Cfg.Scopes[i].Name][p]
			fmt.Fprintf(b, " %d+%d (%d) |", x.Hand, x.Gen, x.Hit)
		}
		b.WriteString("\n")
	}
	primary := a.Cfg.Scopes[0].Name
	fmt.Fprintf(b, "\nLargest classes outside `%s` (RTA), branches (hit by the union):", primary)
	for i, c := range a.OutsideByClass(primary, RTA) {
		if i == 25 {
			break
		}
		fmt.Fprintf(b, " `%s` %d (%d);", c.Class[strings.LastIndexByte(c.Class, '/')+1:], c.Branch, c.Hit)
	}
	b.WriteString("\n\nThe per-method answer is in `candidate-map.tsv`: each scope's column holds `R` (RTA reaches it), `C` (only CHA) or `-`.\n\n")
}

const lexerSource = "tla2sany/parser/TLAplusParserTokenManager.java"

func armCoverage(cov *Coverage, arm LexerArm) (int, int, int) {
	lines := cov.Lines[lexerSource]
	tot, hit, insn := 0, 0, 0
	for l := arm.Line; l <= arm.EndLine; l++ {
		c := lines[l]
		tot += c.BranchHit + c.BranchMissed
		hit += c.BranchHit
		insn += c.InsnHit
	}
	return tot, hit, insn
}

// checkLexer holds the lexer's unreachable arms to both reports.
func (a *Analysis) checkLexer(rep *SoundnessReport) {
	lx, err := AnalyseLexer(filepath.Join(a.Paths.Sources(), lexerSource))
	if err != nil {
		rep.Errors = append(rep.Errors, err.Error())
		return
	}
	for _, st := range lx.States {
		for _, d := range st.DeadArms {
			for _, cov := range []*Coverage{a.Specs, a.Union} {
				if _, _, insn := armCoverage(cov, d); insn > 0 {
					rep.Errors = append(rep.Errors, fmt.Sprintf("lexer arm in %s, state %v at line %d, is unreachable but ran in %s", st.Name, d.States, d.Line, cov.Path))
				}
			}
		}
	}
}

func (a *Analysis) writeLexer(b *strings.Builder) error {
	lx, err := AnalyseLexer(filepath.Join(a.Paths.Sources(), lexerSource))
	if err != nil {
		return err
	}
	total, hitU, hitS := 0, 0, 0
	for _, k := range sortedKeys(a.Union) {
		if a.methodOf(k).Class.Name == "tla2sany/parser/TLAplusParserTokenManager" {
			total += a.Union.Methods[k].BranchTotal
			hitU += a.Union.Methods[k].BranchHit
			hitS += a.Specs.Methods[k].BranchHit
		}
	}
	primary := a.Cfg.Scopes[0].Name
	gen := a.TallyScope(primary, RTA).Gen
	b.WriteString("## The lexer\n\n")
	fmt.Fprintf(b, "The token manager has %d branches; specs hit %d, the union %d. Reachability of its automaton states is decidable and the state graph is easy to read from the generated source, so `lexer.go` does: every `jjMoveNfa_k` arm, every transition (`jjCheckNAdd*`, `jjAddStates`, direct `jjstateSet` writes, through `jjnextStates`), and every start (`jjMoveNfa_k(n, ...)`, `jjStartNfaWithStates_k`, the returns of `jjStopStringLiteralDfa_k`). This is about automaton arms only: grammar productions and choice points stay with the climb.\n\n", total, hitS, hitU)
	b.WriteString("All six lexical states are reachable from `jjnewLexState` and `SwitchTo`.\n\n")
	b.WriteString("| lexical state | NFA states | reached | arms | unreachable arms | branches on them | literal DFA bit tests | unsatisfiable |\n|---|---:|---:|---:|---:|---:|---:|---:|\n")
	sumDead, sumBr, sumHit := 0, 0, 0
	var detail strings.Builder
	sections := []string{"curChar<64", "64..127", ">=128"}
	for _, st := range lx.States {
		br := 0
		for _, d := range st.DeadArms {
			t, h, _ := armCoverage(a.Union, d)
			br += t
			sumBr += t
			sumHit += h
		}
		sumDead += len(st.DeadArms)
		fmt.Fprintf(b, "| %s | %d | %d | %d | %d | %d | %d | %d |\n", st.Name, st.NfaStates, st.Reached, st.Arms, len(st.DeadArms), br, st.DfaTests, len(st.DfaDead))
		if len(st.DeadArms) > 0 {
			fmt.Fprintf(&detail, "- %s:", st.Name)
			for s, label := range sections {
				var parts []string
				for _, d := range st.DeadArms {
					if d.Section == s {
						parts = append(parts, fmt.Sprintf("%s (line %d)", intsText(d.States), d.Line))
					}
				}
				if len(parts) > 0 {
					fmt.Fprintf(&detail, " `%s`: %s;", label, strings.Join(parts, ", "))
				}
			}
			detail.WriteString("\n")
		}
	}
	fmt.Fprintf(b, "\n**%d arms cannot run, on any input:** no transition or start names their state. They hold %d branches, of which the union hit %d. That is the check, and the tool fails if one ever runs. Most likely JavaCC merges these states' moves into composite start states and leaves their original arms with nothing to enter them; the generated code shows such merged arms, but this explanation was not checked against JavaCC's generator. The arms stay in the generated denominators above; subtracting them turns generated `%s` under RTA from %s to %s of %d, specs, and is still a statement about the lexer alone.\n\n", sumDead, sumBr, sumHit, primary, p1(gen.Specs, gen.Branches), p1(gen.Specs, gen.Branches-sumBr), gen.Branches-sumBr)
	b.WriteString("By lexical state, input section and NFA state:\n\n")
	b.WriteString(detail.String())
	b.WriteString("\nThe literal DFA methods get a forward union of the active bit masks, in depth order: no bit test is unsatisfiable by the masks that reach it. That is a statement about the masks only.\n\n")
	b.WriteString("What deciding the rest would take: the remaining missed branches sit inside reachable arms, on tests such as `kind > N` and the character masks, whose outcome depends on which NFA states are active together. Deciding those exactly means running the subset construction over the generated moves, active state sets against character classes, recording which outcome each test can take. It is finite and on the order of a few hundred lines that interpret the generated arms; it is not done here. `jjCanMove` above 127 would need the character tables too.\n\n")
	return nil
}

func intsText(v []int) string {
	parts := make([]string, len(v))
	for i, n := range v {
		parts[i] = strconv.Itoa(n)
	}
	return strings.Join(parts, "+")
}

func (a *Analysis) writeLimits(b *strings.Builder) {
	b.WriteString("## Limits\n\n")
	b.WriteString("- **Upper bound.** Reachable means not ruled out. A method in the map can be dead for reasons the graph cannot see, such as a check no value can fail. Unhit branches in the map are candidates, not known-reachable targets.\n")
	b.WriteString("- **Whole methods.** Except for the fold and the cuts, every branch in a reachable method counts.\n")
	b.WriteString("- **JDK as a black box.** JDK bodies are not analysed. Any override of a JDK-declared method on an instantiated class is assumed called; nothing checks that the object reaches that JDK call. This keeps it sound and keeps it loose.\n")
	a.writeLooseness(b)
	b.WriteString("- **The test rows are weak.** Rooting all test methods reaches a lot; they mainly show the analysis does not drop reflection through the product. Spec runs are the bar.\n")
	b.WriteString("- **Not run.** TLC was not started for this report. The `parse` scope follows `SpecProcessor.processSpec` from reading its source, so its seeds are by hand, and TLC's later use of the semantic graph is not an entry.\n")
	b.WriteString("\n")
	b.WriteString("## Inputs\n\n| input | sha256 prefix |\n|---|---|\n")
	fmt.Fprintf(b, "| class files, %d application classes | `%s` |\n", len(a.World.AppNames), a.classFingerprint())
	for _, rel := range []string{"data/staticreach/specs.xml.gz", "data/staticreach/specs-cli.xml.gz", "data/staticreach/union.xml.gz", "data/staticreach/generated-missed.tsv"} {
		fmt.Fprintf(b, "| `%s` | `%s` |\n", rel, fileHash(a.Paths.abs(rel)))
	}
	sum := sha256.Sum256([]byte(a.Cfg.Text))
	fmt.Fprintf(b, "| scopes file | `%x` |\n", sum[:8])
	b.WriteString("\n")
}

// writeLooseness measures how much of the primary scope hangs on the JDK callback
// assumption alone. It is a measurement: the map keeps these methods.
func (a *Analysis) writeLooseness(b *strings.Builder) {
	primary := a.Cfg.Scope(a.Cfg.Scopes[0].Name)
	o, err := scopeOptions(a.Cfg, a.World, primary, RTA)
	if err != nil {
		fatalf("%v", err)
	}
	o.NoCallbacks = true
	res := Analyse(a.World, o)
	base := a.Results[primary.Name][RTA]
	inScope, hit := 0, 0
	for _, k := range sortedKeys(a.Union) {
		m := a.methodOf(k)
		if base.Reached[m.ID] && !res.Reached[m.ID] && !IsGenerated(m.Class) {
			inScope += a.Union.Methods[k].BranchTotal
			hit += a.Specs.Methods[k].BranchHit
		}
	}
	fmt.Fprintf(b, "- **Measuring that looseness.** %d handwritten branches of `%s` (RTA) are reachable only through an assumed JDK callback, and specs hit %d of them, so the assumption is doing real work and also inflating the denominator by an unknown part of the rest.\n", inScope, primary.Name, hit)
	if len(a.World.Missing) > 0 {
		names := make([]string, 0, len(a.World.Missing))
		for n := range a.World.Missing {
			names = append(names, "`"+n+"`")
		}
		sort.Strings(names)
		fmt.Fprintf(b, "- **Classes no input provides.** %d names looked up during the analysis are in neither the build nor the JDK, so nothing is inherited through them: %s.\n", len(names), strings.Join(names, ", "))
	}
}

func fileHash(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		fatalf("%v", err)
	}
	sum := sha256.Sum256(data)
	return fmt.Sprintf("%x", sum[:8])
}

func (a *Analysis) classFingerprint() string {
	h := sha256.New()
	root := a.Paths.Classes()
	var paths []string
	err := filepath.WalkDir(root, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(p, ".class") {
			paths = append(paths, p)
		}
		return nil
	})
	if err != nil {
		fatalf("%v", err)
	}
	sort.Strings(paths)
	for _, p := range paths {
		rel, err := filepath.Rel(root, p)
		if err != nil {
			fatalf("%v", err)
		}
		data, err := os.ReadFile(p)
		if err != nil {
			fatalf("%v", err)
		}
		h.Write([]byte(rel))
		h.Write(data)
	}
	return fmt.Sprintf("%x", h.Sum(nil)[:8])
}
