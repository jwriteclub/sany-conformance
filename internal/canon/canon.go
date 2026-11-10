// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

// Package canon canonicalises SANY's XML export so that two graphs compare equal
// exactly when they have the same structure. docs/reference.md, "The harness column", says what it removes.
package canon

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

// xmlNode is an element, or a text run when name is empty.
type xmlNode struct {
	name     string
	attrs    []xml.Attr
	text     string
	children []*xmlNode
}

// GraphStats says what the canonicaliser saw, so a surprise is visible.
type GraphStats struct {
	Entries     int // context entries
	Refs        int // <UID> elements outside entry keys
	Unreachable int // entries no walk from the module list reached
	Dangling    int // UIDs referenced with no entry
}

// CanonicalXML turns SANY's XML export into a canonical text form.
//
// It removes exactly three things: (1) the numeric values of <UID>, relabelled
// by order of first reference in a walk from the module list; (2) the order of
// <context> entries, which comes from a HashMap, now sorted by canonical label;
// (3) XML serialisation details: the declaration, CDATA versus escaped text, and
// whitespace-only text between elements. All other text, including whitespace
// inside leaf elements, locations, levels, names and sibling order, is kept.
func CanonicalXML(raw []byte) ([]byte, GraphStats, error) {
	var st GraphStats
	root, err := parseXML(raw)
	if err != nil {
		return nil, st, err
	}
	var context *xmlNode
	for _, c := range root.children {
		if c.name == "context" {
			context = c
			break
		}
	}
	entries := map[string]*xmlNode{}
	var order []string // entries in numeric UID order, for the unreachable tail
	if context != nil {
		for _, e := range context.children {
			if e.name != "entry" {
				continue
			}
			uid := e.childText("UID")
			if uid == "" {
				return nil, st, fmt.Errorf("context entry without UID")
			}
			if _, dup := entries[uid]; dup {
				return nil, st, fmt.Errorf("duplicate context UID %s", uid)
			}
			entries[uid] = e
			order = append(order, uid)
		}
	}
	st.Entries = len(entries)
	sort.Slice(order, func(i, j int) bool {
		a, errA := strconv.Atoi(order[i])
		b, errB := strconv.Atoi(order[j])
		if errA != nil || errB != nil {
			return order[i] < order[j]
		}
		return a < b
	})

	labels := map[string]string{}
	var seq []string // entries with a label, in label order
	var assign func(uid string)
	var walk func(n *xmlNode)
	assign = func(uid string) {
		if _, ok := labels[uid]; ok {
			return
		}
		labels[uid] = "c" + strconv.Itoa(len(labels))
		e, ok := entries[uid]
		if !ok {
			st.Dangling++
			return
		}
		seq = append(seq, uid)
		for _, c := range e.children {
			if c.name == "UID" {
				continue // the key, not a reference
			}
			walk(c)
		}
	}
	walk = func(n *xmlNode) {
		if n.name == "UID" {
			st.Refs++
			assign(n.textOnly())
			return
		}
		for _, c := range n.children {
			walk(c)
		}
	}
	for _, c := range root.children {
		if c != context {
			walk(c)
		}
	}
	reached := len(seq)
	for _, uid := range order {
		assign(uid)
	}
	st.Unreachable = len(seq) - reached

	var b bytes.Buffer
	fmt.Fprintf(&b, "%s\n", root.name)
	for _, c := range root.children {
		if c == context {
			b.WriteString(" context\n")
			for _, uid := range seq {
				emit(&b, entries[uid], 2, labels)
			}
			continue
		}
		emit(&b, c, 1, labels)
	}
	return b.Bytes(), st, nil
}

func parseXML(raw []byte) (*xmlNode, error) {
	dec := xml.NewDecoder(bytes.NewReader(raw))
	var stack []*xmlNode
	var root *xmlNode
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			n := &xmlNode{name: t.Name.Local, attrs: t.Copy().Attr}
			if len(stack) > 0 {
				p := stack[len(stack)-1]
				p.children = append(p.children, n)
			} else if root == nil {
				root = n
			} else {
				return nil, fmt.Errorf("more than one root element")
			}
			stack = append(stack, n)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if len(stack) == 0 {
				continue
			}
			p := stack[len(stack)-1]
			if k := len(p.children); k > 0 && p.children[k-1].name == "" {
				p.children[k-1].text += string(t)
			} else {
				p.children = append(p.children, &xmlNode{text: string(t)})
			}
		}
	}
	if root == nil {
		return nil, fmt.Errorf("no root element")
	}
	stripInterElementSpace(root)
	return root, nil
}

// stripInterElementSpace drops whitespace-only text from elements that also
// have element children. Text in leaf elements is never touched.
func stripInterElementSpace(n *xmlNode) {
	hasElem := false
	for _, c := range n.children {
		if c.name != "" {
			hasElem = true
		}
	}
	if hasElem {
		kept := n.children[:0]
		for _, c := range n.children {
			if c.name == "" && strings.TrimSpace(c.text) == "" {
				continue
			}
			kept = append(kept, c)
		}
		n.children = kept
	}
	for _, c := range n.children {
		if c.name != "" {
			stripInterElementSpace(c)
		}
	}
}

func (n *xmlNode) textOnly() string {
	var s strings.Builder
	for _, c := range n.children {
		if c.name == "" {
			s.WriteString(c.text)
		}
	}
	return s.String()
}

func (n *xmlNode) childText(name string) string {
	for _, c := range n.children {
		if c.name == name {
			return c.textOnly()
		}
	}
	return ""
}

func emit(b *bytes.Buffer, n *xmlNode, depth int, labels map[string]string) {
	b.WriteString(strings.Repeat(" ", depth))
	if n.name == "" {
		b.WriteString("#text " + strconv.Quote(n.text) + "\n")
		return
	}
	b.WriteString(n.name)
	for _, a := range n.attrs {
		b.WriteString(" @" + a.Name.Local + "=" + strconv.Quote(a.Value))
	}
	if n.name == "UID" {
		b.WriteString(" " + labels[n.textOnly()] + "\n")
		return
	}
	if len(n.children) == 1 && n.children[0].name == "" {
		b.WriteString(" = " + strconv.Quote(n.children[0].text) + "\n")
		return
	}
	b.WriteByte('\n')
	for _, c := range n.children {
		emit(b, c, depth+1, labels)
	}
}
