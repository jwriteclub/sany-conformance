// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"reflect"
	"testing"
)

// world builds a universe directly: one upstream case, then one case per entry of
// sizes, each reaching the elements listed for it.
func world(t *testing.T, upstream []int32, sizes []int, reach [][]int32) (*universe, []int32) {
	t.Helper()
	ix := &Index{}
	ix.Cases = append(ix.Cases, Case{ID: "up", Group: "tlaplus-other", Size: 1})
	u := &universe{ix: ix, evidence: map[string]string{}}
	u.caseEl = [][]int32{upstream}
	u.upCase = []bool{true}
	n := int32(0)
	for _, l := range reach {
		for _, e := range l {
			if e >= n {
				n = e + 1
			}
		}
	}
	for _, e := range upstream {
		if e >= n {
			n = e + 1
		}
	}
	for i, s := range sizes {
		ix.Cases = append(ix.Cases, Case{ID: string(rune('a' + i)), Group: "mutants", Size: s})
		u.caseEl = append(u.caseEl, reach[i])
		u.upCase = append(u.upCase, false)
	}
	u.count = make([]int32, n)
	u.upstream = make([]bool, n)
	u.names = make([]string, n)
	for e := range u.names {
		u.names[e] = "e" + string(rune('A'+e))
	}
	for ci, l := range u.caseEl {
		for _, e := range l {
			u.count[e]++
			if u.upCase[ci] {
				u.upstream[e] = true
			}
		}
	}
	order := make([]int32, len(ix.Cases))
	for i := range order {
		order[i] = int32(i)
	}
	return u, order
}

func ids(u *universe, chosen map[int32]bool) map[string]bool {
	out := map[string]bool{}
	for ci := range chosen {
		out[u.ix.Cases[ci].ID] = true
	}
	return out
}

func TestSmallCasesBeatTheKitchenSink(t *testing.T) {
	// a reaches everything and is huge. b, c, d each reach one thing and are small.
	// Plain greedy set cover would take a alone; the rule takes the small ones.
	u, order := world(t, nil, []int{1000, 10, 11, 12}, [][]int32{{0, 1, 2}, {0}, {1}, {2}})
	got := ids(u, u.selectCases(order))
	want := map[string]bool{"b": true, "c": true, "d": true}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestUpstreamCoverageCountsAndRedundantCasesDrop(t *testing.T) {
	// Element 0 is upstream's, so nothing is kept for it. c is the smallest for 1 and
	// b for 2, and b also reaches 1, so c turns out redundant and is dropped largest
	// first only if nothing needs it.
	u, order := world(t, []int32{0}, []int{20, 30, 10}, [][]int32{{0}, {1, 2}, {1}})
	got := ids(u, u.selectCases(order))
	want := map[string]bool{"b": true}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestAnyOrderGivesTheSameSelection(t *testing.T) {
	u, order := world(t, nil, []int{5, 5, 7, 9}, [][]int32{{0, 1}, {1, 2}, {2, 3}, {0, 3}})
	first := ids(u, u.selectCases(order))
	rev := make([]int32, len(order))
	for i := range order {
		rev[i] = order[len(order)-1-i]
	}
	if got := ids(u, u.selectCases(rev)); !reflect.DeepEqual(got, first) {
		t.Fatalf("got %v reversed, %v forward", got, first)
	}
}
