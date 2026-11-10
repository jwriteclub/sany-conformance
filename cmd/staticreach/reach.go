// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"sort"
)

// Precision picks how virtual calls are resolved.
type Precision int

const (
	// CHA dispatches a virtual call to every concrete application subtype.
	CHA Precision = iota
	// RTA dispatches only to classes instantiated in reachable code.
	RTA
	// RTAConst is RTA on top of the intra-method constant field refinement.
	RTAConst
)

func (p Precision) String() string {
	switch p {
	case CHA:
		return "CHA"
	case RTA:
		return "RTA"
	}
	return "RTA+const"
}

// Options describes one reachability question.
type Options struct {
	Entries   []*Method
	RootClass []*Class // instantiated from outside, every method a root
	Cuts      []Cut
	// CutMethods are methods nothing may reach, by method ID.
	CutMethods map[int]bool
	Precision  Precision
	// Reflect maps a method to the class it initialises by reflection.
	Reflect map[string]string
	// NoCallbacks drops the JDK callback assumption. Unsound; only for measuring it.
	NoCallbacks bool
}

// Cut removes the sites on a range of source lines of one class.
type Cut struct {
	Name     string
	Class    string
	From, To int
	Why      string
}

// EdgeKind says why one method can make another run.
type EdgeKind uint8

const (
	EdgeCall EdgeKind = iota
	EdgeClinit
	EdgeCallback
	EdgeLambda
)

// Result is the reachable method set with its call graph.
type Result struct {
	W       *World
	Opts    Options
	Reached []bool
	Inst    map[string]bool
	edges   map[int][]int
	edgeSet map[uint64]bool
	kinds   map[uint64]EdgeKind
	roots   []int
	parent  []int32
	depth   []int32
	queue   []*Method
	vcalls  map[string]*vcall
	byClass map[string][]*vcall
}

type vcall struct {
	class, sig string
	callers    []int
	targets    []*Method
	seen       map[int]bool
	dispatched map[string]bool
	extra      *Method
}

// Analyse computes the fixed point. The reached set does not depend on the order
// of the work list; the parent links come from a breadth first pass afterwards.
func Analyse(w *World, o Options) *Result {
	r := &Result{
		W: w, Opts: o,
		Reached: make([]bool, len(w.Methods)),
		Inst:    map[string]bool{},
		edges:   map[int][]int{},
		edgeSet: map[uint64]bool{},
		kinds:   map[uint64]EdgeKind{},
		vcalls:  map[string]*vcall{},
		byClass: map[string][]*vcall{},
	}
	for _, m := range o.Entries {
		r.root(m)
		// A constructor as an entry means the caller made the object.
		if m.Name == "<init>" {
			r.instantiate(-1, m.Class)
		}
	}
	for _, c := range o.RootClass {
		for _, m := range c.Methods {
			r.root(m)
		}
		if !c.IsAbstract() {
			r.instantiate(-1, c)
		}
	}
	for len(r.queue) > 0 {
		m := r.queue[0]
		r.queue = r.queue[1:]
		r.process(m)
	}
	r.shortestChains()
	return r
}

func (r *Result) root(m *Method) {
	if m.ID < 0 || r.Opts.CutMethods[m.ID] {
		return
	}
	r.roots = append(r.roots, m.ID)
	if !r.Reached[m.ID] {
		r.Reached[m.ID] = true
		r.queue = append(r.queue, m)
	}
}

func (r *Result) edge(from int, to *Method, kind EdgeKind) {
	if to.ID < 0 || r.Opts.CutMethods[to.ID] {
		return
	}
	if from >= 0 {
		key := uint64(from)<<32 | uint64(to.ID)
		if !r.edgeSet[key] {
			r.edgeSet[key] = true
			r.kinds[key] = kind
			r.edges[from] = append(r.edges[from], to.ID)
		}
	} else if !r.Reached[to.ID] {
		r.roots = append(r.roots, to.ID)
	}
	if !r.Reached[to.ID] {
		r.Reached[to.ID] = true
		r.queue = append(r.queue, to)
	}
}

func (r *Result) cutAway(m *Method, line int) bool {
	for i := range r.Opts.Cuts {
		c := &r.Opts.Cuts[i]
		if c.Class == m.Class.Name && line >= c.From && line <= c.To {
			return true
		}
	}
	return false
}

func (r *Result) process(m *Method) {
	r.initClass(m.ID, m.Class)
	if t, ok := r.Opts.Reflect[m.Key()]; ok {
		r.initClass(m.ID, r.W.mustClass(t))
	}
	events, err := m.Events()
	if err != nil {
		fatalf("%v", err)
	}
	var dead map[int]bool
	if r.Opts.Precision == RTAConst {
		dead = r.W.DeadOffsets(m)
	}
	for i := range events {
		ev := &events[i]
		if r.cutAway(m, ev.Line) || dead[ev.Off] {
			continue
		}
		switch ev.Kind {
		case EvInvoke:
			r.invoke(m.ID, ev.Op, ev.Ref)
		case EvNew:
			c := r.W.mustClass(ev.Ref.Class)
			r.initClass(m.ID, c)
			r.instantiate(m.ID, c)
		case EvStatic:
			if f := r.W.ResolveField(ev.Ref.Class, ev.Ref.Name+":"+ev.Ref.Desc); f != nil {
				r.initClass(m.ID, f.Class)
			} else {
				r.initClass(m.ID, r.W.mustClass(ev.Ref.Class))
			}
		case EvLambda:
			r.handle(m.ID, ev.HandleKind, ev.Ref, EdgeLambda)
			r.instantiate(m.ID, r.W.LambdaClass(ev.Iface))
		case EvHandle:
			r.handle(m.ID, ev.HandleKind, ev.Ref, EdgeLambda)
		case EvConcat:
			for _, t := range ev.Args {
				r.virtual(m.ID, t, "toString:()Ljava/lang/String;")
			}
		}
	}
}

func (r *Result) invoke(from int, op uint8, ref MemberRef) {
	if ref.Class[0] == '[' {
		return
	}
	sig := ref.Name + ":" + ref.Desc
	switch op {
	case opInvokestat:
		if t := r.W.ResolveMethod(ref.Class, sig); t != nil {
			r.initClass(from, t.Class)
			r.edge(from, t, EdgeCall)
		}
	case opInvokespec:
		if t := r.W.ResolveMethod(ref.Class, sig); t != nil && !t.IsAbstract() {
			r.edge(from, t, EdgeCall)
		}
	default:
		r.virtual(from, ref.Class, sig)
	}
}

func (r *Result) handle(from int, kind uint8, ref MemberRef, ek EdgeKind) {
	sig := ref.Name + ":" + ref.Desc
	switch kind {
	case refGetStatic, refPutStatic:
		if f := r.W.ResolveField(ref.Class, sig); f != nil {
			r.initClass(from, f.Class)
		}
	case refInvokeStatic:
		if t := r.W.ResolveMethod(ref.Class, sig); t != nil {
			r.initClass(from, t.Class)
			r.edge(from, t, ek)
		}
	case refInvokeSpecial:
		if t := r.W.ResolveMethod(ref.Class, sig); t != nil && !t.IsAbstract() {
			r.edge(from, t, ek)
		}
	case refNewInvokeSpecial:
		c := r.W.mustClass(ref.Class)
		r.initClass(from, c)
		r.instantiate(from, c)
		if t := c.MethodBy[sig]; t != nil {
			r.edge(from, t, ek)
		}
	case refInvokeVirtual, refInvokeInterface:
		r.virtual(from, ref.Class, sig)
	}
}

// virtual records a virtual call site. A call whose static receiver is a JDK
// type needs nothing here: its application targets are exactly the methods that
// Callbacks already reaches for each instantiated class.
func (r *Result) virtual(from int, class, sig string) {
	if class == "" || class[0] == '[' {
		return
	}
	c := r.W.mustClass(class)
	if !c.App {
		return
	}
	key := class + " " + sig
	v := r.vcalls[key]
	if v == nil {
		v = &vcall{class: class, sig: sig, seen: map[int]bool{}, dispatched: map[string]bool{}}
		r.vcalls[key] = v
		r.byClass[class] = append(r.byClass[class], v)
		res := r.W.ResolveMethod(class, sig)
		if res != nil && res.IsPrivate() {
			v.extra = res
			v.addTarget(r, res)
		} else {
			if res != nil && res.IsPackagePrivate() && !res.IsAbstract() {
				// An override in another package does not replace a package-private
				// method, so selection alone could miss it.
				v.extra = res
			}
			if r.Opts.Precision == CHA {
				for _, k := range r.W.ConcreteSubtypes(class) {
					r.dispatch(v, k)
				}
			}
			for _, name := range r.instSorted() {
				k := r.W.mustClass(name)
				if contains(r.W.Ancestors(k), class) {
					r.dispatch(v, k)
				}
			}
		}
	}
	found := false
	for _, x := range v.callers {
		if x == from {
			found = true
			break
		}
	}
	if !found {
		v.callers = append(v.callers, from)
		for _, t := range v.targets {
			r.edge(from, t, EdgeCall)
		}
	}
}

func (r *Result) instSorted() []string {
	out := make([]string, 0, len(r.Inst))
	for n := range r.Inst {
		out = append(out, n)
	}
	sort.Strings(out)
	return out
}

func (r *Result) dispatch(v *vcall, k *Class) {
	if v.dispatched[k.Name] {
		return
	}
	v.dispatched[k.Name] = true
	if v.extra != nil {
		v.addTarget(r, v.extra)
	}
	for _, t := range r.W.Select(k, v.sig) {
		if t.Class.App {
			v.addTarget(r, t)
		}
	}
}

func (v *vcall) addTarget(r *Result, t *Method) {
	if v.seen[t.ID] {
		return
	}
	v.seen[t.ID] = true
	v.targets = append(v.targets, t)
	for _, c := range v.callers {
		r.edge(c, t, EdgeCall)
	}
}

// instantiate records that an object of class k exists. Its JDK callbacks become
// reachable, with an edge from every site that makes one, so that a call chain
// shows where the object came from.
func (r *Result) instantiate(from int, k *Class) {
	if !k.App && !k.Lambda {
		return
	}
	if !r.Opts.NoCallbacks {
		for _, t := range r.W.Callbacks(k) {
			r.edge(from, t, EdgeCallback)
		}
	}
	if r.Inst[k.Name] {
		return
	}
	r.Inst[k.Name] = true
	for _, a := range r.W.Ancestors(k) {
		for _, v := range r.byClass[a] {
			r.dispatch(v, k)
		}
	}
}

// initClass reaches the static initialisers that JVMS 5.5 runs for a class: its
// own, its superclasses', and those of superinterfaces with default methods.
func (r *Result) initClass(from int, c *Class) {
	if !c.App {
		return
	}
	if m := c.MethodBy["<clinit>:()V"]; m != nil {
		r.edge(from, m, EdgeClinit)
	}
	if c.IsInterface() {
		return
	}
	for _, n := range r.W.Ancestors(c) {
		if n == c.Name {
			continue
		}
		a := r.W.mustClass(n)
		if !a.App {
			continue
		}
		m := a.MethodBy["<clinit>:()V"]
		if m == nil {
			continue
		}
		if a.IsInterface() && !hasDefaultMethod(a) {
			continue
		}
		r.edge(from, m, EdgeClinit)
	}
}

func hasDefaultMethod(c *Class) bool {
	for _, m := range c.Methods {
		if !m.IsStatic() && !m.IsAbstract() && !m.IsPrivate() {
			return true
		}
	}
	return false
}

// shortestChains runs breadth first over the final graph, neighbours in a fixed
// order, so that chains are shortest and stable.
func (r *Result) shortestChains() {
	r.parent = make([]int32, len(r.Reached))
	r.depth = make([]int32, len(r.Reached))
	for i := range r.parent {
		r.parent[i] = -2
	}
	roots := append([]int(nil), r.roots...)
	sort.Ints(roots)
	q := make([]int, 0, len(roots))
	for _, id := range roots {
		if r.parent[id] == -2 {
			r.parent[id] = -1
			q = append(q, id)
		}
	}
	for i := 0; i < len(q); i++ {
		id := q[i]
		out := append([]int(nil), r.edges[id]...)
		sort.Ints(out)
		for _, t := range out {
			if r.parent[t] == -2 {
				r.parent[t] = int32(id)
				r.depth[t] = r.depth[id] + 1
				q = append(q, t)
			}
		}
	}
}

// Chain is the shortest call chain from a root to the method, root first.
func (r *Result) Chain(m *Method) []*Method {
	if m.ID < 0 || !r.Reached[m.ID] || r.parent[m.ID] == -2 {
		return nil
	}
	var rev []*Method
	for id := int32(m.ID); id >= 0; id = r.parent[id] {
		rev = append(rev, r.W.Methods[id])
	}
	out := make([]*Method, len(rev))
	for i, x := range rev {
		out[len(rev)-1-i] = x
	}
	return out
}

// Depth is the number of calls from a root, -1 if unreached.
func (r *Result) Depth(m *Method) int {
	if m.ID < 0 || !r.Reached[m.ID] || r.parent[m.ID] == -2 {
		return -1
	}
	return int(r.depth[m.ID])
}
