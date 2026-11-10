// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	_ "embed"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Label names a run's scope with any cuts held against it on top.
func (r RunDef) Label() string {
	if len(r.ExtraCuts) == 0 {
		return r.Scope
	}
	return r.Scope + "+" + strings.Join(r.ExtraCuts, "+")
}

const parserClass = "tla2sany/parser/TLAplusParser"

// NtkSite is one (jj_ntk==-1)?jj_ntk():jj_ntk, found in the class file.
type NtkSite struct {
	M    *Method
	Line int
	Off  int
	// InPort says that the port scope reaches the method, under RTA.
	InPort bool
	Total  int
	// Compute and Memo say which side an exec file took, specs then union.
	Compute, Memo [2]bool
}

// ntkSites scans every method of the parser for the ternary, so that a site in a
// method no report covers still counts.
func (a *Analysis) ntkSites() []NtkSite {
	c := a.World.Classes[parserClass]
	if c == nil {
		fatalf("no class %s", parserClass)
	}
	port := a.Results["port"][RTA]
	var out []NtkSite
	for _, m := range c.Methods {
		shapes, err := a.World.ClassifyMethod(m)
		if err != nil {
			fatalf("%v", err)
		}
		insns, err := m.Decode()
		if err != nil {
			fatalf("%v", err)
		}
		ords := make([]int, 0, len(shapes))
		for ord, s := range shapes {
			if s.Class == GuardNtk {
				ords = append(ords, ord)
			}
		}
		sort.Ints(ords)
		for _, ord := range ords {
			s := shapes[ord]
			// javac writes if_icmpne to the memo side: arm 0 falls into jj_ntk(), arm 1
			// jumps to the field read. Anything else would make the names wrong.
			if insns[ord].Op != opIfIcmpne {
				fatalf("%s line %d: ternary compiled to opcode %d", m.Key(), s.Line, insns[ord].Op)
			}
			site := NtkSite{M: m, Line: s.Line, Off: s.Off, InPort: port.Reached[m.ID]}
			for i, table := range []ProbeTable{a.ProbesSpecs, a.ProbesUnion} {
				for _, r := range table[m.Key()] {
					if r.Ordinal == ord {
						site.Total = r.Total
						site.Compute[i] = r.Has(0)
						site.Memo[i] = r.Has(1)
					}
				}
			}
			out = append(out, site)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Line < out[j].Line })
	return out
}

// sourceNtkCount counts the ternary in the generated source, as a check on the scan.
func (a *Analysis) sourceNtkCount() (int, error) {
	b, err := os.ReadFile(a.Paths.Sources() + "/" + parserClass + ".java")
	if err != nil {
		return 0, err
	}
	return strings.Count(string(b), "(jj_ntk==-1)?jj_ntk():jj_ntk"), nil
}

func (a *Analysis) ntkSitesTSV() string {
	var sb strings.Builder
	sb.WriteString(tsvHeader)
	sb.WriteString("method\tdescriptor\tline\toffset\tin_port\toutcomes\tcompute_specs\tmemo_specs\tcompute_union\tmemo_union\n")
	yn := func(b bool) string {
		if b {
			return "yes"
		}
		return "no"
	}
	for _, s := range a.ntkSites() {
		fmt.Fprintf(&sb, "%s\t%s\t%d\t%d\t%s\t%d\t%s\t%s\t%s\t%s\n", s.M.Name, s.M.Desc, s.Line, s.Off, yn(s.InPort), s.Total, yn(s.Compute[0]), yn(s.Memo[0]), yn(s.Compute[1]), yn(s.Memo[1]))
	}
	return sb.String()
}

// NtkCount is the ternary's outcomes, for a set of sites and one exec file.
type NtkCount struct {
	Sites, Outcomes, Hit, Both, ComputeOnly, MemoOnly, Neither int
}

func countNtk(sites []NtkSite, which int, onlyPort bool) NtkCount {
	var n NtkCount
	for _, s := range sites {
		if onlyPort && !s.InPort {
			continue
		}
		n.Sites++
		n.Outcomes += s.Total
		c, m := s.Compute[which], s.Memo[which]
		switch {
		case c && m:
			n.Both++
			n.Hit += 2
		case c:
			n.ComputeOnly++
			n.Hit++
		case m:
			n.MemoOnly++
			n.Hit++
		default:
			n.Neither++
		}
	}
	return n
}

//go:embed generated-review.conf
var generatedReviewText string

// ReviewRow is the review's verdict counts for one kind of generated code.
type ReviewRow struct {
	Kind                                          string
	Shown, Believed, Ternary, Unknown, Blind, Big int
}

// Total is what the review accounts for in this kind.
func (r ReviewRow) Total() int {
	return r.Shown + r.Believed + r.Ternary + r.Unknown + r.Blind + r.Big
}

// parseReview reads generated-review.conf.
func parseReview(text string) ([]ReviewRow, error) {
	var out []ReviewRow
	for n, line := range strings.Split(text, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || line[0] == '#' {
			continue
		}
		w := strings.Split(line, "|")
		if len(w) != 7 {
			return nil, fmt.Errorf("generated-review.conf line %d: want 7 fields", n+1)
		}
		r := ReviewRow{Kind: w[0]}
		for i, p := range []*int{&r.Shown, &r.Believed, &r.Ternary, &r.Unknown, &r.Blind, &r.Big} {
			v, err := strconv.Atoi(w[i+1])
			if err != nil {
				return nil, fmt.Errorf("generated-review.conf line %d: %w", n+1, err)
			}
			*p = v
		}
		out = append(out, r)
	}
	return out, nil
}

// lineKinds reads data/staticreach/generated-missed.tsv: file and line to kind.
func (a *Analysis) lineKinds() (map[string]string, error) {
	b, err := os.ReadFile(a.Paths.abs("data/staticreach/generated-missed.tsv"))
	if err != nil {
		return nil, err
	}
	out := map[string]string{}
	for i, l := range strings.Split(string(b), "\n") {
		if i == 0 || l == "" || l[0] == '#' || strings.HasPrefix(l, "file\t") {
			continue
		}
		w := strings.Split(l, "\t")
		if len(w) < 4 {
			return nil, fmt.Errorf("generated-missed.tsv line %d: too few fields", i+1)
		}
		out[w[0]+":"+w[1]] = w[3]
	}
	return out, nil
}

// KindMissed is the union's missed generated outcomes of one kind.
type KindMissed struct {
	Kind              string
	All, InScope      int
	NtkAll, NtkInPort int
}

// generatedMissed counts missed generated outcomes per kind, per instruction, in
// the union, inside and outside a scope.
func (a *Analysis) generatedMissed(scope string) (map[string]*KindMissed, error) {
	kinds, err := a.lineKinds()
	if err != nil {
		return nil, err
	}
	res := a.Results[scope][RTA]
	out := map[string]*KindMissed{}
	for _, k := range sortedKeys(a.Union) {
		m := a.methodOf(k)
		if !IsGenerated(m.Class) {
			continue
		}
		shapes, err := a.World.ClassifyMethod(m)
		if err != nil {
			return nil, err
		}
		pkg := m.Class.Name[:strings.LastIndexByte(m.Class.Name, '/')]
		for _, r := range a.ProbesUnion[k] {
			miss := r.Total - r.HitCount()
			if miss == 0 {
				continue
			}
			kind := kinds[pkg+"/"+m.Class.SourceFile+":"+strconv.Itoa(r.Line)]
			if kind == "" {
				kind = "(line not in generated-missed.tsv)"
			}
			km := out[kind]
			if km == nil {
				km = &KindMissed{Kind: kind}
				out[kind] = km
			}
			km.All += miss
			isNtk := shapes[r.Ordinal] != nil && shapes[r.Ordinal].Class == GuardNtk
			if isNtk {
				km.NtkAll += miss
			}
			if res.Reached[m.ID] {
				km.InScope += miss
				if isNtk {
					km.NtkInPort += miss
				}
			}
		}
	}
	return out, nil
}

// PortReview is the review's verdicts applied to the port scope.
type PortReview struct {
	Rows                                          []PortReviewRow
	Shown, Believed, Ternary, Unknown, Blind, Big int
	// Unreconciled names kinds whose measured misses differ from the review's counts.
	Unreconciled []string
}

// PortReviewRow is one kind inside the scope.
type PortReviewRow struct {
	ReviewRow
	Measured, Outside int
}

// applyReview takes the review's verdict counts into a scope. What the scope leaves
// out comes off the shown-dead count of its kind: those outcomes are in methods with
// no static path, which is the review's own reason for its shown-dead calls there.
func applyReview(review []ReviewRow, missed map[string]*KindMissed) PortReview {
	var out PortReview
	seen := map[string]bool{}
	for _, r := range review {
		seen[r.Kind] = true
		km := missed[r.Kind]
		if km == nil {
			km = &KindMissed{}
		}
		row := PortReviewRow{ReviewRow: r, Measured: km.InScope, Outside: km.All - km.InScope}
		if km.All != r.Total() {
			out.Unreconciled = append(out.Unreconciled, fmt.Sprintf("%s: review accounts for %d, %d missed now", r.Kind, r.Total(), km.All))
		}
		row.Shown -= row.Outside
		if row.Shown < 0 {
			out.Unreconciled = append(out.Unreconciled, fmt.Sprintf("%s: %d outside the scope, more than the %d shown dead", r.Kind, row.Outside, r.Shown))
			row.Shown = 0
		}
		out.Rows = append(out.Rows, row)
		out.Shown += row.Shown
		out.Believed += row.Believed
		out.Ternary += row.Ternary
		out.Unknown += row.Unknown
		out.Blind += row.Blind
		out.Big += row.Big
	}
	names := make([]string, 0, len(missed))
	for k := range missed {
		names = append(names, k)
	}
	sort.Strings(names)
	for _, k := range names {
		if !seen[k] {
			out.Unreconciled = append(out.Unreconciled, fmt.Sprintf("%s: %d missed now, no verdict from the review", k, missed[k].All))
		}
	}
	return out
}

// directUses counts the direct calls and string concatenations on a method's class
// and name in the bytecode, the two facts the cut's justification rests on.
func (a *Analysis) directUses(ref string) (calls, concats int, err error) {
	target, err := a.World.MethodByRef(ref)
	if err != nil {
		return 0, 0, err
	}
	holder := target.Class.Name
	for _, m := range a.World.Methods {
		events, err := m.Events()
		if err != nil {
			return 0, 0, err
		}
		for i := range events {
			ev := &events[i]
			if ev.Kind == EvInvoke && ev.Ref.Class == holder && ev.Ref.Name == target.Name && ev.Ref.Desc == target.Desc {
				calls++
			}
			for _, t := range ev.Args {
				if t == holder {
					concats++
				}
			}
		}
	}
	return calls, concats, nil
}

// checkMethodCuts fails the soundness check when a cut's justification stops being
// true of the bytecode, or when the ternary scan and the source disagree.
func (a *Analysis) checkMethodCuts(rep *SoundnessReport) {
	if n, err := a.sourceNtkCount(); err != nil {
		rep.Errors = append(rep.Errors, err.Error())
	} else if sites := a.ntkSites(); n != len(sites) {
		rep.Errors = append(rep.Errors, fmt.Sprintf("the source has %d ternary sites, the bytecode scan found %d", n, len(sites)))
	}
	for i := range a.Cfg.MethodCut {
		mc := &a.Cfg.MethodCut[i]
		calls, concats, err := a.directUses(mc.Method)
		if err != nil {
			rep.Errors = append(rep.Errors, err.Error())
			continue
		}
		if calls+concats > 0 {
			rep.Errors = append(rep.Errors, fmt.Sprintf("cut %s: %s has %d direct calls and %d string concatenations", mc.Name, mc.Method, calls, concats))
		}
	}
}

func rate(hit, total int) string {
	return fmt.Sprintf("%.1f%% (%d of %d)", Pct(hit, total), hit, total)
}

// writePort is the port section of the report.
func (a *Analysis) writePort(b *strings.Builder, rep *SoundnessReport) error {
	raw := a.tallyScopeView("port", RTA, false)
	view := a.tallyScopeView("port", RTA, true)
	cha := a.tallyScopeView("port", CHA, true)
	konst := a.tallyScopeView("port", RTAConst, true)
	base := a.tallyScopeView("cli-noexp", RTA, false)

	review, err := parseReview(generatedReviewText)
	if err != nil {
		return err
	}
	missed, err := a.generatedMissed("port")
	if err != nil {
		return err
	}
	pr := applyReview(review, missed)
	sites := a.ntkSites()
	srcCount, err := a.sourceNtkCount()
	if err != nil {
		return err
	}
	ntkS, ntkU := countNtk(sites, 0, false), countNtk(sites, 1, false)
	portS, portU := countNtk(sites, 0, true), countNtk(sites, 1, true)

	b.WriteString("## The `port` scope\n\n")
	fmt.Fprintf(b, "**What a port has to match, under RTA: handwritten %s by specs and %s by the union; generated, with the ternary as constant, %s and %s.** Raw scopes stay in the table above. Coverage here means the code ran. It does not mean any output was checked.\n\n",
		rate(view.Hand.Specs, view.Hand.Branches), rate(view.Hand.Union, view.Hand.Branches), rate(view.Gen.Specs, view.Gen.Branches), rate(view.Gen.Union, view.Gen.Branches))
	b.WriteString("`port` is `tla2sany.SANY.main` with three things taken out, each a named line in `scopes.conf` that Chris can veto:\n\n")
	b.WriteString("- `explorer`, a line cut: the `-d` block that creates an `Explorer`.\n- XMLExporter: not an entry; nothing in the command line reaches it.\n- `symboltable-tostring`, a method cut, below.\n\nThe ternary is a reporting view on top, and the map and the soundness check do not use it.\n\n")

	b.WriteString("| code | view | outcomes | specs | union |\n|---|---|---:|---:|---:|\n")
	cell := func(hit, total int) string { return rate(hit, total) }
	row := func(code, label string, t, hs, hu int) {
		fmt.Fprintf(b, "| %s | %s | %d | %s | %s |\n", code, label, t, cell(hs, t), cell(hu, t))
	}
	h := view.Hand
	row("handwritten", "raw", h.Branches, h.Specs, h.Union)
	fmt.Fprintf(b, "| handwritten | without guards on instructions that ran | | %s | %s |\n", cell(h.Specs, h.Branches-h.ExclSpecs.Partly), cell(h.Union, h.Branches-h.ExclUnion.Partly))
	fmt.Fprintf(b, "| handwritten | without guards, also on instructions that never ran | | %s | %s |\n", cell(h.Specs, h.Branches-h.ExclSpecs.Partly-h.ExclSpecs.Never), cell(h.Union, h.Branches-h.ExclUnion.Partly-h.ExclUnion.Never))
	g := raw.Gen
	row("generated", "raw, ternary kept", g.Branches, g.Specs, g.Union)
	g = view.Gen
	row("generated", "ternary constant", g.Branches, g.Specs, g.Union)
	fmt.Fprintf(b, "| generated | ternary constant, without guards on instructions that ran | | %s | %s |\n", cell(g.Specs, g.Branches-g.ExclSpecs.Partly), cell(g.Union, g.Branches-g.ExclUnion.Partly))
	fmt.Fprintf(b, "| generated | ternary constant, without guards, also never ran | | %s | %s |\n", cell(g.Specs, g.Branches-g.ExclSpecs.Partly-g.ExclSpecs.Never), cell(g.Union, g.Branches-g.ExclUnion.Partly-g.ExclUnion.Never))
	row("generated", "ternary constant, without the review's shown dead", g.Branches-pr.Shown, g.Specs, g.Union)
	row("generated", "ternary constant, without shown and believed dead", g.Branches-pr.Shown-pr.Believed, g.Specs, g.Union)
	fmt.Fprintf(b, "\nPrecision: CHA has %d handwritten and %d generated outcomes with the view on, RTA+const %d and %d. Guards are shapes, not proof (next sections), and they overlap the review's dead code, so the two are never subtracted together. Specs hit %d and the union %d handwritten outcomes in `cli-noexp`, against %d and %d here.\n\n",
		cha.Hand.Branches, cha.Gen.Branches, konst.Hand.Branches, konst.Gen.Branches, base.Hand.Specs, base.Hand.Union, view.Hand.Specs, view.Hand.Union)

	cliHand, cliGen, specHand, specGen := a.portPlainCLI()
	fmt.Fprintf(b, "**How much of \"specs\" is the harness.** `specs` is `SanyProbe` plus, on accepted specs, `XMLExporter`, merged per case. Hits inside `port` that only those runs can have made: at most %d of %d handwritten and %d of %d generated, with the ternary kept in both counts. That is an upper bound, method by method: specs minus the plain command-line passes, and those passes were last run before tiers D and E stopped growing, so cases kept since then count here too.\n\n", specHand-cliHand, specHand, specGen-cliGen, specGen)
	if gone := a.cutAwayHits(); len(gone) > 0 {
		fmt.Fprintf(b, "Hits the cut takes with it, because the method is in `cli-noexp` and not in `port`: %s. They ran only through a harness: the plain command-line rows hold to `port` with nothing outside it.\n\n", strings.Join(gone, ", "))
	}

	b.WriteString("### Cutting `SymbolTable.toString()`\n\n")
	for i := range a.Cfg.MethodCut {
		mc := &a.Cfg.MethodCut[i]
		calls, concats, err := a.directUses(mc.Method)
		if err != nil {
			return err
		}
		fmt.Fprintf(b, "`%s` in `scopes.conf`: `%s`. %s\n\n", mc.Name, mc.Method, mc.Why)
		cut := 0
		for _, rr := range rep.Runs {
			if cov, err := ReadCoverage(a.Paths.abs(rr.Path)); err == nil {
				if c := cov.Methods[mc.Method]; c != nil && c.InsnHit > 0 {
					cut++
				}
			}
		}
		fmt.Fprintf(b, "Measured: %d direct calls, %d string concatenations on the type, %d of %d report rows ran it. `staticreach check` fails if any of those three stops being zero. Not shown: that no `SymbolTable` reaches a JDK call that stringifies it, which needs points-to analysis. This is a decision, not a proof. Without it, `cli-noexp` has %d handwritten outcomes under RTA; `port` has %d.\n\n", calls, concats, cut, len(rep.Runs), base.Hand.Branches, view.Hand.Branches)
	}

	b.WriteString("### The `jj_ntk` ternary\n\n")
	fmt.Fprintf(b, "`(jj_ntk==-1)?jj_ntk():jj_ntk`: **%d sites in the bytecode, %d outcomes; the source has %d.** Handled per instruction: each site's two outcomes leave the denominator and each that was hit leaves the numerator. Default off, on only in `port` and `staticreach view`. Basis: `FINDINGS.md`, 0 stale reads in 8,721 cases and 0 mismatches with always recomputing; that is corpus evidence, not a proof, and the view rests on it. Every site, with the sides it took: `ntk-sites.tsv`.\n\n", len(sites), ntkU.Outcomes, srcCount)
	b.WriteString("| | sites | outcomes | hit | missed | both sides | compute only | memo only | neither |\n|---|---:|---:|---:|---:|---:|---:|---:|---:|\n")
	for _, r := range []struct {
		name string
		n    NtkCount
	}{{"all, specs", ntkS}, {"all, union", ntkU}, {"`port`, specs", portS}, {"`port`, union", portU}} {
		fmt.Fprintf(b, "| %s | %d | %d | %d | %d | %d | %d | %d | %d |\n", r.name, r.n.Sites, r.n.Outcomes, r.n.Hit, r.n.Outcomes-r.n.Hit, r.n.Both, r.n.ComputeOnly, r.n.MemoOnly, r.n.Neither)
	}
	b.WriteString("\nThe estimate going in was 177 hit and 83 missed of 260. A counter patched into the parser, in `FINDINGS.md`, saw both sides at 48 sites, compute only at 54, memo only at 27 and neither at 1. JaCoCo's arms above are measured separately, on the current exec files.\n\n")

	b.WriteString("### Generated code: what the review says is dead\n\n")
	fmt.Fprintf(b, "**Judgement, not measured here.** Verdicts are from one reviewer's reading of the generated code, as counts per kind in `cmd/staticreach/generated-review.conf`, and this tool does not recheck them. The per-line verdicts were not kept, so these are not per instruction. Missed outcomes per kind are measured against the union; what `port` leaves out comes off that kind's shown-dead count, because those outcomes have no static path. The ternary sides are the review's believed-dead 81; with the view on they are already out, so they do not count again.\n\n")
	b.WriteString("| kind | missed, all | outside `port` | shown dead | believed dead | ternary, believed dead | unknown | ran, JaCoCo blind | 2 GiB input only |\n|---|---:|---:|---:|---:|---:|---:|---:|---:|\n")
	for _, r := range pr.Rows {
		fmt.Fprintf(b, "| %s | %d | %d | %d | %d | %d | %d | %d | %d |\n", r.Kind, r.Measured+r.Outside, r.Outside, r.Shown, r.Believed, r.Ternary, r.Unknown, r.Blind, r.Big)
	}
	fmt.Fprintf(b, "| **total in `port`** | | | %d | %d | %d | %d | %d | %d |\n\n", pr.Shown, pr.Believed, pr.Ternary, pr.Unknown, pr.Blind, pr.Big)
	if len(pr.Unreconciled) == 0 {
		b.WriteString("Every kind's measured misses equal the review's counts.\n\n")
	} else {
		b.WriteString("Measured misses do not match the review's counts here:\n\n")
		for _, u := range pr.Unreconciled {
			fmt.Fprintf(b, "- %s\n", u)
		}
		b.WriteString("\n")
	}
	return nil
}

// portPlainCLI compares specs with the plain command line, method by method, inside
// port under RTA. Upper bound on what only a harness hit.
func (a *Analysis) portPlainCLI() (cliHand, cliGen, specHand, specGen int) {
	res := a.Results["port"][RTA]
	for _, k := range sortedKeys(a.Union) {
		m := a.methodOf(k)
		if !res.Reached[m.ID] {
			continue
		}
		s, c := a.Specs.Methods[k], a.SpecCLI.Methods[k]
		if s == nil {
			continue
		}
		cli := 0
		if c != nil {
			cli = c.BranchHit
		}
		if cli > s.BranchHit {
			cli = s.BranchHit
		}
		if IsGenerated(m.Class) {
			specGen += s.BranchHit
			cliGen += cli
		} else {
			specHand += s.BranchHit
			cliHand += cli
		}
	}
	return cliHand, cliGen, specHand, specGen
}

// cutAwayHits names methods in cli-noexp and not in port that specs hit.
func (a *Analysis) cutAwayHits() []string {
	base, port := a.Results["cli-noexp"][RTA], a.Results["port"][RTA]
	var out []string
	for _, k := range sortedKeys(a.Union) {
		m := a.methodOf(k)
		if !base.Reached[m.ID] || port.Reached[m.ID] {
			continue
		}
		if s := a.Specs.Methods[k]; s != nil && s.BranchHit > 0 {
			out = append(out, fmt.Sprintf("`%s` %d of %d", short(m), s.BranchHit, s.BranchTotal))
		}
	}
	return out
}

// viewText prints every scope under NtkAll, for `staticreach view`.
func (a *Analysis) viewText() string {
	var b strings.Builder
	b.WriteString("jj_ntk ternary treated as constant in every scope. A reporting view: the map, the scopes and the soundness check do not use it.\n\n")
	b.WriteString("| scope | precision | handwritten | specs | union | generated | specs | union | ternary outcomes out | hit by specs | hit by union |\n|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|\n")
	for i := range a.Cfg.Scopes {
		n := a.Cfg.Scopes[i].Name
		for _, p := range precisions {
			t := a.TallyScope(n, p)
			fmt.Fprintf(&b, "| `%s` | %s | %d | %s | %s | %d | %s | %s | %d | %d | %d |\n", n, p, t.Hand.Branches, p1(t.Hand.Specs, t.Hand.Branches), p1(t.Hand.Union, t.Hand.Branches), t.Gen.Branches, p1(t.Gen.Specs, t.Gen.Branches), p1(t.Gen.Union, t.Gen.Branches), t.NtkTotal, t.NtkSpecs, t.NtkUn)
		}
	}
	return b.String()
}

// portCutArmsTSV lists the outcomes the port scope's line cuts remove, for readers
// that count arms: cmd/casemap. arm is a number, or "all" for a branch on a cut line.
func (a *Analysis) portCutArmsTSV() string {
	var sb strings.Builder
	sb.WriteString(tsvHeader)
	sb.WriteString("class\tname\tdescriptor\tordinal\tline\tarm\n")
	res := a.Results["port"][RTA]
	keys := make([]string, 0, len(a.World.Methods))
	byKey := map[string]*Method{}
	for _, m := range a.World.Methods {
		keys = append(keys, m.Key())
		byKey[m.Key()] = m
	}
	sort.Strings(keys)
	inCut := func(m *Method, line int) bool { return res.cutAway(m, line) }
	for _, k := range keys {
		m := byKey[k]
		if !res.Reached[m.ID] {
			continue
		}
		cuts := false
		for i := range res.Opts.Cuts {
			if res.Opts.Cuts[i].Class == m.Class.Name {
				cuts = true
			}
		}
		if !cuts {
			continue
		}
		insns, err := m.Decode()
		if err != nil {
			fatalf("%v", err)
		}
		for ord, in := range insns {
			if conditionalOutcomes(in) == 0 {
				continue
			}
			line := m.LineAt(in.Off)
			row := func(arm string) {
				fmt.Fprintf(&sb, "%s\t%s\t%s\t%d\t%d\t%s\n", m.Class.Name, m.Name, m.Desc, ord, line, arm)
			}
			switch {
			case inCut(m, line):
				row("all")
			case in.IsConditional() && inCut(m, m.LineAt(in.Off+in.Len)):
				row("0")
			case in.IsConditional() && inCut(m, m.LineAt(in.Targets[0])):
				row("1")
			}
		}
	}
	return sb.String()
}
