// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// LexerArm is one case of a jjMoveNfa switch.
type LexerArm struct {
	LexState int
	Section  int // 0: curChar < 64, 1: curChar < 128, 2: above
	States   []int
	Line     int
	EndLine  int
}

// LexerStateReport is one lexical state's result.
type LexerStateReport struct {
	LexState  int
	Name      string
	Reachable bool
	Starts    int
	NfaStates int
	Reached   int
	Arms      int
	DeadArms  []LexerArm
	DfaTests  int
	DfaDead   []int // source lines
}

// LexerReport is the whole token manager.
type LexerReport struct {
	States []LexerStateReport
}

var (
	reNfaHead   = regexp.MustCompile(`^private final int jjMoveNfa_(\d+)\(int startState, int curPos\)`)
	reCase      = regexp.MustCompile(`^\s*case (\d+):\s*$`)
	reSwitch    = regexp.MustCompile(`switch\(jjstateSet\[--i\]\)`)
	reAdd1      = regexp.MustCompile(`jjCheckNAdd\((\d+)\)`)
	reAdd2      = regexp.MustCompile(`jjCheckNAddTwoStates\((\d+), (\d+)\)`)
	reAddRange  = regexp.MustCompile(`jjCheckNAddStates\((\d+), (\d+)\)|jjAddStates\((\d+), (\d+)\)`)
	reAddPair   = regexp.MustCompile(`jjCheckNAddStates\((\d+)\)`)
	reAddDirect = regexp.MustCompile(`jjstateSet\[jjnewStateCnt\+\+\] = (\d+)`)
	reStartNfa  = regexp.MustCompile(`jjMoveNfa_(\d+)\((\d+), `)
	reStartWith = regexp.MustCompile(`jjStartNfaWithStates_(\d+)\(\d+, \d+, (\d+)\)`)
	reStopHead  = regexp.MustCompile(`^private final int jjStopStringLiteralDfa_(\d+)\(`)
	reReturnNum = regexp.MustCompile(`return (\d+);`)
	reDfaHead   = regexp.MustCompile(`^private final int jjMoveStringLiteralDfa(\d+)_(\d+)\((.*)\)`)
	reDfaCall   = regexp.MustCompile(`jjMoveStringLiteralDfa(\d+)_(\d+)\(([^)]*)\)`)
	reDfaTest   = regexp.MustCompile(`\(active(\d+) & (0x[0-9a-fA-F]+)L\) != 0L`)
	reMethodTop = regexp.MustCompile(`^(private|public|protected|static)\b`)
)

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fatalf("lexer: %v", err)
	}
	return n
}

// AnalyseLexer reads the generated token manager's source. The state graph is
// exact for the NFA: an arm whose state no transition or start names cannot run,
// whatever the input. The DFA pass is a forward union of the literal bit masks.
func AnalyseLexer(path string) (*LexerReport, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	next, err := lexerIntTable(lines, "jjnextStates")
	if err != nil {
		return nil, err
	}
	names, newLex, err := lexerStates(lines)
	if err != nil {
		return nil, err
	}
	reach := map[int]bool{0: true}
	for _, v := range newLex {
		if v >= 0 {
			reach[v] = true
		}
	}
	for _, l := range lines {
		if i := strings.Index(l, "SwitchTo( "); i >= 0 {
			name := strings.TrimSpace(strings.TrimSuffix(strings.TrimSuffix(strings.TrimSpace(l[i+len("SwitchTo( "):]), ";"), ")"))
			for k, n := range names {
				if strings.HasPrefix(name, n) {
					reach[k] = true
				}
			}
		}
	}
	rep := &LexerReport{}
	for k, name := range names {
		st := LexerStateReport{LexState: k, Name: name, Reachable: reach[k]}
		starts := map[int]bool{}
		for i, l := range lines {
			if m := reStartNfa.FindStringSubmatch(l); m != nil && atoi(m[1]) == k {
				starts[atoi(m[2])] = true
			}
			if m := reStartWith.FindStringSubmatch(l); m != nil && atoi(m[1]) == k {
				starts[atoi(m[2])] = true
			}
			if m := reStopHead.FindStringSubmatch(l); m != nil && atoi(m[1]) == k {
				for j := i + 1; j < len(lines) && !reMethodTop.MatchString(lines[j]); j++ {
					if r := reReturnNum.FindStringSubmatch(lines[j]); r != nil {
						starts[atoi(r[1])] = true
					}
				}
			}
		}
		arms, succ := lexerArms(lines, k, next)
		seen := map[int]bool{}
		var work []int
		for s := range starts {
			work = append(work, s)
		}
		sort.Ints(work)
		st.Starts = len(work)
		for _, s := range work {
			seen[s] = true
		}
		for len(work) > 0 {
			s := work[len(work)-1]
			work = work[:len(work)-1]
			for _, t := range succ[s] {
				if !seen[t] {
					seen[t] = true
					work = append(work, t)
				}
			}
		}
		all := map[int]bool{}
		for _, a := range arms {
			for _, s := range a.States {
				all[s] = true
			}
		}
		st.NfaStates = len(all)
		for s := range all {
			if seen[s] {
				st.Reached++
			}
		}
		st.Arms = len(arms)
		for _, a := range arms {
			alive := false
			for _, s := range a.States {
				if seen[s] {
					alive = true
				}
			}
			if !alive {
				st.DeadArms = append(st.DeadArms, a)
			}
		}
		st.DfaTests, st.DfaDead = lexerDfa(lines, k)
		rep.States = append(rep.States, st)
	}
	return rep, nil
}

func lexerIntTable(lines []string, name string) ([]int, error) {
	for i, l := range lines {
		if strings.Contains(l, " "+name+" = {") {
			var out []int
			for j := i + 1; j < len(lines); j++ {
				t := strings.TrimSpace(lines[j])
				if strings.HasPrefix(t, "}") {
					return out, nil
				}
				for _, w := range strings.Split(t, ",") {
					w = strings.TrimSpace(w)
					if w != "" {
						out = append(out, atoi(w))
					}
				}
			}
		}
	}
	return nil, fmt.Errorf("no table %s", name)
}

func lexerStates(lines []string) ([]string, []int, error) {
	var names []string
	for i, l := range lines {
		if strings.Contains(l, "lexStateNames = {") {
			for j := i + 1; j < len(lines); j++ {
				t := strings.TrimSpace(lines[j])
				if strings.HasPrefix(t, "}") {
					break
				}
				names = append(names, strings.Trim(t, `", `))
			}
			break
		}
	}
	if len(names) == 0 {
		return nil, nil, fmt.Errorf("no lexStateNames")
	}
	table, err := lexerIntTable(lines, "jjnewLexState")
	return names, table, err
}

// lexerArms reads one jjMoveNfa method's arms and its state transition graph.
func lexerArms(lines []string, k int, next []int) ([]LexerArm, map[int][]int) {
	var arms []LexerArm
	succ := map[int][]int{}
	start := -1
	for i, l := range lines {
		if m := reNfaHead.FindStringSubmatch(l); m != nil && atoi(m[1]) == k {
			start = i
			break
		}
	}
	if start < 0 {
		return nil, succ
	}
	section := -1
	var cur *LexerArm
	flush := func(end int) {
		if cur != nil {
			cur.EndLine = end
			arms = append(arms, *cur)
			cur = nil
		}
	}
	inBody := false
	for i := start + 1; i < len(lines); i++ {
		l := lines[i]
		if reMethodTop.MatchString(l) {
			break
		}
		if reSwitch.MatchString(l) {
			section++
			continue
		}
		if m := reCase.FindStringSubmatch(l); m != nil {
			s := atoi(m[1])
			if cur == nil || inBody {
				flush(i)
				cur = &LexerArm{LexState: k, Section: section, Line: i + 1}
			}
			inBody = false
			cur.States = append(cur.States, s)
			continue
		}
		if strings.Contains(l, "default :") || strings.Contains(l, "while(i != startsAt)") {
			flush(i)
			inBody = false
			continue
		}
		if cur == nil {
			continue
		}
		inBody = true
		var targets []int
		for _, m := range reAdd1.FindAllStringSubmatch(l, -1) {
			targets = append(targets, atoi(m[1]))
		}
		for _, m := range reAdd2.FindAllStringSubmatch(l, -1) {
			targets = append(targets, atoi(m[1]), atoi(m[2]))
		}
		for _, m := range reAddRange.FindAllStringSubmatch(l, -1) {
			a, b := m[1], m[2]
			if a == "" {
				a, b = m[3], m[4]
			}
			for x := atoi(a); x <= atoi(b) && x < len(next); x++ {
				targets = append(targets, next[x])
			}
		}
		for _, m := range reAddPair.FindAllStringSubmatch(l, -1) {
			x := atoi(m[1])
			if x+1 < len(next) {
				targets = append(targets, next[x], next[x+1])
			}
		}
		for _, m := range reAddDirect.FindAllStringSubmatch(l, -1) {
			targets = append(targets, atoi(m[1]))
		}
		for _, s := range cur.States {
			succ[s] = append(succ[s], targets...)
		}
	}
	flush(len(lines))
	return arms, succ
}

// lexerDfa pushes the possible active bits through the literal DFA methods, in
// depth order, and counts the bit tests no literal can satisfy.
func lexerDfa(lines []string, k int) (int, []int) {
	type body struct {
		depth      int
		start, end int
	}
	var bodies []body
	for i, l := range lines {
		m := reDfaHead.FindStringSubmatch(l)
		if m == nil || atoi(m[2]) != k {
			continue
		}
		end := len(lines)
		for j := i + 1; j < len(lines); j++ {
			if reMethodTop.MatchString(lines[j]) {
				end = j
				break
			}
		}
		bodies = append(bodies, body{atoi(m[1]), i, end})
	}
	sort.Slice(bodies, func(a, b int) bool { return bodies[a].depth < bodies[b].depth })
	union := map[int]map[int]uint64{} // depth, word, bits
	get := func(d int) map[int]uint64 {
		if union[d] == nil {
			union[d] = map[int]uint64{}
		}
		return union[d]
	}
	tests := 0
	var dead []int
	for _, b := range bodies {
		here := get(b.depth)
		for i := b.start + 1; i < b.end; i++ {
			l := lines[i]
			for _, m := range reDfaTest.FindAllStringSubmatch(l, -1) {
				tests++
				mask, err := strconv.ParseUint(strings.TrimPrefix(m[2], "0x"), 16, 64)
				if err != nil {
					fatalf("lexer: %v", err)
				}
				if here[atoi(m[1])]&mask == 0 {
					dead = append(dead, i+1)
				}
			}
			for _, m := range reDfaCall.FindAllStringSubmatch(l, -1) {
				if atoi(m[2]) != k {
					continue
				}
				callee := get(atoi(m[1]))
				args := strings.Split(m[3], ",")
				for x := range args {
					args[x] = strings.TrimSpace(args[x])
				}
				if len(args) > 0 && strings.HasPrefix(args[0], "active") {
					for x := 0; x+1 < len(args); x += 2 {
						w := atoi(strings.TrimPrefix(args[x], "active"))
						callee[w] |= here[w] & parseJavaLong(args[x+1])
					}
				} else {
					for w, a := range args {
						if a != "" {
							callee[w] |= parseJavaLong(a)
						}
					}
				}
			}
		}
	}
	return tests, dead
}

func parseJavaLong(s string) uint64 {
	s = strings.TrimSuffix(s, "L")
	if strings.HasPrefix(s, "0x") {
		v, err := strconv.ParseUint(s[2:], 16, 64)
		if err != nil {
			fatalf("lexer: %v", err)
		}
		return v
	}
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		fatalf("lexer: %v", err)
	}
	return v
}
