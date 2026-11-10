// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import "sort"

// The one intra-method refinement: a branch on a field whose value is fixed in
// the closed world, or on a static method that returns a constant. Sound as long
// as no reachable code writes the field by reflection, which the reflection
// listing in the report covers.

// constValue is the fixed value of a field or the result of a method.
type constValue struct {
	Null bool // a reference known to be null
	N    int64
}

// FieldWrite is one bytecode write of a field.
type FieldWrite struct {
	Method *Method
	Off    int
	Const  bool
	Value  constValue
}

type constantFields struct {
	values map[string]constValue
	why    map[string]string
	writes map[string][]FieldWrite
}

// pushedConstant reads a constant-pushing instruction.
func pushedConstant(c *Class, in Insn, code []byte) (constValue, bool) {
	switch {
	case in.Op == opAconstNull:
		return constValue{Null: true}, true
	case in.Op >= opIconstM1 && in.Op <= opIconst5:
		return constValue{N: int64(in.Op) - 3}, true
	case in.Op == opBipush:
		return constValue{N: int64(int8(code[in.Off+1]))}, true
	case in.Op == opSipush:
		return constValue{N: int64(int16(uint16(code[in.Off+1])<<8 | uint16(code[in.Off+2])))}, true
	case in.Op == opLdc || in.Op == opLdcW:
		if e := c.entry(in.Index); e.tag == tagInteger {
			return constValue{N: e.n}, true
		}
	}
	return constValue{}, false
}

// Constants scans every write of every field in the world, once.
func (w *World) Constants() *constantFields {
	if w.constants != nil {
		return w.constants
	}
	cf := &constantFields{values: map[string]constValue{}, why: map[string]string{}, writes: map[string][]FieldWrite{}}
	for _, m := range w.Methods {
		insns, err := m.Decode()
		if err != nil {
			fatalf("%v", err)
		}
		for i, in := range insns {
			if in.Op != opPutfield && in.Op != opPutstatic {
				continue
			}
			ref := m.Class.memberRef(in.Index)
			f := w.ResolveField(ref.Class, ref.Name+":"+ref.Desc)
			if f == nil || !f.Class.App {
				continue
			}
			fw := FieldWrite{Method: m, Off: in.Off}
			if i > 0 {
				if v, ok := pushedConstant(m.Class, insns[i-1], m.Code); ok {
					// The push must really feed this write: no jump may land between them.
					if !isJumpTarget(m, in.Off) {
						fw.Const, fw.Value = true, v
					}
				}
			}
			cf.writes[f.Key()] = append(cf.writes[f.Key()], fw)
		}
	}
	for _, n := range w.AppNames {
		for _, f := range w.Classes[n].Fields {
			w.judgeField(cf, f)
		}
	}
	w.constants = cf
	return cf
}

func isJumpTarget(m *Method, off int) bool {
	insns, _ := m.Decode()
	for _, in := range insns {
		for _, t := range in.Targets {
			if t == off {
				return true
			}
		}
	}
	for _, h := range m.Handlers {
		if h.Target == off {
			return true
		}
	}
	return false
}

func (w *World) judgeField(cf *constantFields, f *Field) {
	if f.HasConst {
		return // javac inlines these; a branch on one was already folded at compile time
	}
	var def constValue
	switch f.Desc[0] {
	case 'Z', 'B', 'S', 'C', 'I':
	case 'L', '[':
		def = constValue{Null: true}
	default:
		return
	}
	writes := cf.writes[f.Key()]
	allDefault := true
	for _, fw := range writes {
		if !fw.Const || fw.Value != def {
			allDefault = false
			break
		}
	}
	if allDefault {
		cf.values[f.Key()] = def
		if len(writes) == 0 {
			cf.why[f.Key()] = "never written"
		} else {
			cf.why[f.Key()] = "only ever written with its default"
		}
		return
	}
	// A static field written once, with a constant, in its own initialiser, before
	// the initialiser can run any other code: no reader can see the default.
	if f.Access&accStatic == 0 || len(writes) != 1 {
		return
	}
	fw := writes[0]
	if !fw.Const || fw.Method.Class != f.Class || fw.Method.Name != "<clinit>" {
		return
	}
	events, err := fw.Method.Events()
	if err != nil {
		fatalf("%v", err)
	}
	for _, ev := range events {
		if ev.Off >= fw.Off {
			break
		}
		if ev.Kind != EvStatic || ev.Ref.Class != f.Class.Name {
			return
		}
	}
	cf.values[f.Key()] = fw.Value
	cf.why[f.Key()] = "written once, with a constant, in its own initialiser"
}

// constantReturn says what a method returns when its whole body is a constant push and a return.
func constantReturn(m *Method) (constValue, bool) {
	if m == nil || !m.Class.App {
		return constValue{}, false
	}
	insns, err := m.Decode()
	if err != nil || len(insns) != 2 {
		return constValue{}, false
	}
	if insns[1].Op != opIreturn && insns[1].Op != 176 {
		return constValue{}, false
	}
	return pushedConstant(m.Class, insns[0], m.Code)
}

// LineOutcome says how many branch outcomes a refinement or cut removes on a line.
type LineOutcome struct {
	Line     int
	Outcomes int
	// Fact is what licensed the removal: a field or a method, or empty for a dead block.
	Fact string
}

type deadInfo struct {
	offsets map[int]bool
	removed []LineOutcome
}

// DeadOffsets lists the instructions the refinement rules out in a method.
func (w *World) DeadOffsets(m *Method) map[int]bool { return w.dead(m).offsets }

var noDead = &deadInfo{}

func (w *World) dead(m *Method) *deadInfo {
	if w.deadCache == nil {
		w.deadCache = map[*Method]*deadInfo{}
	}
	if d, ok := w.deadCache[m]; ok {
		return d
	}
	d := w.computeDead(m)
	w.deadCache[m] = d
	return d
}

// foldedBranch says that a conditional has one possible next offset.
type foldedBranch struct {
	next int
	fact string
}

func (w *World) computeDead(m *Method) *deadInfo {
	insns, err := m.Decode()
	if err != nil {
		fatalf("%v", err)
	}
	if len(insns) == 0 {
		return noDead
	}
	targets := map[int]bool{}
	for _, in := range insns {
		for _, t := range in.Targets {
			targets[t] = true
		}
	}
	for _, h := range m.Handlers {
		targets[h.Target] = true
	}
	folds := map[int]foldedBranch{}
	for i, in := range insns {
		v, fact, ok := w.constantLoad(m, in)
		if !ok || i+1 >= len(insns) {
			continue
		}
		next := insns[i+1]
		if targets[next.Off] {
			continue
		}
		if next.IsConditional() {
			if taken, ok := decideUnary(next.Op, v); ok {
				folds[next.Off] = foldedBranch{next: chooseNext(next, taken), fact: fact}
			}
			continue
		}
		if i+2 >= len(insns) || targets[insns[i+2].Off] {
			continue
		}
		rhs, ok := pushedConstant(m.Class, next, m.Code)
		if !ok {
			continue
		}
		cmp := insns[i+2]
		if taken, ok := decideBinary(cmp.Op, v, rhs); ok {
			folds[cmp.Off] = foldedBranch{next: chooseNext(cmp, taken), fact: fact}
		}
	}
	if len(folds) == 0 {
		return noDead
	}
	index := map[int]int{}
	for i, in := range insns {
		index[in.Off] = i
	}
	alive := make([]bool, len(insns))
	var work []int
	mark := func(off int) {
		if i, ok := index[off]; ok && !alive[i] {
			alive[i] = true
			work = append(work, i)
		}
	}
	mark(0)
	for {
		for len(work) > 0 {
			i := work[len(work)-1]
			work = work[:len(work)-1]
			in := insns[i]
			if f, ok := folds[in.Off]; ok {
				mark(f.next)
				continue
			}
			for _, t := range in.Targets {
				mark(t)
			}
			if fallsThrough(in.Op) && i+1 < len(insns) {
				mark(insns[i+1].Off)
			}
		}
		// An exception handler is alive when anything it covers is.
		grew := false
		for _, h := range m.Handlers {
			if alive[index[h.Target]] {
				continue
			}
			for i, in := range insns {
				if alive[i] && in.Off >= h.Start && in.Off < h.End {
					mark(h.Target)
					grew = true
					break
				}
			}
		}
		if !grew {
			break
		}
	}
	d := &deadInfo{offsets: map[int]bool{}}
	for i, in := range insns {
		line := m.LineAt(in.Off)
		if !alive[i] {
			d.offsets[in.Off] = true
			if n := conditionalOutcomes(in); n > 0 {
				d.removed = append(d.removed, LineOutcome{Line: line, Outcomes: n})
			}
			continue
		}
		if f, ok := folds[in.Off]; ok {
			d.removed = append(d.removed, LineOutcome{Line: line, Outcomes: 1, Fact: f.fact})
		}
	}
	return d
}

// constantLoad recognises an instruction that yields a value known in advance.
func (w *World) constantLoad(m *Method, in Insn) (constValue, string, bool) {
	switch in.Op {
	case opGetstatic, opGetfield:
		ref := m.Class.memberRef(in.Index)
		f := w.ResolveField(ref.Class, ref.Name+":"+ref.Desc)
		if f == nil || !f.Class.App {
			return constValue{}, "", false
		}
		v, ok := w.Constants().values[f.Key()]
		return v, f.Key(), ok
	case opInvokestat:
		ref := m.Class.memberRef(in.Index)
		t := w.ResolveMethod(ref.Class, ref.Name+":"+ref.Desc)
		v, ok := constantReturn(t)
		if !ok {
			return constValue{}, "", false
		}
		return v, t.Key(), true
	}
	return constValue{}, "", false
}

func chooseNext(in Insn, taken bool) int {
	if taken {
		return in.Targets[0]
	}
	return in.Off + in.Len
}

func fallsThrough(op uint8) bool {
	switch op {
	case opGoto, opGotoW, opAthrow, opRet, opTableswi, opLookupswi, opJsr, opJsrW:
		return false
	}
	return op < opIreturn || op > opReturn
}

func decideUnary(op uint8, v constValue) (bool, bool) {
	switch op {
	case opIfnull:
		return v.Null, v.Null
	case opIfnonnull:
		return false, v.Null
	}
	if v.Null {
		return false, false
	}
	switch op {
	case opIfeq:
		return v.N == 0, true
	case opIfne:
		return v.N != 0, true
	case opIflt:
		return v.N < 0, true
	case opIfge:
		return v.N >= 0, true
	case opIfgt:
		return v.N > 0, true
	case opIfle:
		return v.N <= 0, true
	}
	return false, false
}

func decideBinary(op uint8, a, b constValue) (bool, bool) {
	if a.Null || b.Null {
		return false, false
	}
	switch op {
	case opIfIcmpeq:
		return a.N == b.N, true
	case opIfIcmpne:
		return a.N != b.N, true
	case opIfIcmplt:
		return a.N < b.N, true
	case opIfIcmpge:
		return a.N >= b.N, true
	case opIfIcmpgt:
		return a.N > b.N, true
	case opIfIcmple:
		return a.N <= b.N, true
	}
	return false, false
}

// ConstantFacts lists the fields and methods that actually folded a branch in
// the given methods, with outcomes removed, sorted.
func (w *World) ConstantFacts(ms []*Method) []FactCount {
	counts := map[string]int{}
	for _, m := range ms {
		for _, lo := range w.dead(m).removed {
			counts[lo.Fact] += lo.Outcomes
		}
	}
	out := make([]FactCount, 0, len(counts))
	for k, v := range counts {
		out = append(out, FactCount{k, v})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Fact < out[j].Fact })
	return out
}

// FactCount pairs a constant with the outcomes it removes.
type FactCount struct {
	Fact     string
	Outcomes int
}

// CutRemoved counts the outcomes a set of cuts removes in a method: branches on
// the cut lines, and the guarding outcome of a branch that jumps into them.
func CutRemoved(m *Method, cuts []Cut) []LineOutcome {
	inCut := func(line int) bool {
		for i := range cuts {
			if cuts[i].Class == m.Class.Name && line >= cuts[i].From && line <= cuts[i].To {
				return true
			}
		}
		return false
	}
	any := false
	for i := range cuts {
		if cuts[i].Class == m.Class.Name {
			any = true
			break
		}
	}
	if !any {
		return nil
	}
	insns, err := m.Decode()
	if err != nil {
		fatalf("%v", err)
	}
	var out []LineOutcome
	for _, in := range insns {
		n := conditionalOutcomes(in)
		if n == 0 {
			continue
		}
		line := m.LineAt(in.Off)
		if inCut(line) {
			out = append(out, LineOutcome{Line: line, Outcomes: n})
			continue
		}
		if !in.IsConditional() {
			continue
		}
		if inCut(m.LineAt(in.Targets[0])) || inCut(m.LineAt(in.Off+in.Len)) {
			out = append(out, LineOutcome{Line: line, Outcomes: 1})
		}
	}
	return out
}
