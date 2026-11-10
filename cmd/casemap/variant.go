// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// variantParts is what an exec file covers, by index part.
func (ix *Index) variantParts(data ExecData) (map[int32]bool, error) {
	byID := map[uint64]int{}
	for i := range ix.Classes {
		byID[ix.Classes[i].ID] = i
	}
	out := map[int32]bool{}
	for id, c := range data {
		ci, ok := byID[id]
		if !ok {
			continue
		}
		cl := &ix.Classes[ci]
		if int(cl.NProbes) != len(c.Probes) {
			return nil, fmt.Errorf("%s: %d probes against %d in the index", c.Name, len(c.Probes), cl.NProbes)
		}
		for p, on := range c.Probes {
			if on {
				for _, part := range cl.Images[p] {
					out[part] = true
				}
			}
		}
	}
	return out, nil
}

func parseKinds(s string) []uint16 {
	out := []uint16{}
	for _, w := range strings.Split(s, ",") {
		if w != "" {
			n, _ := strconv.Atoi(w)
			out = append(out, uint16(n))
		}
	}
	return out
}

// cmdVariantCheck says whether a changed copy of a case still reaches what the
// original is here for. Everything the case reaches beyond upstream's sets must still
// be reached; the rest is reported. Used to clean a case's bytes without losing its why.
//
//	casemap variant-check -case OLD -exec FILE -tokens 1,2,3
func cmdVariantCheck(p Paths, args []string) error {
	fs := flag.NewFlagSet("variant-check", flag.ExitOnError)
	id := fs.String("case", "", "the original case")
	execPath := fs.String("exec", "", "JaCoCo exec file of the variant")
	tokens := fs.String("tokens", "", "token kinds of the variant, comma separated")
	_ = fs.Parse(args)
	ix, err := LoadIndex(p)
	if err != nil {
		return err
	}
	u, err := loadUniverse(p, ix)
	if err != nil {
		return err
	}
	ci, ok := u.byID[*id]
	if !ok {
		return fmt.Errorf("no case %s", *id)
	}
	data, err := ReadExec(*execPath)
	if err != nil {
		return err
	}
	have, err := ix.variantParts(data)
	if err != nil {
		return err
	}
	kinds := map[uint16]bool{}
	for _, k := range parseKinds(*tokens) {
		kinds[k] = true
	}
	var lostWhy, lostOther []string
	inOld := map[int32]bool{}
	for _, e := range u.caseEl[ci] {
		inOld[e] = true
		reached := true
		switch {
		case int(e) < u.nParts:
			reached = have[e]
		case int(e) < u.nParts+u.nTokens:
			reached = kinds[uint16(int(e)-u.nParts)]
		default:
			continue // extra elements are checked from the stock and tlx runs
		}
		if reached {
			continue
		}
		if u.upstream[e] {
			lostOther = append(lostOther, u.names[e])
		} else {
			lostWhy = append(lostWhy, u.names[e])
		}
	}
	gained := 0
	for e := range have {
		if !inOld[e] {
			gained++
		}
	}
	sort.Strings(lostWhy)
	sort.Strings(lostOther)
	return json.NewEncoder(os.Stdout).Encode(map[string]any{"case": *id, "parts_reached": len(have), "lost_why": lostWhy, "lost_upstream_covered": len(lostOther), "lost_upstream_sample": first(lostOther, 5), "gained": gained})
}

func first(l []string, n int) []string {
	if len(l) < n {
		return l
	}
	return l[:n]
}

// cmdVariantApply swaps a case in the index for its changed copy: new id, content
// hash, size, probes and token kinds.
//
//	casemap variant-apply -old ID -new ID -hash H -size N -exec FILE -tokens 1,2,3
func cmdVariantApply(p Paths, args []string) error {
	fs := flag.NewFlagSet("variant-apply", flag.ExitOnError)
	oldID := fs.String("old", "", "the original case")
	newID := fs.String("new", "", "the changed case")
	hash := fs.String("hash", "", "its content hash")
	size := fs.Int("size", 0, "its size in bytes")
	execPath := fs.String("exec", "", "JaCoCo exec file")
	tokens := fs.String("tokens", "", "token kinds, comma separated")
	_ = fs.Parse(args)
	ix, err := LoadIndex(p)
	if err != nil {
		return err
	}
	ci := ix.findCase(*oldID)
	if ci < 0 {
		return fmt.Errorf("no case %s", *oldID)
	}
	data, err := ReadExec(*execPath)
	if err != nil {
		return err
	}
	for _, pn := range ix.CasePats[ci] {
		pat := &ix.Patterns[pn]
		kept := pat.Cases[:0]
		for _, c := range pat.Cases {
			if int(c) != ci {
				kept = append(kept, c)
			}
		}
		pat.Cases = kept
	}
	patIndex := map[string]int32{}
	for pn := range ix.Patterns {
		patIndex[strconv.Itoa(int(ix.Patterns[pn].Class))+" "+string(ix.Patterns[pn].Bits)] = int32(pn)
	}
	byID := map[uint64]int32{}
	for i := range ix.Classes {
		byID[ix.Classes[i].ID] = int32(i)
	}
	ids := make([]uint64, 0, len(data))
	for id := range data {
		ids = append(ids, id)
	}
	sort.Slice(ids, func(a, b int) bool { return ids[a] < ids[b] })
	var pats []int32
	for _, id := range ids {
		cls, ok := byID[id]
		if !ok {
			continue
		}
		bits := make([]byte, (len(data[id].Probes)+7)/8)
		any := false
		for i, v := range data[id].Probes {
			if v {
				bits[i/8] |= 1 << (i % 8)
				any = true
			}
		}
		if !any {
			continue
		}
		k := strconv.Itoa(int(cls)) + " " + string(bits)
		pn, ok := patIndex[k]
		if !ok {
			pn = int32(len(ix.Patterns))
			patIndex[k] = pn
			ix.Patterns = append(ix.Patterns, Pattern{Class: cls, Bits: bits})
			ix.Classes[cls].Patterns = append(ix.Classes[cls].Patterns, pn)
		}
		ix.Patterns[pn].Cases = append(ix.Patterns[pn].Cases, int32(ci))
		sort.Slice(ix.Patterns[pn].Cases, func(a, b int) bool { return ix.Patterns[pn].Cases[a] < ix.Patterns[pn].Cases[b] })
		pats = append(pats, pn)
	}
	ix.CasePats[ci] = pats
	ix.CaseTokens[ci] = parseKinds(*tokens)
	c := &ix.Cases[ci]
	c.ID, c.Hash, c.Size = *newID, *hash, *size
	if err := ix.recount(); err != nil {
		return err
	}
	return ix.Save(p.IndexFile())
}
