// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func (ix *Index) groupIndex(name string) int {
	for i, g := range ix.Groups {
		if g == name {
			return i
		}
	}
	return -1
}

func (ix *Index) findCase(id string) int {
	for i := range ix.Cases {
		if ix.Cases[i].ID == id {
			return i
		}
	}
	return -1
}

// caseParts is every part one case reaches, in part order.
func (ix *Index) caseParts(ci int) []int32 {
	var out []int32
	for _, pn := range ix.CasePats[ci] {
		out = append(out, ix.patternParts(&ix.Patterns[pn])...)
	}
	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
	return out
}

func simple(class string) string {
	return class[strings.LastIndexByte(class, '/')+1:]
}

func pkgOf(class string) string {
	return class[:strings.LastIndexByte(class, '/')]
}

// armLabel names an arm so that a reader can find it.
func (ix *Index) armLabel(pt *Part) string {
	c := &ix.Classes[pt.Class]
	return fmt.Sprintf("%s:%d %s #%d arm %d", c.File, pt.Line, c.Methods[pt.Method].Name, pt.Ord, pt.Arm)
}

func (ix *Index) methodLabel(pt *Part) string {
	c := &ix.Classes[pt.Class]
	m := &c.Methods[pt.Method]
	overload := 0
	for i := range c.Methods {
		if c.Methods[i].Name == m.Name {
			overload++
		}
	}
	if overload > 1 {
		return simple(c.Name) + "." + m.Name + m.Desc
	}
	return simple(c.Name) + "." + m.Name
}

func elemLabel(e *ElemInfo) string {
	if e.Kind == EKProduction {
		return e.Prod
	}
	return fmt.Sprintf("%s %s %s: %s", e.Prod, e.Path, e.Kind, e.Text)
}

func cmdCase(p Paths, args []string) error {
	fs := flag.NewFlagSet("case", flag.ExitOnError)
	summary := fs.Bool("summary", false, "counts only")
	_ = fs.Parse(args)
	if fs.NArg() != 1 {
		return fmt.Errorf("usage: casemap case [-summary] <id>")
	}
	ix, err := LoadIndex(p)
	if err != nil {
		return err
	}
	ci := ix.findCase(fs.Arg(0))
	if ci < 0 {
		return fmt.Errorf("no case %q", fs.Arg(0))
	}
	fmt.Print(ix.caseText(ci, *summary))
	return readWhy(p, func(r WhyRow) {
		if r.ID != ix.Cases[ci].ID {
			return
		}
		fmt.Printf("\nplace: %s\nwhy: %s\n", r.Place, r.Reason)
		for _, l := range r.Keeps {
			fmt.Printf("  keeps: %s\n", l)
		}
		for _, l := range r.Rarest {
			fmt.Printf("  rarest: %s\n", l)
		}
	})
}

// caseText answers "case to parts".
func (ix *Index) caseText(ci int, summary bool) string {
	var b strings.Builder
	c := &ix.Cases[ci]
	group := ix.groupIndex(c.Group)
	bit := uint64(1) << group
	fmt.Fprintf(&b, "case %s\nsource: %s (%s), tier %s\n", c.ID, c.Group, c.Source, c.Tier)
	fmt.Fprintf(&b, "files: %d, %d bytes, root %s\norigin: %s\n", len(c.Files), c.Size, c.RootFile, c.Origin)
	if c.Commit != "" {
		fmt.Fprintf(&b, "commit: %s\n", c.Commit)
	}
	fmt.Fprintf(&b, "asserts: %s\n", c.Expect)
	fmt.Fprintf(&b, "kept: %s", c.Keep)
	if len(c.Reasons) > 0 {
		fmt.Fprintf(&b, " (%s)", strings.Join(c.Reasons, ", "))
	}
	b.WriteString("\n")
	for _, kv := range [][2]string{{"model", c.Model}, {"target", c.Target}, {"aim", c.Aim}, {"parent", c.Parent}, {"rule", c.Rule}, {"finding", c.Finding}} {
		if kv[1] != "" {
			fmt.Fprintf(&b, "%s: %s\n", kv[0], kv[1])
		}
	}
	outcome := c.Outcome
	if outcome == "" {
		outcome = "not recorded per case"
	}
	fmt.Fprintf(&b, "result in the coverage run, not the stock check: outcome %s, verdict %q, codes %q, warnings %q\n", outcome, c.L0, c.L1e, c.L1w)
	b.WriteString("note: coverage means the code ran under the measuring wrapper and, on accepted specs, XMLExporter. It does not mean output was checked.\n\n")

	parts := ix.caseParts(ci)
	type section struct {
		title string
		items []string
	}
	secs := map[string]*section{}
	order := []string{}
	add := func(key, title, item string) {
		s := secs[key]
		if s == nil {
			s = &section{title: title}
			secs[key] = s
			order = append(order, key)
		}
		s.items = append(s.items, item)
	}
	pkgs := map[string]bool{}
	classes := map[string]bool{}
	for _, pid := range parts {
		pt := &ix.Parts[pid]
		cl := &ix.Classes[pt.Class]
		pkgs[pkgOf(cl.Name)] = true
		classes[cl.Name] = true
		switch pt.Kind {
		case PKElem:
			e := &ix.Elements[pt.Elem]
			key := "g-" + e.Kind
			add(key, "grammar: "+e.Kind, elemLabel(e))
			if pt.Reach == 1 {
				add("u-case-g", "grammar reached only by this case", elemLabel(e))
			} else if pt.Groups == bit {
				add("u-src-g", "grammar reached only by this source", elemLabel(e))
			}
		case PKMethod:
			add("c-method", "code: methods", ix.methodLabel(pt))
			if pt.Reach == 1 {
				add("u-case-m", "methods reached only by this case", ix.methodLabel(pt))
			} else if pt.Groups == bit {
				add("u-src-m", "methods reached only by this source", ix.methodLabel(pt))
			}
		case PKLine:
			if pt.Reach == 1 {
				add("u-case-l", "lines reached only by this case", fmt.Sprintf("%s:%d", cl.File, pt.Line))
			} else if pt.Groups == bit {
				add("u-src-l", "lines reached only by this source", fmt.Sprintf("%s:%d", cl.File, pt.Line))
			}
		case PKArm:
			add("c-arm", "code: branch outcomes", ix.armLabel(pt))
			if pt.Reach == 1 {
				add("u-case-a", "branch outcomes reached only by this case", ix.armLabel(pt))
			} else if pt.Groups == bit {
				add("u-src-a", "branch outcomes reached only by this source", ix.armLabel(pt))
			}
		}
	}
	var tok []string
	for _, k := range ix.CaseTokens[ci] {
		t := &ix.Tokens[k]
		tok = append(tok, t.Const+" "+t.Name)
		if t.Reach == 1 {
			add("u-case-t", "token kinds only in this case", t.Name)
		} else if t.Groups == bit {
			add("u-src-t", "token kinds only in this source", t.Name)
		}
	}
	secs["tok"] = &section{title: "grammar: token kinds (from the parser's token chain, not JaCoCo; the case's own files)", items: tok}
	secs["pkg"] = &section{title: "code: packages", items: keysOf(pkgs)}
	secs["cls"] = &section{title: "code: classes", items: keysOf(classes)}
	for _, key := range []string{"g-production", "g-alternative", "g-optional", "g-optional-skip", "g-loop-enter", "g-loop-exit", "g-syntax-error", "g-lookahead-true", "g-lookahead-false", "tok", "pkg", "cls", "c-method", "c-arm",
		"u-case-g", "u-case-t", "u-case-m", "u-case-l", "u-case-a", "u-src-g", "u-src-t", "u-src-m", "u-src-l", "u-src-a"} {
		s := secs[key]
		if s == nil {
			if strings.HasPrefix(key, "u-") {
				continue
			}
			s = &section{title: titleFor(key)}
		}
		fmt.Fprintf(&b, "%s: %d\n", s.title, len(s.items))
		if !summary {
			for _, it := range s.items {
				fmt.Fprintf(&b, "  %s\n", it)
			}
		}
	}
	return b.String()
}

func titleFor(key string) string {
	switch {
	case strings.HasPrefix(key, "g-"):
		return "grammar: " + key[2:]
	case key == "c-method":
		return "code: methods"
	case key == "c-arm":
		return "code: branch outcomes"
	}
	return key
}

func keysOf(m map[string]bool) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

var reFileLine = regexp.MustCompile(`^(.+):(\d+)$`)

// matchParts resolves a who query to the parts it names.
func (ix *Index) matchParts(q string) (parts []int32, desc string, err error) {
	if m := reFileLine.FindStringSubmatch(q); m != nil {
		line, _ := strconv.Atoi(m[2])
		for pid := range ix.Parts {
			pt := &ix.Parts[pid]
			if pt.Kind == PKLine && int(pt.Line) == line && fileMatches(ix.Classes[pt.Class].File, m[1]) {
				parts = append(parts, int32(pid))
			}
		}
		if len(parts) == 0 {
			return nil, "", fmt.Errorf("no counted line %s", q)
		}
		return parts, "line " + q, nil
	}
	for _, name := range ix.Productions {
		if name == q {
			for i := range ix.Elements {
				if ix.Elements[i].Prod == q && ix.Elements[i].Kind == EKProduction {
					return []int32{ix.Elements[i].Part}, "production " + q, nil
				}
			}
		}
	}
	if prod, n, ok := strings.Cut(q, "#"); ok {
		i, e := strconv.Atoi(n)
		if e == nil && i > 0 {
			k := 0
			for j := range ix.Elements {
				el := &ix.Elements[j]
				if el.Prod == prod && el.Kind != EKProduction && IsRule(el.Kind) {
					k++
					if k == i {
						return []int32{el.Part}, elemLabel(el), nil
					}
				}
			}
		}
		return nil, "", fmt.Errorf("no element %s; `casemap grammar -prod %s` lists them", q, prod)
	}
	if dot := strings.LastIndexByte(q, '.'); dot > 0 {
		cls, meth := q[:dot], q[dot+1:]
		for ci := range ix.Classes {
			c := &ix.Classes[ci]
			if c.Name != cls && simple(c.Name) != cls {
				continue
			}
			for mi := range c.Methods {
				if c.Methods[mi].Name == meth {
					parts = append(parts, c.Methods[mi].Part)
				}
			}
		}
		if len(parts) > 0 {
			return parts, "method " + q, nil
		}
	}
	return nil, "", fmt.Errorf("%q is not a production, Production#n, Class.method, file:line or token:NAME", q)
}

func fileMatches(file, q string) bool {
	return file == q || strings.HasSuffix(file, "/"+q)
}

// casesOf lists the cases reaching any of the parts, in index order.
func (ix *Index) casesOf(parts []int32) []int32 {
	seen := map[int32]bool{}
	for _, pid := range parts {
		pt := &ix.Parts[pid]
		cls := &ix.Classes[pt.Class]
		var probes []int
		for i := range cls.Images {
			for _, q := range cls.Images[i] {
				if q == pid {
					probes = append(probes, i)
					break
				}
			}
		}
		for _, pn := range cls.Patterns {
			pat := &ix.Patterns[pn]
			for _, pr := range probes {
				if hasBit(pat.Bits, pr) {
					for _, ci := range pat.Cases {
						seen[ci] = true
					}
					break
				}
			}
		}
	}
	out := make([]int32, 0, len(seen))
	for c := range seen {
		out = append(out, c)
	}
	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
	return out
}

func cmdWho(p Paths, args []string) error {
	fs := flag.NewFlagSet("who", flag.ExitOnError)
	limit := fs.Int("limit", 5, "cases to list per source")
	all := fs.Bool("all", false, "list every case")
	_ = fs.Parse(args)
	if fs.NArg() != 1 {
		return fmt.Errorf("usage: casemap who [-limit n] [-all] <production | Production#n | Class.method | file:line | token:NAME>")
	}
	ix, err := LoadIndex(p)
	if err != nil {
		return err
	}
	n := *limit
	if *all {
		n = 1 << 30
	}
	out, err := ix.whoText(fs.Arg(0), n)
	if err != nil {
		return err
	}
	fmt.Print(out)
	return nil
}

// whoText answers "part to cases": smallest first, grouped by source.
func (ix *Index) whoText(q string, limit int) (string, error) {
	var cases []int32
	var desc string
	extra := ""
	if name, ok := strings.CutPrefix(q, "token:"); ok {
		k := -1
		for i := range ix.Tokens {
			if ix.Tokens[i].Const == name || ix.Tokens[i].Name == name || ix.Tokens[i].Name == "\""+name+"\"" || ix.Tokens[i].Name == "<"+name+">" {
				k = i
				break
			}
		}
		if k < 0 {
			return "", fmt.Errorf("no token kind %s", name)
		}
		for ci := range ix.CaseTokens {
			for _, t := range ix.CaseTokens[ci] {
				if int(t) == k {
					cases = append(cases, int32(ci))
					break
				}
			}
		}
		desc = "token kind " + ix.Tokens[k].Name + " (" + ix.Tokens[k].Const + "), from the parser's token chain"
	} else {
		parts, d, err := ix.matchParts(q)
		if err != nil {
			return "", err
		}
		desc = d
		cases = ix.casesOf(parts)
		if m := reFileLine.FindStringSubmatch(q); m != nil {
			line, _ := strconv.Atoi(m[2])
			var sb strings.Builder
			for pid := range ix.Parts {
				pt := &ix.Parts[pid]
				if pt.Kind == PKArm && int(pt.Line) == line && fileMatches(ix.Classes[pt.Class].File, m[1]) {
					fmt.Fprintf(&sb, "  %s: %d cases\n", ix.armLabel(pt), pt.Reach)
				}
			}
			extra = sb.String()
		}
	}
	var b strings.Builder
	fmt.Fprintf(&b, "%s: %d cases\n%s", desc, len(cases), extra)
	byGroup := map[string][]int32{}
	for _, c := range cases {
		byGroup[ix.Cases[c].Group] = append(byGroup[ix.Cases[c].Group], c)
	}
	for _, g := range ix.Groups {
		list := byGroup[g]
		if len(list) == 0 {
			continue
		}
		sort.Slice(list, func(i, j int) bool {
			a, b := &ix.Cases[list[i]], &ix.Cases[list[j]]
			if a.Size != b.Size {
				return a.Size < b.Size
			}
			return a.ID < b.ID
		})
		fmt.Fprintf(&b, "%s: %d\n", g, len(list))
		for i, c := range list {
			if i == limit {
				fmt.Fprintf(&b, "  ... %d more\n", len(list)-limit)
				break
			}
			fmt.Fprintf(&b, "  %7d  %s\n", ix.Cases[c].Size, ix.Cases[c].ID)
		}
	}
	return b.String(), nil
}

// Count is one number three ways: reached by a source at all, first reached by it in
// keep order, and reached by it alone, whatever the order.
type Count struct {
	Reached, First, Only int
}

// SourceRow is one line of the per-source table.
type SourceRow struct {
	Group, Category, Title string
	Cases, Distinct, New   int
	Kept                   map[string]int
	PortHand, PortGen      Count
	PortRules              Count
	AllHand, AllGen        Count
	AllRules               Count
}

// Totals are the denominators and what the whole suite reaches.
type Totals struct {
	PortHand, PortGen, PortRules, AllHand, AllGen, AllRules                   int
	PortHandHit, PortGenHit, PortRulesHit, AllHandHit, AllGenHit, AllRulesHit int
	Cases, Distinct                                                           int
}

// SourceTable computes the table. In keep order, each case comes after every case
// that could have earned its parts first: tier A, B, C in manifest order, then the
// kept cases by keep number, then what the gate did not keep.
func (ix *Index) SourceTable() ([]SourceRow, Totals) {
	rows := make([]SourceRow, len(ix.Groups))
	idx := map[string]int{}
	for i, g := range ix.Groups {
		gr := groupByName(g)
		rows[i] = SourceRow{Group: g, Category: gr.Category, Title: gr.Title, Kept: map[string]int{}}
		idx[g] = i
	}
	var tot Totals
	seen := map[string]bool{}
	within := make([]map[string]bool, len(ix.Groups))
	for i := range within {
		within[i] = map[string]bool{}
	}
	for ci := range ix.Cases {
		c := &ix.Cases[ci]
		g := idx[c.Group]
		r := &rows[g]
		r.Cases++
		r.Kept[c.Keep]++
		tot.Cases++
		if !within[g][c.Hash] {
			within[g][c.Hash] = true
			r.Distinct++
		}
		if !seen[c.Hash] {
			seen[c.Hash] = true
			r.New++
			tot.Distinct++
		}
	}
	bump := func(pt *Part, get func(*SourceRow) *Count) {
		for g := range ix.Groups {
			if pt.Groups&(1<<g) != 0 {
				get(&rows[g]).Reached++
			}
		}
		if pt.First >= 0 {
			get(&rows[idx[ix.Cases[pt.First].Group]]).First++
		}
		if pt.Groups != 0 && pt.Groups&(pt.Groups-1) == 0 {
			for g := range ix.Groups {
				if pt.Groups == 1<<g {
					get(&rows[g]).Only++
				}
			}
		}
	}
	for pid := range ix.Parts {
		pt := &ix.Parts[pid]
		hit := 0
		if pt.Reach > 0 {
			hit = 1
		}
		switch pt.Kind {
		case PKArm:
			gen := ix.Classes[pt.Class].Generated
			if gen {
				tot.AllGen++
				tot.AllGenHit += hit
				bump(pt, func(r *SourceRow) *Count { return &r.AllGen })
			} else {
				tot.AllHand++
				tot.AllHandHit += hit
				bump(pt, func(r *SourceRow) *Count { return &r.AllHand })
			}
			if pt.InPort {
				if gen {
					tot.PortGen++
					tot.PortGenHit += hit
					bump(pt, func(r *SourceRow) *Count { return &r.PortGen })
				} else {
					tot.PortHand++
					tot.PortHandHit += hit
					bump(pt, func(r *SourceRow) *Count { return &r.PortHand })
				}
			}
		case PKElem:
			if !IsRule(ix.Elements[pt.Elem].Kind) {
				continue
			}
			tot.AllRules++
			tot.AllRulesHit += hit
			bump(pt, func(r *SourceRow) *Count { return &r.AllRules })
			if pt.InPort {
				tot.PortRules++
				tot.PortRulesHit += hit
				bump(pt, func(r *SourceRow) *Count { return &r.PortRules })
			}
		}
	}
	return rows, tot
}

// CategoryRow sums one kind of source: what any of its sources reaches, and what
// only its sources reach.
type CategoryRow struct {
	Key, Title                                                            string
	Cases                                                                 int
	PortHand, PortGen, PortRules, OnlyPortHand, OnlyPortGen, OnlyPortRule int
	AllHand, AllGen, OnlyAllHand, OnlyAllGen                              int
}

// CategoryTable computes the category sums.
func (ix *Index) CategoryTable() []CategoryRow {
	masks := make([]uint64, len(Categories))
	rows := make([]CategoryRow, len(Categories))
	for i, cat := range Categories {
		rows[i] = CategoryRow{Key: cat.Key, Title: cat.Title}
		for g, name := range ix.Groups {
			if groupByName(name).Category == cat.Key {
				masks[i] |= 1 << g
			}
		}
	}
	for ci := range ix.Cases {
		g := ix.groupIndex(ix.Cases[ci].Group)
		for i := range masks {
			if masks[i]&(1<<g) != 0 {
				rows[i].Cases++
			}
		}
	}
	for pid := range ix.Parts {
		pt := &ix.Parts[pid]
		if pt.Reach == 0 || (pt.Kind != PKArm && pt.Kind != PKElem) {
			continue
		}
		if pt.Kind == PKElem && !IsRule(ix.Elements[pt.Elem].Kind) {
			continue
		}
		gen := ix.Classes[pt.Class].Generated
		for i := range masks {
			any := pt.Groups&masks[i] != 0
			only := pt.Groups&^masks[i] == 0
			r := &rows[i]
			if !any {
				continue
			}
			switch {
			case pt.Kind == PKElem && pt.InPort:
				r.PortRules++
				if only {
					r.OnlyPortRule++
				}
			case pt.Kind == PKArm && gen:
				r.AllGen++
				if only {
					r.OnlyAllGen++
				}
				if pt.InPort {
					r.PortGen++
					if only {
						r.OnlyPortGen++
					}
				}
			case pt.Kind == PKArm:
				r.AllHand++
				if only {
					r.OnlyAllHand++
				}
				if pt.InPort {
					r.PortHand++
					if only {
						r.OnlyPortHand++
					}
				}
			}
		}
	}
	return rows
}

func categoriesMarkdown(rows []CategoryRow) string {
	var b strings.Builder
	b.WriteString("Each cell: reached by any source of this kind / reached by no other kind.\n\n")
	b.WriteString("| kind of source | cases | `port` handwritten | `port` generated | `port` rules | overall handwritten | overall generated |\n|---|---:|---:|---:|---:|---:|---:|\n")
	for _, r := range rows {
		fmt.Fprintf(&b, "| %s | %d | %d / %d | %d / %d | %d / %d | %d / %d | %d / %d |\n", r.Title, r.Cases, r.PortHand, r.OnlyPortHand, r.PortGen, r.OnlyPortGen, r.PortRules, r.OnlyPortRule, r.AllHand, r.OnlyAllHand, r.AllGen, r.OnlyAllGen)
	}
	return b.String()
}

func cmdSources(p Paths, args []string) error {
	fs := flag.NewFlagSet("sources", flag.ExitOnError)
	_ = fs.Parse(args)
	ix, err := LoadIndex(p)
	if err != nil {
		return err
	}
	places := map[string]map[string]int{}
	if err := readWhy(p, func(r WhyRow) {
		if places[r.Set] == nil {
			places[r.Set] = map[string]int{}
		}
		places[r.Set][r.Place]++
	}); err != nil {
		return err
	}
	rows, tot := ix.SourceTable()
	fmt.Print(sourcesMarkdown(rows, tot, places))
	fmt.Print("\n" + categoriesMarkdown(ix.CategoryTable()))
	return nil
}

func cnt(c Count) string { return fmt.Sprintf("%d / %d", c.Reached, c.Only) }

func sourcesMarkdown(rows []SourceRow, tot Totals, places map[string]map[string]int) string {
	var b strings.Builder
	b.WriteString("Each cell: reached by this set / reached only by it, whatever the order. Branch outcomes and grammar rules. `port` is the scope in `data/staticreach/candidate-map.md`, with the jj_ntk ternary out; overall is all of `tla2sany`, ternary in. Reached means the code ran under the measuring wrapper; it does not mean an output was checked. A set that copies another's files shows nothing in \"only\", as the rfcs and tree-sitter corpora do.\n\n")
	b.WriteString("| set | cases | selected | deep cuts | distinct | `port` handwritten | `port` generated | `port` rules | overall handwritten | overall generated | overall rules |\n|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|\n")
	for _, cat := range Categories {
		fmt.Fprintf(&b, "| **%s** | | | | | | | | | | |\n", cat.Title)
		for i := range rows {
			r := &rows[i]
			if r.Category != cat.Key {
				continue
			}
			pl := places[r.Group]
			sel := pl["selected"] + pl["upstream"]
			fmt.Fprintf(&b, "| `%s` | %d | %d | %d | %d | %s | %s | %s | %s | %s | %s |\n", r.Group, r.Cases, sel, pl["deep-cut"], r.Distinct, cnt(r.PortHand), cnt(r.PortGen), cnt(r.PortRules), cnt(r.AllHand), cnt(r.AllGen), cnt(r.AllRules))
		}
	}
	fmt.Fprintf(&b, "| **whole suite, reached of total** | %d | | | %d | %d of %d | %d of %d | %d of %d | %d of %d | %d of %d | %d of %d |\n", tot.Cases, tot.Distinct, tot.PortHandHit, tot.PortHand, tot.PortGenHit, tot.PortGen, tot.PortRulesHit, tot.PortRules, tot.AllHandHit, tot.AllHand, tot.AllGenHit, tot.AllGen, tot.AllRulesHit, tot.AllRules)
	return b.String()
}

func cmdGrammar(p Paths, args []string) error {
	fs := flag.NewFlagSet("grammar", flag.ExitOnError)
	prod := fs.String("prod", "", "list one production's elements")
	_ = fs.Parse(args)
	ix, err := LoadIndex(p)
	if err != nil {
		return err
	}
	if *prod != "" {
		k := 0
		for i := range ix.Elements {
			e := &ix.Elements[i]
			if e.Prod != *prod {
				continue
			}
			tag := ""
			if e.Kind != EKProduction && IsRule(e.Kind) {
				k++
				tag = fmt.Sprintf("#%d ", k)
			}
			fmt.Fprintf(os.Stdout, "%s%s  (%d cases)\n", tag, elemLabel(e), ix.Parts[e.Part].Reach)
		}
		return nil
	}
	fmt.Print(ix.grammarText())
	return nil
}

// grammarText is the totals by kind and every unmapped element.
func (ix *Index) grammarText() string {
	type c struct{ total, hit, portTotal, portHit int }
	by := map[string]*c{}
	var keys []string
	for i := range ix.Elements {
		e := &ix.Elements[i]
		x := by[e.Kind]
		if x == nil {
			x = &c{}
			by[e.Kind] = x
			keys = append(keys, e.Kind)
		}
		x.total++
		reached := ix.Parts[e.Part].Reach > 0
		if reached {
			x.hit++
		}
		if e.InPort {
			x.portTotal++
			if reached {
				x.portHit++
			}
		}
	}
	sort.Strings(keys)
	var b strings.Builder
	fmt.Fprintf(&b, "| element | mapped | reached by a case | in `port` | reached |\n|---|---:|---:|---:|---:|\n")
	for _, k := range keys {
		x := by[k]
		fmt.Fprintf(&b, "| %s | %d | %d | %d | %d |\n", k, x.total, x.hit, x.portTotal, x.portHit)
	}
	both := 0
	dec := 0
	type key struct{ prod, path string }
	ways := map[key]int{}
	for i := range ix.Elements {
		e := &ix.Elements[i]
		if (e.Kind == EKLookTrue || e.Kind == EKLookFalse) && ix.Parts[e.Part].Reach > 0 {
			ways[key{e.Prod, fmt.Sprint(e.Path, "@", e.Ref.Line)}]++
		}
		if e.Kind == EKLookTrue {
			dec++
		}
	}
	for _, n := range ways {
		if n == 2 {
			both++
		}
	}
	fmt.Fprintf(&b, "\nLOOKAHEAD decisions and semantic checks seen both ways, per JaCoCo: %d of %d. JaCoCo cannot record an arm whose first call throws, so an arm can have run and show as unreached.\n", both, dec)
	seenKinds := 0
	for i := range ix.Tokens {
		if ix.Tokens[i].Reach > 0 {
			seenKinds++
		}
	}
	fmt.Fprintf(&b, "Token kinds seen: %d of %d. The rest include private regular expressions and skipped kinds, which the parser never returns.\n", seenKinds, len(ix.Tokens))
	fmt.Fprintf(&b, "\nUnmapped: %d.\n", len(ix.Unmapped))
	for _, u := range ix.Unmapped {
		fmt.Fprintf(&b, "- %s, %s: %s (%s)\n", u.Prod, u.Kind, u.Text, u.Reason)
	}
	return b.String()
}
