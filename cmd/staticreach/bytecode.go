// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/binary"
	"fmt"
	"sort"
)

// Opcodes the analysis cares about, JVMS 6.5.
const (
	opAconstNull = 1
	opIconstM1   = 2
	opIconst5    = 8
	opBipush     = 16
	opSipush     = 17
	opLdc        = 18
	opLdcW       = 19
	opLdc2W      = 20
	opIfeq       = 153
	opIfne       = 154
	opIflt       = 155
	opIfge       = 156
	opIfgt       = 157
	opIfle       = 158
	opIfIcmpeq   = 159
	opIfIcmpne   = 160
	opIfIcmplt   = 161
	opIfIcmpge   = 162
	opIfIcmpgt   = 163
	opIfIcmple   = 164
	opIfAcmpeq   = 165
	opIfAcmpne   = 166
	opGoto       = 167
	opJsr        = 168
	opRet        = 169
	opTableswi   = 170
	opLookupswi  = 171
	opIreturn    = 172
	opReturn     = 177
	opGetstatic  = 178
	opPutstatic  = 179
	opGetfield   = 180
	opPutfield   = 181
	opInvokevirt = 182
	opInvokespec = 183
	opInvokestat = 184
	opInvokeifc  = 185
	opInvokedyn  = 186
	opNew        = 187
	opAthrow     = 191
	opWide       = 196
	opIfnull     = 198
	opIfnonnull  = 199
	opGotoW      = 200
	opJsrW       = 201
)

// opLength holds the fixed length of each opcode; 0 means variable or invalid.
var opLength = func() [256]uint8 {
	var t [256]uint8
	set := func(lo, hi int, n uint8) {
		for i := lo; i <= hi; i++ {
			t[i] = n
		}
	}
	set(0, 15, 1)
	set(16, 16, 2)
	set(17, 17, 3)
	set(18, 18, 2)
	set(19, 20, 3)
	set(21, 25, 2)
	set(26, 53, 1)
	set(54, 58, 2)
	set(59, 131, 1)
	set(132, 132, 3)
	set(133, 152, 1)
	set(153, 168, 3)
	set(169, 169, 2)
	set(172, 177, 1)
	set(178, 184, 3)
	set(185, 186, 5)
	set(187, 187, 3)
	set(188, 188, 2)
	set(189, 189, 3)
	set(190, 191, 1)
	set(192, 193, 3)
	set(194, 195, 1)
	set(197, 197, 4)
	set(198, 199, 3)
	set(200, 201, 5)
	return t
}()

// Insn is one decoded instruction.
type Insn struct {
	Off     int
	Op      uint8
	Len     int
	Index   uint16 // constant pool index, when the opcode has one
	Targets []int  // jump or switch targets, absolute
}

// IsConditional says that the instruction is a two-way branch.
func (i Insn) IsConditional() bool {
	return (i.Op >= opIfeq && i.Op <= opIfAcmpne) || i.Op == opIfnull || i.Op == opIfnonnull
}

// IsSwitch says that the instruction is a table or lookup switch.
func (i Insn) IsSwitch() bool { return i.Op == opTableswi || i.Op == opLookupswi }

// Decode returns the instructions of the method, in order.
func (m *Method) Decode() ([]Insn, error) {
	if m.insns != nil || len(m.Code) == 0 {
		return m.insns, nil
	}
	code := m.Code
	var out []Insn
	be := binary.BigEndian
	for off := 0; off < len(code); {
		op := code[off]
		in := Insn{Off: off, Op: op}
		n := int(opLength[op])
		switch op {
		case opTableswi, opLookupswi:
			p := (off + 4) &^ 3
			if p+8 > len(code) {
				return nil, fmt.Errorf("%s: switch at %d runs off the code", m.Key(), off)
			}
			def := int(int32(be.Uint32(code[p:])))
			in.Targets = append(in.Targets, off+def)
			if op == opTableswi {
				if p+12 > len(code) {
					return nil, fmt.Errorf("%s: tableswitch at %d runs off the code", m.Key(), off)
				}
				low := int(int32(be.Uint32(code[p+4:])))
				high := int(int32(be.Uint32(code[p+8:])))
				cnt := high - low + 1
				p += 12
				if cnt < 0 || p+4*cnt > len(code) {
					return nil, fmt.Errorf("%s: tableswitch at %d runs off the code", m.Key(), off)
				}
				for i := 0; i < cnt; i++ {
					in.Targets = append(in.Targets, off+int(int32(be.Uint32(code[p+4*i:]))))
				}
				n = p + 4*cnt - off
			} else {
				cnt := int(int32(be.Uint32(code[p+4:])))
				p += 8
				if cnt < 0 || p+8*cnt > len(code) {
					return nil, fmt.Errorf("%s: lookupswitch at %d runs off the code", m.Key(), off)
				}
				for i := 0; i < cnt; i++ {
					in.Targets = append(in.Targets, off+int(int32(be.Uint32(code[p+8*i+4:]))))
				}
				n = p + 8*cnt - off
			}
		case opWide:
			if off+1 >= len(code) {
				return nil, fmt.Errorf("%s: wide at %d runs off the code", m.Key(), off)
			}
			n = 4
			if code[off+1] == 132 { // iinc
				n = 6
			}
		default:
			if n == 0 {
				return nil, fmt.Errorf("%s: unknown opcode %d at %d", m.Key(), op, off)
			}
		}
		if off+n > len(code) {
			return nil, fmt.Errorf("%s: instruction at %d runs off the code", m.Key(), off)
		}
		switch {
		case op == opLdc:
			in.Index = uint16(code[off+1])
		case op == opLdcW || op == opLdc2W || (op >= opGetstatic && op <= opInvokedyn) || op == opNew:
			in.Index = be.Uint16(code[off+1:])
		case in.IsConditional() || op == opGoto || op == opJsr:
			in.Targets = []int{off + int(int16(be.Uint16(code[off+1:])))}
		case op == opGotoW || op == opJsrW:
			in.Targets = []int{off + int(int32(be.Uint32(code[off+1:])))}
		}
		in.Len = n
		out = append(out, in)
		off += n
	}
	m.insns = out
	return out, nil
}

// EventKind classifies a bytecode site that can make other code run.
type EventKind uint8

const (
	// EvInvoke is an invokevirtual, invokeinterface, invokespecial or invokestatic.
	EvInvoke EventKind = iota
	// EvNew is a new instruction: the class is instantiated and initialised.
	EvNew
	// EvStatic is a getstatic or putstatic: the declaring class is initialised.
	EvStatic
	// EvLambda is an invokedynamic through LambdaMetafactory.
	EvLambda
	// EvConcat is an invokedynamic through StringConcatFactory.
	EvConcat
	// EvHandle is a method handle constant loaded with ldc.
	EvHandle
)

// Event is one such site.
type Event struct {
	Kind EventKind
	Off  int
	Line int
	Op   uint8
	Ref  MemberRef
	// Handle is the implementation method handle kind of a lambda or ldc.
	HandleKind uint8
	// Iface is the interface a lambda implements; extra markers follow it.
	Iface []string
	// Args are the reference type arguments of a string concatenation.
	Args []string
}

// Events lists the method's sites in bytecode order.
func (m *Method) Events() ([]Event, error) {
	if m.events != nil || len(m.Code) == 0 {
		return m.events, nil
	}
	insns, err := m.Decode()
	if err != nil {
		return nil, err
	}
	c := m.Class
	out := []Event{}
	for _, in := range insns {
		line := 0
		switch in.Op {
		case opInvokevirt, opInvokespec, opInvokestat, opInvokeifc:
			line = m.LineAt(in.Off)
			out = append(out, Event{Kind: EvInvoke, Off: in.Off, Line: line, Op: in.Op, Ref: c.memberRef(in.Index)})
		case opGetstatic, opPutstatic:
			line = m.LineAt(in.Off)
			out = append(out, Event{Kind: EvStatic, Off: in.Off, Line: line, Op: in.Op, Ref: c.memberRef(in.Index)})
		case opNew:
			line = m.LineAt(in.Off)
			out = append(out, Event{Kind: EvNew, Off: in.Off, Line: line, Op: in.Op, Ref: MemberRef{Class: c.className(in.Index)}})
		case opInvokedyn:
			line = m.LineAt(in.Off)
			ev, err := c.indyEvent(in.Index)
			if err != nil {
				return nil, fmt.Errorf("%s at %d: %w", m.Key(), in.Off, err)
			}
			ev.Off, ev.Line, ev.Op = in.Off, line, in.Op
			out = append(out, ev)
		case opLdc, opLdcW:
			if e := c.entry(in.Index); e.tag == tagMethodHandle {
				line = m.LineAt(in.Off)
				out = append(out, Event{Kind: EvHandle, Off: in.Off, Line: line, Op: in.Op, HandleKind: uint8(e.a), Ref: c.memberRef(e.b)})
			}
		}
	}
	m.events = out
	return out, nil
}

// indyEvent reads the real bootstrap arguments of an invokedynamic.
func (c *Class) indyEvent(index uint16) (Event, error) {
	e := c.entry(index)
	if e.tag != tagInvokeDynamic {
		return Event{}, fmt.Errorf("invokedynamic constant has tag %d", e.tag)
	}
	if int(e.a) >= len(c.bootstraps) {
		return Event{}, fmt.Errorf("bootstrap index %d out of range", e.a)
	}
	boot := c.bootstraps[e.a]
	bh := c.entry(boot.Handle)
	owner := c.memberRef(bh.b)
	nt := c.entry(e.b)
	desc := c.utf(nt.b)
	args, ret := parseDesc(desc)
	switch owner.Class + "." + owner.Name {
	case "java/lang/invoke/LambdaMetafactory.metafactory", "java/lang/invoke/LambdaMetafactory.altMetafactory":
		if len(boot.Args) < 2 {
			return Event{}, fmt.Errorf("lambda bootstrap with %d arguments", len(boot.Args))
		}
		impl := c.entry(boot.Args[1])
		if impl.tag != tagMethodHandle {
			return Event{}, fmt.Errorf("lambda implementation constant has tag %d", impl.tag)
		}
		ev := Event{Kind: EvLambda, HandleKind: uint8(impl.a), Ref: c.memberRef(impl.b), Iface: []string{descClass(ret)}}
		if owner.Name == "altMetafactory" {
			ev.Iface = append(ev.Iface, c.altMarkers(boot.Args)...)
		}
		return ev, nil
	case "java/lang/invoke/StringConcatFactory.makeConcatWithConstants", "java/lang/invoke/StringConcatFactory.makeConcat":
		ev := Event{Kind: EvConcat}
		for _, a := range args {
			if n := descClass(a); n != "" {
				ev.Args = append(ev.Args, n)
			}
		}
		return ev, nil
	}
	return Event{}, fmt.Errorf("unmodelled invokedynamic bootstrap %s.%s", owner.Class, owner.Name)
}

// altMarkers reads the marker interfaces of an altMetafactory call. The flags
// argument is index 3; FLAG_MARKERS is 2 and FLAG_SERIALIZABLE is 1.
func (c *Class) altMarkers(args []uint16) []string {
	if len(args) < 4 {
		return nil
	}
	flags := c.entry(args[3]).n
	var out []string
	i := 4
	if flags&1 != 0 {
		out = append(out, "java/io/Serializable")
	}
	if flags&2 != 0 && i < len(args) {
		n := int(c.entry(args[i]).n)
		i++
		for j := 0; j < n && i < len(args); j++ {
			out = append(out, c.className(args[i]))
			i++
		}
	}
	return out
}

// parseDesc splits a method descriptor into argument and return types.
func parseDesc(desc string) ([]string, string) {
	var args []string
	i := 1
	one := func() string {
		s := i
		for i < len(desc) && desc[i] == '[' {
			i++
		}
		if i < len(desc) && desc[i] == 'L' {
			for i < len(desc) && desc[i] != ';' {
				i++
			}
		}
		i++
		if i > len(desc) {
			i = len(desc)
		}
		return desc[s:i]
	}
	for i < len(desc) && desc[i] != ')' {
		args = append(args, one())
	}
	i++
	return args, one()
}

// descClass is the class of an object type descriptor, "" for arrays and primitives.
func descClass(d string) string {
	if len(d) > 2 && d[0] == 'L' {
		return d[1 : len(d)-1]
	}
	return ""
}

// conditionalOutcomes counts the outcomes JaCoCo would count for a branching
// instruction: two for a conditional, one per distinct target for a switch.
func conditionalOutcomes(in Insn) int {
	switch {
	case in.IsConditional():
		return 2
	case in.IsSwitch():
		t := append([]int(nil), in.Targets...)
		sort.Ints(t)
		n := 0
		for i, v := range t {
			if i == 0 || v != t[i-1] {
				n++
			}
		}
		return n
	}
	return 0
}
