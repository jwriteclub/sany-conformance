// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/binary"
	"errors"
	"fmt"
)

// Constant pool tags, JVMS 4.4.
const (
	tagUtf8          = 1
	tagInteger       = 3
	tagFloat         = 4
	tagLong          = 5
	tagDouble        = 6
	tagClass         = 7
	tagString        = 8
	tagFieldref      = 9
	tagMethodref     = 10
	tagIfaceMethod   = 11
	tagNameAndType   = 12
	tagMethodHandle  = 15
	tagMethodType    = 16
	tagDynamic       = 17
	tagInvokeDynamic = 18
	tagModule        = 19
	tagPackage       = 20
)

// Access flags, JVMS 4.1 and 4.6.
const (
	accPublic    = 0x0001
	accPrivate   = 0x0002
	accProtected = 0x0004
	accStatic    = 0x0008
	accFinal     = 0x0010
	accBridge    = 0x0040
	accNative    = 0x0100
	accInterface = 0x0200
	accAbstract  = 0x0400
	accSynthetic = 0x1000
)

// Method handle kinds, JVMS 5.4.3.5.
const (
	refGetField         = 1
	refGetStatic        = 2
	refPutField         = 3
	refPutStatic        = 4
	refInvokeVirtual    = 5
	refInvokeStatic     = 6
	refInvokeSpecial    = 7
	refNewInvokeSpecial = 8
	refInvokeInterface  = 9
)

type cpEntry struct {
	tag  uint8
	a, b uint16 // indices, or kind and index for a method handle
	s    string // Utf8
	n    int64  // Integer, or Long
}

// Field is a declared field.
type Field struct {
	Class  *Class
	Name   string
	Desc   string
	Access uint16
	// HasConst says that the field carries a ConstantValue attribute.
	HasConst bool
}

// Key names the field inside the whole world.
func (f *Field) Key() string { return f.Class.Name + "." + f.Name + ":" + f.Desc }

// LineEntry is one LineNumberTable row.
type LineEntry struct {
	Off  int
	Line int
}

// Handler is one exception table row.
type Handler struct {
	Start, End, Target int
}

// Method is a declared method. ID is set only for application methods.
type Method struct {
	Class    *Class
	Name     string
	Desc     string
	Access   uint16
	Code     []byte
	Handlers []Handler
	Lines    []LineEntry
	ID       int

	insns  []Insn
	events []Event
}

// Key names the method inside the whole world.
func (m *Method) Key() string { return m.Class.Name + "." + m.Name + ":" + m.Desc }

// Sig is name plus descriptor.
func (m *Method) Sig() string { return m.Name + ":" + m.Desc }

// IsStatic says that the method is static.
func (m *Method) IsStatic() bool { return m.Access&accStatic != 0 }

// IsAbstract says that the method has no body to select.
func (m *Method) IsAbstract() bool { return m.Access&accAbstract != 0 }

// IsPrivate says that the method is private.
func (m *Method) IsPrivate() bool { return m.Access&accPrivate != 0 }

// IsPackagePrivate says that the method is visible only inside its package.
func (m *Method) IsPackagePrivate() bool {
	return m.Access&(accPublic|accProtected|accPrivate) == 0
}

// LineAt is the source line of a bytecode offset, 0 when the method has no table.
func (m *Method) LineAt(off int) int {
	line := 0
	best := -1
	for _, e := range m.Lines {
		if e.Off <= off && e.Off > best {
			best = e.Off
			line = e.Line
		}
	}
	return line
}

// Bootstrap is one BootstrapMethods row.
type Bootstrap struct {
	Handle uint16
	Args   []uint16
}

// Class is one parsed class file.
type Class struct {
	Name       string
	Super      string
	Interfaces []string
	Access     uint16
	SourceFile string
	Fields     []*Field
	Methods    []*Method
	MethodBy   map[string]*Method // by Sig
	FieldBy    map[string]*Field  // by name and descriptor

	// App marks a class of the analysed program, as opposed to the JDK.
	App bool
	// Missing marks a referenced class that no input provides.
	Missing bool
	// Lambda marks a modelled class for a lambda's runtime object.
	Lambda bool
	// Origin is the index of the input directory the class came from.
	Origin int

	cp         []cpEntry
	bootstraps []Bootstrap
}

// IsInterface says that the class is an interface.
func (c *Class) IsInterface() bool { return c.Access&accInterface != 0 }

// IsAbstract says that the class cannot be instantiated directly.
func (c *Class) IsAbstract() bool { return c.Access&(accAbstract|accInterface) != 0 }

type reader struct {
	b   []byte
	off int
	err error
}

func (r *reader) need(n int) bool {
	if r.err != nil {
		return false
	}
	if n < 0 || r.off+n > len(r.b) {
		r.err = errors.New("truncated class file")
		return false
	}
	return true
}

func (r *reader) u1() uint8 {
	if !r.need(1) {
		return 0
	}
	v := r.b[r.off]
	r.off++
	return v
}

func (r *reader) u2() uint16 {
	if !r.need(2) {
		return 0
	}
	v := binary.BigEndian.Uint16(r.b[r.off:])
	r.off += 2
	return v
}

func (r *reader) u4() uint32 {
	if !r.need(4) {
		return 0
	}
	v := binary.BigEndian.Uint32(r.b[r.off:])
	r.off += 4
	return v
}

func (r *reader) bytes(n int) []byte {
	if !r.need(n) {
		return nil
	}
	v := r.b[r.off : r.off+n]
	r.off += n
	return v
}

// ParseClass reads a class file. The identifiers are kept in their modified
// UTF-8 bytes: only equality matters here, and the bytes are stable.
func ParseClass(data []byte) (*Class, error) {
	r := &reader{b: data}
	if magic := r.u4(); magic != 0xCAFEBABE {
		return nil, fmt.Errorf("bad magic %x", magic)
	}
	r.u2() // minor
	r.u2() // major
	count := int(r.u2())
	cp := make([]cpEntry, count)
	for i := 1; i < count; i++ {
		e := cpEntry{tag: r.u1()}
		switch e.tag {
		case tagUtf8:
			e.s = string(r.bytes(int(r.u2())))
		case tagInteger:
			e.n = int64(int32(r.u4()))
		case tagFloat:
			r.u4()
		case tagLong:
			e.n = int64(uint64(r.u4())<<32 | uint64(r.u4()))
		case tagDouble:
			r.u4()
			r.u4()
		case tagClass, tagString, tagMethodType, tagModule, tagPackage:
			e.a = r.u2()
		case tagFieldref, tagMethodref, tagIfaceMethod, tagNameAndType, tagDynamic, tagInvokeDynamic:
			e.a = r.u2()
			e.b = r.u2()
		case tagMethodHandle:
			e.a = uint16(r.u1())
			e.b = r.u2()
		default:
			return nil, fmt.Errorf("constant pool tag %d at %d", e.tag, i)
		}
		cp[i] = e
		if e.tag == tagLong || e.tag == tagDouble {
			i++
		}
		if r.err != nil {
			return nil, r.err
		}
	}
	c := &Class{cp: cp, MethodBy: map[string]*Method{}, FieldBy: map[string]*Field{}}
	c.Access = r.u2()
	c.Name = c.className(r.u2())
	if s := r.u2(); s != 0 {
		c.Super = c.className(s)
	}
	for n := int(r.u2()); n > 0; n-- {
		c.Interfaces = append(c.Interfaces, c.className(r.u2()))
	}
	for n := int(r.u2()); n > 0; n-- {
		f := &Field{Class: c, Access: r.u2(), Name: c.utf(r.u2()), Desc: c.utf(r.u2())}
		for a := int(r.u2()); a > 0; a-- {
			name := c.utf(r.u2())
			body := r.bytes(int(r.u4()))
			if name == "ConstantValue" && body != nil {
				f.HasConst = true
			}
		}
		c.Fields = append(c.Fields, f)
		c.FieldBy[f.Name+":"+f.Desc] = f
	}
	for n := int(r.u2()); n > 0; n-- {
		m := &Method{Class: c, Access: r.u2(), Name: c.utf(r.u2()), Desc: c.utf(r.u2()), ID: -1}
		for a := int(r.u2()); a > 0; a-- {
			name := c.utf(r.u2())
			body := r.bytes(int(r.u4()))
			if name == "Code" && body != nil {
				if err := c.parseCode(m, body); err != nil {
					return nil, fmt.Errorf("%s: %w", m.Key(), err)
				}
			}
		}
		c.Methods = append(c.Methods, m)
		c.MethodBy[m.Sig()] = m
	}
	for a := int(r.u2()); a > 0; a-- {
		name := c.utf(r.u2())
		body := r.bytes(int(r.u4()))
		if body == nil {
			break
		}
		switch name {
		case "SourceFile":
			br := &reader{b: body}
			c.SourceFile = c.utf(br.u2())
		case "BootstrapMethods":
			br := &reader{b: body}
			for k := int(br.u2()); k > 0; k-- {
				b := Bootstrap{Handle: br.u2()}
				for j := int(br.u2()); j > 0; j-- {
					b.Args = append(b.Args, br.u2())
				}
				c.bootstraps = append(c.bootstraps, b)
			}
			if br.err != nil {
				return nil, br.err
			}
		}
	}
	if r.err != nil {
		return nil, r.err
	}
	return c, nil
}

func (c *Class) parseCode(m *Method, body []byte) error {
	r := &reader{b: body}
	r.u2() // max_stack
	r.u2() // max_locals
	m.Code = r.bytes(int(r.u4()))
	for n := int(r.u2()); n > 0; n-- {
		h := Handler{Start: int(r.u2()), End: int(r.u2()), Target: int(r.u2())}
		r.u2() // catch type
		m.Handlers = append(m.Handlers, h)
	}
	for a := int(r.u2()); a > 0; a-- {
		name := c.utf(r.u2())
		attr := r.bytes(int(r.u4()))
		if name != "LineNumberTable" || attr == nil {
			continue
		}
		lr := &reader{b: attr}
		for k := int(lr.u2()); k > 0; k-- {
			m.Lines = append(m.Lines, LineEntry{Off: int(lr.u2()), Line: int(lr.u2())})
		}
		if lr.err != nil {
			return lr.err
		}
	}
	return r.err
}

func (c *Class) entry(i uint16) cpEntry {
	if int(i) >= len(c.cp) {
		return cpEntry{}
	}
	return c.cp[i]
}

func (c *Class) utf(i uint16) string { return c.entry(i).s }

func (c *Class) className(i uint16) string { return c.utf(c.entry(i).a) }

// MemberRef is a resolved constant pool Fieldref or Methodref.
type MemberRef struct {
	Class, Name, Desc string
	// Iface says that the constant was an InterfaceMethodref.
	Iface bool
}

func (c *Class) memberRef(i uint16) MemberRef {
	e := c.entry(i)
	nt := c.entry(e.b)
	return MemberRef{Class: c.className(e.a), Name: c.utf(nt.a), Desc: c.utf(nt.b), Iface: e.tag == tagIfaceMethod}
}
