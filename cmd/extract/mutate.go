// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"regexp"
	"sort"
	"strings"
)

// Tier C rules. Each takes the root file text and an RNG and returns the
// mutated text plus a one-line description, or ok=false when it has no site.
//
//	delete-token       remove one identifier or operator token from the module body
//	swap-delimiter     replace one closing ) ] } by a different closer
//	misalign-junction  shift one /\ or \/ bullet line 1-2 columns left or right
//	unknown-name       rename one use of a defined operator to an undefined name
//	prime-in-invariant add a prime to a VARIABLE inside an *Inv*/TypeOK definition
//	bad-arity          drop or add one argument at a call of a defined operator
var mutRules = []string{"delete-token", "swap-delimiter", "misalign-junction", "unknown-name", "prime-in-invariant", "bad-arity"}

type tok struct {
	kind       byte // i identifier, n number, s symbol, d delimiter
	start, end int
}

var reTok = regexp.MustCompile(`[A-Za-z_][A-Za-z0-9_]*|[0-9]+(?:\.[0-9]+)?|[()\[\]{}<>,]|[/\\=~#|+\-*^.:!@$%&?'<>]+`)

// body returns the byte range between the first MODULE header and the last
// ==== line, and the code with comments and strings blanked.
func body(src string) (lo, hi int, code string) {
	code = stripComments(src)
	m := reModule.FindStringIndex(code)
	if m == nil {
		return 0, 0, code
	}
	lo = m[1]
	hi = len(code)
	if i := strings.LastIndex(code, "\n===="); i > lo {
		hi = i
	}
	return
}

func tokens(code string, lo, hi int) []tok {
	var ts []tok
	for _, m := range reTok.FindAllStringIndex(code[lo:hi], -1) {
		s, e := m[0]+lo, m[1]+lo
		c := code[s]
		k := byte('s')
		switch {
		case c == '_' || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z'):
			k = 'i'
		case c >= '0' && c <= '9':
			k = 'n'
		case strings.ContainsRune("()[]{},", rune(c)) && e-s == 1:
			k = 'd'
		}
		ts = append(ts, tok{k, s, e})
	}
	return ts
}

var reDefHead = regexp.MustCompile(`(?m)^\s*(?:LOCAL\s+)?([A-Za-z_][A-Za-z0-9_]*)\s*(\(([^()]*)\))?\s*==`)
var reVars = regexp.MustCompile(`\bVARIABLES?\s+([A-Za-z0-9_]+(?:\s*,\s*[A-Za-z0-9_]+)*)`)

func mutate(rule, src string, r *rand.Rand) (string, string, bool) {
	lo, hi, code := body(src)
	if hi <= lo {
		return "", "", false
	}
	ts := tokens(code, lo, hi)
	pick := func(idx []int) int { return idx[r.Intn(len(idx))] }
	switch rule {
	case "delete-token":
		var idx []int
		for i, t := range ts {
			txt := code[t.start:t.end]
			if (t.kind == 'i' || t.kind == 's') && txt != "MODULE" && !strings.HasPrefix(txt, "==") && !strings.HasPrefix(txt, "--") {
				idx = append(idx, i)
			}
		}
		if len(idx) == 0 {
			return "", "", false
		}
		t := ts[pick(idx)]
		return src[:t.start] + src[t.end:], fmt.Sprintf("deleted %q at byte %d", code[t.start:t.end], t.start), true
	case "swap-delimiter":
		var idx []int
		for i, t := range ts {
			if t.kind == 'd' && strings.Contains(")]}", code[t.start:t.end]) {
				idx = append(idx, i)
			}
		}
		if len(idx) == 0 {
			return "", "", false
		}
		t := ts[pick(idx)]
		old := code[t.start : t.start+1]
		alts := strings.Split(strings.Replace(")]}", old, "", 1), "")
		nw := alts[r.Intn(len(alts))]
		return src[:t.start] + nw + src[t.end:], fmt.Sprintf("replaced %s by %s at byte %d", old, nw, t.start), true
	case "misalign-junction":
		lines := strings.SplitAfter(src, "\n")
		var idx []int
		off := 0
		starts := make([]int, len(lines))
		for i, l := range lines {
			starts[i] = off
			off += len(l)
			trim := strings.TrimLeft(l, " ")
			if starts[i] > lo && starts[i] < hi && (strings.HasPrefix(trim, "/\\") || strings.HasPrefix(trim, "\\/")) &&
				code[starts[i]+len(l)-len(trim)] != ' ' {
				idx = append(idx, i)
			}
		}
		if len(idx) == 0 {
			return "", "", false
		}
		i := pick(idx)
		d := 1 + r.Intn(2)
		indent := len(lines[i]) - len(strings.TrimLeft(lines[i], " "))
		var nl string
		if r.Intn(2) == 0 && indent >= d {
			nl = lines[i][d:]
			d = -d
		} else {
			nl = strings.Repeat(" ", d) + lines[i]
		}
		lines[i] = nl
		return strings.Join(lines, ""), fmt.Sprintf("shifted junction line %d by %+d columns", i+1, d), true
	}
	// rules needing definitions
	defs := map[string]int{} // name -> arity
	defSite := map[int]bool{}
	for _, m := range reDefHead.FindAllStringSubmatchIndex(code[lo:hi], -1) {
		name := code[lo+m[2] : lo+m[3]]
		ar := 0
		if m[4] >= 0 {
			if strings.TrimSpace(code[lo+m[6]:lo+m[7]]) != "" {
				ar = strings.Count(code[lo+m[6]:lo+m[7]], ",") + 1
			}
		}
		defs[name] = ar
		defSite[lo+m[2]] = true
	}
	switch rule {
	case "unknown-name":
		var idx []int
		for i, t := range ts {
			if t.kind == 'i' && !defSite[t.start] {
				if _, ok := defs[code[t.start:t.end]]; ok {
					idx = append(idx, i)
				}
			}
		}
		if len(idx) == 0 {
			return "", "", false
		}
		t := ts[pick(idx)]
		old := code[t.start:t.end]
		return src[:t.start] + old + "Undefd" + src[t.end:], fmt.Sprintf("renamed use of %s to %sUndefd at byte %d", old, old, t.start), true
	case "prime-in-invariant":
		vars := map[string]bool{}
		for _, m := range reVars.FindAllStringSubmatch(code, -1) {
			for _, v := range strings.Split(m[1], ",") {
				vars[strings.TrimSpace(v)] = true
			}
		}
		if len(vars) == 0 {
			return "", "", false
		}
		reInv := regexp.MustCompile(`(?i)inv|typeok`)
		heads := reDefHead.FindAllStringSubmatchIndex(code[lo:hi], -1)
		var idx []int
		for hi2, m := range heads {
			if !reInv.MatchString(code[lo+m[2] : lo+m[3]]) {
				continue
			}
			end := hi
			if hi2+1 < len(heads) {
				end = lo + heads[hi2+1][0]
			}
			for i, t := range ts {
				if t.kind == 'i' && t.start > lo+m[1] && t.end <= end && vars[code[t.start:t.end]] {
					idx = append(idx, i)
				}
			}
		}
		if len(idx) == 0 {
			return "", "", false
		}
		t := ts[pick(idx)]
		return src[:t.end] + "'" + src[t.end:], fmt.Sprintf("primed %s at byte %d inside an invariant definition", code[t.start:t.end], t.start), true
	case "bad-arity":
		type site struct{ open, close, nargs int }
		var sites []site
		for i, t := range ts {
			if t.kind != 'i' || defSite[t.start] || i+1 >= len(ts) || ts[i+1].kind != 'd' || code[ts[i+1].start] != '(' || ts[i+1].start != t.end {
				continue
			}
			ar, ok := defs[code[t.start:t.end]]
			if !ok || ar == 0 {
				continue
			}
			depth, commas := 0, 0
			cl := -1
			for j := i + 1; j < len(ts); j++ {
				if ts[j].kind == 'd' {
					switch code[ts[j].start] {
					case '(', '[', '{':
						depth++
					case ')', ']', '}':
						depth--
					case ',':
						if depth == 1 {
							commas++
						}
					}
				}
				if depth == 0 {
					cl = j
					break
				}
			}
			if cl > 0 {
				sites = append(sites, site{ts[i+1].start, ts[cl].start, commas + 1})
			}
		}
		if len(sites) == 0 {
			return "", "", false
		}
		s := sites[r.Intn(len(sites))]
		if s.nargs >= 2 {
			// drop the last top-level argument
			depth, last := 0, -1
			for j := s.open; j < s.close; j++ {
				switch code[j] {
				case '(', '[', '{':
					depth++
				case ')', ']', '}':
					depth--
				case ',':
					if depth == 1 {
						last = j
					}
				}
			}
			return src[:last] + src[s.close:], fmt.Sprintf("dropped last argument of call at byte %d (had %d)", s.open, s.nargs), true
		}
		return src[:s.close] + ", 0" + src[s.close:], fmt.Sprintf("added an extra argument to call at byte %d (had 1)", s.open), true
	}
	return "", "", false
}

func hash64(parts ...string) int64 {
	h := fnv.New64a()
	for _, p := range parts {
		h.Write([]byte(p))
		h.Write([]byte{0})
	}
	return int64(h.Sum64() >> 1)
}

const maxParentBytes = 30 << 10
const maxPerParent = 2

// mutants derives tier C cases. Parents are tier A cases expecting accept (no
// known deviations) and tier B real specs under 30 KiB, never duplicates, never
// files without a module header or with a module/file name mismatch, and
// never PlusCal files. Selection is round-robin across rules in a seeded
// order, at most two mutants per parent, tier A parents first (up to a third
// of the cap) then tier B.
func (x *ctx) mutants(all []*Case, canon map[string]*Case, seed int64, limit int) []*Case {
	type cand struct {
		parent *Case
		rule   string
		c      *Case
		key    int64
	}
	var pools = map[string][]cand{}
	for _, p := range all {
		if p.DupOf != "" || len(p.content[p.RootFile]) > maxParentBytes {
			continue
		}
		if len(p.Flags) > 0 && (has(p.Flags, "no_module_header") || has(p.Flags, "module_name_mismatch") || has(p.Flags, "has_pluscal") || has(p.Flags, "nonstandard")) {
			continue
		}
		if p.Tier == "A" && (p.Expect == nil || p.Expect.Result != "accept" || p.Expect.Upstream != "ok") {
			continue
		}
		if p.Tier == "B" && p.Expect != nil {
			continue // other parsers' corpus snippets duplicate tier A ideas
		}
		src := string(p.content[p.RootFile])
		for _, rule := range mutRules {
			r := rand.New(rand.NewSource(seed ^ hash64(p.ID, rule)))
			out, detail, ok := mutate(rule, src, r)
			if !ok || out == src {
				continue
			}
			c := &Case{Tier: "C", Source: "mutant/" + p.Tier, Commit: p.Commit, Origin: p.Origin, Name: p.Name,
				Root: p.Root, RootFile: p.RootFile, ExtDeps: p.ExtDeps, Unresolved: p.Unresolved, Flags: append([]string{}, p.Flags...),
				content: map[string][]byte{}}
			for k, v := range p.content {
				c.content[k] = v
			}
			c.content[p.RootFile] = []byte(out)
			c.Mut = &Mutant{Parent: p.ID, Rule: rule, Seed: seed, Detail: detail,
				Intent: map[string]string{"delete-token": "unknown", "swap-delimiter": "likely-reject", "misalign-junction": "unknown",
					"unknown-name": "likely-reject", "prime-in-invariant": "unknown", "bad-arity": "likely-reject"}[rule]}
			c.ID = "C-" + strings.TrimPrefix(p.ID, p.Tier+"-") + "-" + rule
			c.finalize()
			if _, dup := canon[c.Hash]; dup {
				continue
			}
			pools[p.Tier] = append(pools[p.Tier], cand{p, rule, c, hash64(fmt.Sprint(seed), p.ID, rule, "order")})
		}
	}
	var picked []*Case
	seen := map[string]bool{}
	perParent := map[string]int{}
	take := func(tier string, budget int) {
		byRule := map[string][]cand{}
		for _, cd := range pools[tier] {
			byRule[cd.rule] = append(byRule[cd.rule], cd)
		}
		for _, l := range byRule {
			sort.Slice(l, func(i, j int) bool { return l[i].key < l[j].key })
		}
		n := 0
		for progress := true; progress && n < budget; {
			progress = false
			for _, rule := range mutRules {
				l := byRule[rule]
				for len(l) > 0 {
					cd := l[0]
					l = l[1:]
					if perParent[cd.parent.ID] >= maxPerParent || seen[cd.c.Hash] {
						continue
					}
					perParent[cd.parent.ID]++
					seen[cd.c.Hash] = true
					picked = append(picked, cd.c)
					n++
					progress = true
					break
				}
				byRule[rule] = l
				if n >= budget {
					break
				}
			}
		}
	}
	take("A", limit/3)
	take("B", limit-len(picked))
	return picked
}

func has(l []string, s string) bool {
	for _, x := range l {
		if x == s {
			return true
		}
	}
	return false
}
