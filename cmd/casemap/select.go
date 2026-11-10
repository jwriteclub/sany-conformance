// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Selection, and every case's why. Everything here is a function of the set of cases
// and what each reaches: no keep order, no clock, no random choice. An element is
// something a case can be the reason for: a branch outcome, a line, a method or a
// grammar element from the index, a token kind, or a row of data/elements-extra.tsv
// (error-code combination, SANY-vs-tlx disagreement signature, PIT mutant kill, Java
// crash signature).
//
// The rule, Chris's: for each element upstream's sets do not reach, rarest first, keep
// the smallest case reaching it, then drop what turned out redundant, largest first.
// Plain greedy set cover does the opposite and keeps kitchen sinks.

// universe holds each case's elements as ints and how to name them.
type universe struct {
	ix       *Index
	nParts   int
	nTokens  int
	extra    []string          // names of extra elements
	caseEl   [][]int32         // case to sorted elements
	count    []int32           // element to number of cases reaching it
	upstream []bool            // element reached by an upstream set
	upCase   []bool            // case is in an upstream set
	byID     map[string]int32  // case id to index
	evidence map[string]string // case id to stated purpose
	names    []string          // element names, built once
}

func loadUniverse(p Paths, ix *Index) (*universe, error) {
	u := &universe{ix: ix, nParts: len(ix.Parts), nTokens: len(ix.Tokens), byID: map[string]int32{}, evidence: map[string]string{}}
	for i := range ix.Cases {
		u.byID[ix.Cases[i].ID] = int32(i)
	}
	u.upCase = make([]bool, len(ix.Cases))
	for i := range ix.Cases {
		u.upCase[i] = groupByName(ix.Cases[i].Group).Upstream
	}
	extras := map[string]int32{}
	caseExtra := make([][]int32, len(ix.Cases))
	err := readTSV(p.abs("data/elements-extra.tsv"), func(w []string) error {
		ci, ok := u.byID[w[0]]
		if !ok {
			return fmt.Errorf("elements-extra.tsv: no case %s", w[0])
		}
		e, ok := extras[w[1]]
		if !ok {
			e = int32(len(u.extra))
			extras[w[1]] = e
			u.extra = append(u.extra, w[1])
		}
		caseExtra[ci] = append(caseExtra[ci], e)
		return nil
	})
	if err != nil {
		return nil, err
	}
	if err := readTSV(p.abs("data/evidence.tsv"), func(w []string) error {
		if _, ok := u.byID[w[0]]; !ok {
			return fmt.Errorf("evidence.tsv: no case %s", w[0])
		}
		u.evidence[w[0]] = w[1]
		return nil
	}); err != nil {
		return nil, err
	}
	base := u.nParts + u.nTokens
	total := base + len(u.extra)
	u.count = make([]int32, total)
	u.upstream = make([]bool, total)
	u.caseEl = make([][]int32, len(ix.Cases))
	for ci := range ix.Cases {
		seen := map[int32]bool{}
		var els []int32
		for _, part := range ix.caseParts(ci) {
			if !seen[part] {
				seen[part] = true
				els = append(els, part)
			}
		}
		for _, k := range ix.CaseTokens[ci] {
			e := int32(u.nParts) + int32(k)
			if !seen[e] {
				seen[e] = true
				els = append(els, e)
			}
		}
		for _, x := range caseExtra[ci] {
			e := int32(base) + x
			if !seen[e] {
				seen[e] = true
				els = append(els, e)
			}
		}
		sort.Slice(els, func(a, b int) bool { return els[a] < els[b] })
		u.caseEl[ci] = els
		for _, e := range els {
			u.count[e]++
			if u.upCase[ci] {
				u.upstream[e] = true
			}
		}
	}
	u.names = make([]string, total)
	for e := range u.names {
		u.names[e] = u.label(int32(e))
	}
	return u, nil
}

// label names an element so that a reader can find it. Stable across runs.
func (u *universe) label(e int32) string {
	ix := u.ix
	switch {
	case int(e) < u.nParts:
		pt := &ix.Parts[e]
		switch pt.Kind {
		case PKArm:
			scope := "branch"
			if pt.InPort {
				scope = "port-branch"
			}
			return scope + " " + ix.armLabel(pt)
		case PKLine:
			return fmt.Sprintf("line %s:%d", ix.Classes[pt.Class].File, pt.Line)
		case PKMethod:
			return "method " + ix.methodLabel(pt)
		default:
			return "grammar " + elemLabel(&ix.Elements[pt.Elem])
		}
	case int(e) < u.nParts+u.nTokens:
		t := &ix.Tokens[int(e)-u.nParts]
		return "token " + t.Name
	}
	return u.extra[int(e)-u.nParts-u.nTokens]
}

func readTSV(path string, fn func([]string) error) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 0, 1<<20), 1<<24)
	for sc.Scan() {
		l := sc.Text()
		if l == "" || l[0] == '#' {
			continue
		}
		if err := fn(strings.Split(l, "\t")); err != nil {
			return err
		}
	}
	return sc.Err()
}

// selectCases runs the rule. order is any permutation of the case indices: the result
// must not depend on it.
func (u *universe) selectCases(order []int32) map[int32]bool {
	ix := u.ix
	need := map[int32][]int32{} // element no upstream set reaches to the cases reaching it
	for _, ci := range order {
		if u.upCase[ci] {
			continue
		}
		for _, e := range u.caseEl[ci] {
			if !u.upstream[e] {
				need[e] = append(need[e], ci)
			}
		}
	}
	elems := make([]int32, 0, len(need))
	for e := range need {
		elems = append(elems, e)
	}
	sort.Slice(elems, func(a, b int) bool {
		x, y := elems[a], elems[b]
		if len(need[x]) != len(need[y]) {
			return len(need[x]) < len(need[y])
		}
		return u.names[x] < u.names[y]
	})
	foot := func(ci int32) int {
		n := 0
		for _, e := range u.caseEl[ci] {
			if !u.upstream[e] {
				n++
			}
		}
		return n
	}
	footprint := map[int32]int{}
	for _, ci := range order {
		if !u.upCase[ci] {
			footprint[ci] = foot(ci)
		}
	}
	better := func(a, b int32) bool {
		ca, cb := &ix.Cases[a], &ix.Cases[b]
		if ca.Size != cb.Size {
			return ca.Size < cb.Size
		}
		if footprint[a] != footprint[b] {
			return footprint[a] < footprint[b]
		}
		return ca.ID < cb.ID
	}
	chosen := map[int32]bool{}
	covered := map[int32]int{}
	take := func(ci int32) {
		if chosen[ci] {
			return
		}
		chosen[ci] = true
		for _, e := range u.caseEl[ci] {
			if !u.upstream[e] {
				covered[e]++
			}
		}
	}
	// Evidence first: its purpose is stated, not measured, so it stays whatever it reaches.
	for _, ci := range order {
		if _, ok := u.evidence[ix.Cases[ci].ID]; ok && !u.upCase[ci] {
			take(ci)
		}
	}
	for _, e := range elems {
		if covered[e] > 0 {
			continue
		}
		best := need[e][0]
		for _, c := range need[e][1:] {
			if better(c, best) {
				best = c
			}
		}
		take(best)
	}
	// Drop what turned out redundant, largest first; never evidence.
	list := make([]int32, 0, len(chosen))
	for ci := range chosen {
		list = append(list, ci)
	}
	sort.Slice(list, func(a, b int) bool { return better(list[b], list[a]) })
	for _, ci := range list {
		if _, ok := u.evidence[ix.Cases[ci].ID]; ok {
			continue
		}
		needed := false
		for _, e := range u.caseEl[ci] {
			if !u.upstream[e] && covered[e] == 1 {
				needed = true
				break
			}
		}
		if !needed {
			delete(chosen, ci)
			for _, e := range u.caseEl[ci] {
				if !u.upstream[e] {
					covered[e]--
				}
			}
		}
	}
	return chosen
}

// WhyRow is one case's reason, independent of any order.
type WhyRow struct {
	ID     string `json:"id"`
	Set    string `json:"set"`
	Place  string `json:"place"` // upstream, selected, deep-cut
	Reason string `json:"reason"`
	// OnlyThisCase counts what no other case in the suite reaches; BeyondUpstream what no upstream set reaches.
	OnlyThisCase   int      `json:"only_this_case"`
	BeyondUpstream int      `json:"beyond_upstream"`
	Sole           []string `json:"sole,omitempty"`
	Rarest         []string `json:"rarest,omitempty"`
	Keeps          []string `json:"keeps,omitempty"` // for a selected case, what no other selected or upstream case reaches
	KeepsMore      int      `json:"keeps_more,omitempty"`
}

const whyListed = 6

func (u *universe) whyRows(chosen map[int32]bool) []WhyRow {
	ix := u.ix
	covered := make([]int32, len(u.count))
	for ci := range chosen {
		for _, e := range u.caseEl[ci] {
			covered[e]++
		}
	}
	out := make([]WhyRow, len(ix.Cases))
	for ci := range ix.Cases {
		c := &ix.Cases[ci]
		r := WhyRow{ID: c.ID, Set: c.Group}
		var sole, keeps, rare []int32
		for _, e := range u.caseEl[ci] {
			if u.count[e] == 1 {
				r.OnlyThisCase++
				sole = append(sole, e)
			}
			if !u.upstream[e] {
				r.BeyondUpstream++
				rare = append(rare, e)
				if chosen[int32(ci)] && covered[e] == 1 {
					keeps = append(keeps, e)
				}
			}
		}
		byRare := func(l []int32) {
			sort.Slice(l, func(a, b int) bool {
				if u.count[l[a]] != u.count[l[b]] {
					return u.count[l[a]] < u.count[l[b]]
				}
				return u.names[l[a]] < u.names[l[b]]
			})
		}
		byRare(sole)
		byRare(rare)
		byRare(keeps)
		for i, e := range sole {
			if i == whyListed {
				break
			}
			r.Sole = append(r.Sole, u.names[e])
		}
		for i, e := range rare {
			if i == whyListed {
				break
			}
			r.Rarest = append(r.Rarest, fmt.Sprintf("%s (%d cases)", u.names[e], u.count[e]))
		}
		for i, e := range keeps {
			if i == whyListed {
				r.KeepsMore = len(keeps) - whyListed
				break
			}
			r.Keeps = append(r.Keeps, u.names[e])
		}
		switch {
		case u.upCase[ci]:
			r.Place, r.Reason = "upstream", "upstream's own: upstream sets are kept whole"
		case u.evidence[c.ID] != "":
			r.Place, r.Reason = "selected", "evidence: "+u.evidence[c.ID]
		case chosen[int32(ci)]:
			r.Place = "selected"
			r.Reason = fmt.Sprintf("smallest route to %d element(s) no upstream or other selected case reaches", len(keeps))
		default:
			r.Place = "deep-cut"
			if r.BeyondUpstream == 0 {
				r.Reason = "reaches nothing upstream's sets do not; kept as extra data"
			} else {
				r.Reason = fmt.Sprintf("reaches %d element(s) beyond upstream, all of them also reached by selected cases; kept as extra data", r.BeyondUpstream)
			}
		}
		out[ci] = r
	}
	return out
}

func cmdSelect(p Paths, args []string) error {
	fs := flag.NewFlagSet("select", flag.ExitOnError)
	_ = fs.Parse(args)
	ix, err := LoadIndex(p)
	if err != nil {
		return err
	}
	u, err := loadUniverse(p, ix)
	if err != nil {
		return err
	}
	order := make([]int32, len(ix.Cases))
	for i := range order {
		order[i] = int32(i)
	}
	chosen := u.selectCases(order)
	f, err := os.Create(p.abs("data/why.jsonl"))
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	if _, err := w.WriteString("{\"copyright\":\"Copyright 2026 Christopher O'Connell\",\"spdx\":\"Apache-2.0\",\"generated_by\":\"casemap select\"}\n"); err != nil {
		return err
	}
	rows := u.whyRows(chosen)
	per := map[string][3]int{}
	for i := range rows {
		if err := enc.Encode(&rows[i]); err != nil {
			return err
		}
		n := per[rows[i].Set]
		switch rows[i].Place {
		case "upstream":
			n[0]++
		case "selected":
			n[1]++
		default:
			n[2]++
		}
		per[rows[i].Set] = n
	}
	for _, g := range Groups {
		n := per[g.Name]
		fmt.Printf("%-32s upstream %5d selected %5d deep-cut %5d\n", g.Name, n[0], n[1], n[2])
	}
	return nil
}

// tally is reached counts by the kinds the post's table uses.
var tallyKinds = []string{"port handwritten", "port generated", "port rules", "overall handwritten", "overall generated", "overall rules", "lines", "methods", "lookahead outcomes", "token kinds", "error-code combinations", "tlx disagreements", "mutant kills, primary", "mutant kills, wide", "Java crash signatures"}

// kindsOf says which table rows an element counts in.
func (u *universe) kindsOf(e int32) []string {
	ix := u.ix
	switch {
	case int(e) < u.nParts:
		pt := &ix.Parts[e]
		switch pt.Kind {
		case PKArm:
			gen := ix.Classes[pt.Class].Generated
			var out []string
			if gen {
				out = append(out, "overall generated")
				if pt.InPort {
					out = append(out, "port generated")
				}
			} else {
				out = append(out, "overall handwritten")
				if pt.InPort {
					out = append(out, "port handwritten")
				}
			}
			return out
		case PKLine:
			return []string{"lines"}
		case PKMethod:
			return []string{"methods"}
		default:
			if !IsRule(ix.Elements[pt.Elem].Kind) {
				return []string{"lookahead outcomes"}
			}
			out := []string{"overall rules"}
			if pt.InPort {
				out = append(out, "port rules")
			}
			return out
		}
	case int(e) < u.nParts+u.nTokens:
		return []string{"token kinds"}
	}
	n := u.extra[int(e)-u.nParts-u.nTokens]
	switch {
	case strings.HasPrefix(n, "combo "):
		return []string{"error-code combinations"}
	case strings.HasPrefix(n, "tlx "):
		return []string{"tlx disagreements"}
	case strings.HasPrefix(n, "mutant-wide "):
		return []string{"mutant kills, wide"}
	case strings.HasPrefix(n, "mutant "):
		return []string{"mutant kills, primary"}
	case strings.HasPrefix(n, "crash "):
		return []string{"Java crash signatures"}
	}
	return []string{"other"}
}

// cmdProve reruns the accounting: the selected set plus the upstream sets reach what
// the whole suite reaches, and the selection does not depend on the order of the cases.
func cmdProve(p Paths, args []string) error {
	fs := flag.NewFlagSet("prove", flag.ExitOnError)
	_ = fs.Parse(args)
	ix, err := LoadIndex(p)
	if err != nil {
		return err
	}
	u, err := loadUniverse(p, ix)
	if err != nil {
		return err
	}
	order := make([]int32, len(ix.Cases))
	for i := range order {
		order[i] = int32(i)
	}
	chosen := u.selectCases(order)
	// The rule must not depend on order: the same answer reversed and shuffled.
	rev := make([]int32, len(order))
	for i := range order {
		rev[i] = order[len(order)-1-i]
	}
	shuf := append([]int32(nil), order...)
	rand.New(rand.NewSource(20261105)).Shuffle(len(shuf), func(i, j int) { shuf[i], shuf[j] = shuf[j], shuf[i] })
	for name, o := range map[string][]int32{"reversed": rev, "shuffled": shuf} {
		other := u.selectCases(o)
		if len(other) != len(chosen) {
			return fmt.Errorf("selection differs when cases are %s: %d against %d", name, len(other), len(chosen))
		}
		for ci := range chosen {
			if !other[ci] {
				return fmt.Errorf("selection differs when cases are %s: %s", name, ix.Cases[ci].ID)
			}
		}
	}
	fmt.Println("order: the same selection from manifest order, reversed and shuffled")
	// The committed why.jsonl must be this selection.
	stored := map[string]string{}
	if err := readWhy(p, func(r WhyRow) { stored[r.ID] = r.Place }); err != nil {
		return err
	}
	for i := range ix.Cases {
		want := "deep-cut"
		if u.upCase[i] {
			want = "upstream"
		} else if chosen[int32(i)] {
			want = "selected"
		}
		if stored[ix.Cases[i].ID] != want {
			return fmt.Errorf("data/why.jsonl says %s is %q, the rule says %q", ix.Cases[i].ID, stored[ix.Cases[i].ID], want)
		}
	}
	whole := make([]bool, len(u.count))
	kept := make([]bool, len(u.count))
	for ci := range ix.Cases {
		inKept := u.upCase[ci] || chosen[int32(ci)]
		for _, e := range u.caseEl[ci] {
			whole[e] = true
			if inKept {
				kept[e] = true
			}
		}
	}
	type row struct{ whole, kept int }
	tot := map[string]*row{}
	for _, k := range tallyKinds {
		tot[k] = &row{}
	}
	missing := 0
	for e := range whole {
		if !whole[e] {
			continue
		}
		for _, k := range u.kindsOf(int32(e)) {
			r := tot[k]
			if r == nil {
				r = &row{}
				tot[k] = r
			}
			r.whole++
			if kept[e] {
				r.kept++
			}
		}
		if !kept[e] {
			missing++
		}
	}
	fmt.Println("| reached | whole suite | upstream + selected |")
	fmt.Println("|---|---:|---:|")
	for _, k := range tallyKinds {
		r := tot[k]
		fmt.Printf("| %s | %d | %d |\n", k, r.whole, r.kept)
	}
	fmt.Println("elements the whole suite reaches and upstream plus selected do not: " + strconv.Itoa(missing))
	if missing > 0 {
		return fmt.Errorf("the selection lost %d elements", missing)
	}
	return nil
}

func readWhy(p Paths, fn func(WhyRow)) error {
	f, err := os.Open(p.abs("data/why.jsonl"))
	if err != nil {
		return err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 0, 1<<20), 1<<24)
	for sc.Scan() {
		var r WhyRow
		if err := json.Unmarshal(sc.Bytes(), &r); err != nil {
			return err
		}
		if r.ID == "" {
			continue // the header line
		}
		fn(r)
	}
	return sc.Err()
}
