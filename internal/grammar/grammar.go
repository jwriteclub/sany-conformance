// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package grammar

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// A small reader of tla+.jj and a sentence generator in the manner of Purdom: for
// every choice alternative, optional part and loop in the grammar, one short
// sentence from the start symbol that uses it. Semantic lookahead and Java actions
// are ignored, so some sentences are rejected; those exercise the error arms.

type gKind int

const (
	gSeq gKind = iota
	gChoice
	gOpt
	gStar
	gPlus
	gTok
	gNT
)

// GNode is one grammar node.
type GNode struct {
	Kind gKind
	Kids []*GNode
	Text string // token text or nonterminal name
	Prod string
}

// Grammar is what was read.
type Grammar struct {
	Prods  map[string]*GNode
	Names  []string
	minLen map[string]int
}

type jjLexer struct {
	s string
	i int
}

// next returns the next token of a .jj file, skipping comments.
func (l *jjLexer) next() string {
	for l.i < len(l.s) {
		c := l.s[l.i]
		switch {
		case c == ' ' || c == '\t' || c == '\n' || c == '\r':
			l.i++
		case strings.HasPrefix(l.s[l.i:], "//"):
			if j := strings.IndexByte(l.s[l.i:], '\n'); j >= 0 {
				l.i += j
			} else {
				l.i = len(l.s)
			}
		case strings.HasPrefix(l.s[l.i:], "/*"):
			if j := strings.Index(l.s[l.i:], "*/"); j >= 0 {
				l.i += j + 2
			} else {
				l.i = len(l.s)
			}
		case c == '"':
			j := l.i + 1
			for j < len(l.s) && l.s[j] != '"' {
				if l.s[j] == '\\' {
					j++
				}
				j++
			}
			t := l.s[l.i : j+1]
			l.i = j + 1
			return t
		case c == '\'':
			j := l.i + 1
			for j < len(l.s) && l.s[j] != '\'' {
				if l.s[j] == '\\' {
					j++
				}
				j++
			}
			t := l.s[l.i : j+1]
			l.i = j + 1
			return t
		case c == '_' || c == '$' || c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z':
			j := l.i
			for j < len(l.s) && (l.s[j] == '_' || l.s[j] == '$' || l.s[j] >= '0' && l.s[j] <= '9' || l.s[j] >= 'a' && l.s[j] <= 'z' || l.s[j] >= 'A' && l.s[j] <= 'Z') {
				j++
			}
			t := l.s[l.i:j]
			l.i = j
			return t
		default:
			l.i++
			return string(c)
		}
	}
	return ""
}

func (l *jjLexer) peek() string {
	i := l.i
	t := l.next()
	l.i = i
	return t
}

// skipBalanced consumes through the closer of an opener already read.
func (l *jjLexer) skipBalanced(open, close string) {
	depth := 1
	for depth > 0 {
		t := l.next()
		if t == "" {
			return
		}
		if t == open {
			depth++
		} else if t == close {
			depth--
		}
	}
}

func (l *jjLexer) choices(prod string) *GNode {
	alts := []*GNode{l.seq(prod)}
	for l.peek() == "|" {
		l.next()
		alts = append(alts, l.seq(prod))
	}
	if len(alts) == 1 {
		return alts[0]
	}
	return &GNode{Kind: gChoice, Kids: alts, Prod: prod}
}

func (l *jjLexer) seq(prod string) *GNode {
	n := &GNode{Kind: gSeq, Prod: prod}
	for {
		switch t := l.peek(); t {
		case "", "|", ")", "]", "}":
			return n
		case "LOOKAHEAD":
			l.next()
			l.next() // (
			l.skipBalanced("(", ")")
		case "{":
			l.next()
			l.skipBalanced("{", "}")
		case "[":
			l.next()
			c := l.choices(prod)
			l.next() // ]
			n.Kids = append(n.Kids, &GNode{Kind: gOpt, Kids: []*GNode{c}, Prod: prod})
		case "(":
			l.next()
			c := l.choices(prod)
			l.next() // )
			switch l.peek() {
			case "*":
				l.next()
				n.Kids = append(n.Kids, &GNode{Kind: gStar, Kids: []*GNode{c}, Prod: prod})
			case "+":
				l.next()
				n.Kids = append(n.Kids, &GNode{Kind: gPlus, Kids: []*GNode{c}, Prod: prod})
			case "?":
				l.next()
				n.Kids = append(n.Kids, &GNode{Kind: gOpt, Kids: []*GNode{c}, Prod: prod})
			default:
				n.Kids = append(n.Kids, c)
			}
		case "try":
			l.next()
			l.next() // {
			c := l.choices(prod)
			l.next() // }
			for l.peek() == "catch" || l.peek() == "finally" {
				if l.next() == "catch" {
					l.next() // (
					l.skipBalanced("(", ")")
				}
				l.next() // {
				l.skipBalanced("{", "}")
			}
			n.Kids = append(n.Kids, c)
		case "<":
			l.next()
			name := l.next()
			l.next() // >
			n.Kids = append(n.Kids, &GNode{Kind: gTok, Text: "<" + name + ">", Prod: prod})
		default:
			if strings.HasPrefix(t, "\"") {
				l.next()
				n.Kids = append(n.Kids, &GNode{Kind: gTok, Text: t, Prod: prod})
				continue
			}
			name := l.next()
			// An assignment target: x =, a[0] =, a.b =.
			for l.peek() == "[" || l.peek() == "." {
				if l.next() == "[" {
					l.skipBalanced("[", "]")
				} else {
					l.next()
				}
			}
			if l.peek() == "=" {
				l.next()
				continue
			}
			if l.peek() == "(" {
				l.next()
				l.skipBalanced("(", ")")
				n.Kids = append(n.Kids, &GNode{Kind: gNT, Text: name, Prod: prod})
				continue
			}
			// A bare word: not grammar. Skip it rather than fail.
		}
	}
}

// ParseGrammar reads the productions of a .jj file.
func ParseGrammar(path string) (*Grammar, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	text := string(b)
	i := strings.Index(text, "PARSER_END")
	if i < 0 {
		return nil, fmt.Errorf("no PARSER_END in %s", path)
	}
	l := &jjLexer{s: text, i: i}
	l.next()
	l.next()
	l.skipBalanced("(", ")")
	g := &Grammar{Prods: map[string]*GNode{}}
	var last string
	for {
		t := l.next()
		if t == "" {
			break
		}
		switch t {
		case "TOKEN", "SKIP", "MORE", "SPECIAL_TOKEN", "TOKEN_MGR_DECLS":
			// : { ... } at top level
			l.next()
			if l.next() == "{" {
				l.skipBalanced("{", "}")
			}
			continue
		case "<":
			l.skipBalanced("<", ">") // <STATE> TOKEN
			continue
		case "(":
			// Argument list of a production: name ( args ) : { java } { expansion }
			if last == "" {
				l.skipBalanced("(", ")")
				continue
			}
			l.skipBalanced("(", ")")
			if l.next() != ":" {
				last = ""
				continue
			}
			if l.next() != "{" {
				continue
			}
			l.skipBalanced("{", "}")
			if l.next() != "{" {
				continue
			}
			g.Prods[last] = l.choices(last)
			g.Names = append(g.Names, last)
			l.next() // }
			last = ""
			continue
		}
		last = t
	}
	sort.Strings(g.Names)
	g.computeMin()
	return g, nil
}

func (g *Grammar) computeMin() {
	g.minLen = map[string]int{}
	const inf = 1 << 20
	for _, n := range g.Names {
		g.minLen[n] = inf
	}
	for changed := true; changed; {
		changed = false
		for _, n := range g.Names {
			if m := g.nodeMin(g.Prods[n]); m < g.minLen[n] {
				g.minLen[n] = m
				changed = true
			}
		}
	}
}

func (g *Grammar) nodeMin(n *GNode) int {
	const inf = 1 << 20
	switch n.Kind {
	case gTok:
		return 1
	case gNT:
		if m, ok := g.minLen[n.Text]; ok {
			return m
		}
		return inf
	case gOpt, gStar:
		return 0
	case gPlus:
		return g.nodeMin(n.Kids[0])
	case gChoice:
		best := inf
		for _, k := range n.Kids {
			if m := g.nodeMin(k); m < best {
				best = m
			}
		}
		return best
	}
	t := 0
	for _, k := range n.Kids {
		t += g.nodeMin(k)
		if t > inf {
			t = inf
		}
	}
	return t
}

// tokenText gives a sample of a named token. Unknown names make a plain identifier.
func (g *Grammar) tokenText(t string, table map[string]string) string {
	if strings.HasPrefix(t, "\"") {
		return dictRaw(strings.Trim(t, "\""))
	}
	name := strings.Trim(t, "<>")
	if s, ok := table[name]; ok {
		return s
	}
	return "x"
}

var reTokDef = regexp.MustCompile(`<\s*#?\s*(\w+)\s*:\s*"((?:[^"\\\n]|\\.)*)"`)

// TokenTable samples each named token from its first literal, with the regular
// tokens by hand.
func TokenTable(jj string) (map[string]string, error) {
	b, err := os.ReadFile(jj)
	if err != nil {
		return nil, err
	}
	t := map[string]string{}
	for _, m := range reTokDef.FindAllStringSubmatch(string(b), -1) {
		t[m[1]] = dictRaw(m[2])
	}
	for k, v := range map[string]string{
		"IDENTIFIER": "x", "NUMBER": "1", "NUMBER_LITERAL": "1", "STRING_LITERAL": "\"s\"", "_BM0": "---- MODULE", "_BM1": "---- MODULE", "_BM2": "---- MODULE",
		"BEGIN_MODULE": "---- MODULE", "SEPARATOR": "----", "END_MODULE": "====", "ProofStepLexeme": "<1>1", "ProofImplicitStepLexeme": "<*>1", "ProofStepDotLexeme": "<1>1.",
		"ProofLevelLexeme": "<1>", "ProofImplicitLevelLexeme": "<*>", "ProofImplicitLevelDotLexeme": "<*>.",
	} {
		t[k] = v
	}
	return t, nil
}

// Generator holds a walk's state.
type sentenceGen struct {
	g     *Grammar
	table map[string]string
	force map[*GNode]int // choice alternative, or 1 for take an Opt or loop
	out   []string
	depth int
}

func (s *sentenceGen) emit(n *GNode) {
	if s.depth > 400 {
		return
	}
	switch n.Kind {
	case gTok:
		s.out = append(s.out, s.g.tokenText(n.Text, s.table))
	case gNT:
		if p, ok := s.g.Prods[n.Text]; ok {
			s.depth++
			s.emit(p)
			s.depth--
		}
	case gSeq:
		for _, k := range n.Kids {
			s.emit(k)
		}
	case gOpt, gStar:
		if _, ok := s.force[n]; ok {
			delete(s.force, n)
			s.emit(n.Kids[0])
		}
	case gPlus:
		s.emit(n.Kids[0])
		if _, ok := s.force[n]; ok {
			delete(s.force, n)
			s.emit(n.Kids[0])
		}
	case gChoice:
		if alt, ok := s.force[n]; ok {
			delete(s.force, n)
			s.emit(n.Kids[alt])
			return
		}
		best, bestLen := 0, 1<<30
		for i, k := range n.Kids {
			if m := s.g.nodeMin(k); m < bestLen {
				best, bestLen = i, m
			}
		}
		s.emit(n.Kids[best])
	}
}

// pathInto finds forced choices inside a production that lead to a node.
func pathInto(root, target *GNode) ([]forceStep, bool) {
	if root == target {
		return nil, true
	}
	for i, k := range root.Kids {
		if steps, ok := pathInto(k, target); ok {
			switch root.Kind {
			case gChoice:
				steps = append(steps, forceStep{root, i})
			case gOpt, gStar:
				steps = append(steps, forceStep{root, 1})
			}
			return steps, true
		}
	}
	return nil, false
}

type forceStep struct {
	n   *GNode
	alt int
}

// ntOccurrence finds a call of another production within a production.
func ntOccurrences(n *GNode, out *[]*GNode) {
	if n.Kind == gNT {
		*out = append(*out, n)
	}
	for _, k := range n.Kids {
		ntOccurrences(k, out)
	}
}

// chain finds the shortest production chain from start to prod, with the node to
// force in each.
func (g *Grammar) chain(start, prod string) []*GNode {
	type qi struct {
		name string
		path []*GNode
	}
	seen := map[string]bool{start: true}
	q := []qi{{start, nil}}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		if x.name == prod {
			return x.path
		}
		p := g.Prods[x.name]
		if p == nil {
			continue
		}
		var occ []*GNode
		ntOccurrences(p, &occ)
		for _, o := range occ {
			if !seen[o.Text] {
				seen[o.Text] = true
				q = append(q, qi{o.Text, append(append([]*GNode(nil), x.path...), o)})
			}
		}
	}
	return nil
}

// choicePoints lists every choice alternative, optional part and loop in a node.
func choicePoints(n *GNode, out *[]forceStep) {
	switch n.Kind {
	case gChoice:
		for i := range n.Kids {
			*out = append(*out, forceStep{n, i})
		}
	case gOpt, gStar, gPlus:
		*out = append(*out, forceStep{n, 1})
	}
	for _, k := range n.Kids {
		choicePoints(k, out)
	}
}

// Sentences makes one sentence per choice point of every production, from start.
func (g *Grammar) Sentences(start string, table map[string]string) ([]string, int) {
	uniq := map[string]bool{}
	var out []string
	points := 0
	for _, name := range g.Names {
		p := g.Prods[name]
		var pts []forceStep
		choicePoints(p, &pts)
		chain := g.chain(start, name)
		if chain == nil && name != start {
			// Reached only in lookahead, or not at all: try it from the production.
			chain = nil
		}
		for _, pt := range pts {
			points++
			force := map[*GNode]int{pt.n: pt.alt}
			// Force the way down through every production in the chain.
			cur := g.Prods[start]
			for _, occ := range chain {
				if steps, ok := pathInto(cur, occ); ok {
					for _, st := range steps {
						force[st.n] = st.alt
					}
				}
				cur = g.Prods[occ.Text]
			}
			if steps, ok := pathInto(p, pt.n); ok {
				for _, st := range steps {
					force[st.n] = st.alt
				}
			}
			s := &sentenceGen{g: g, table: table, force: force}
			s.emit(g.Prods[start])
			text := strings.Join(s.out, " ")
			if !uniq[text] {
				uniq[text] = true
				out = append(out, text)
			}
		}
	}
	return out, points
}

// Wrap puts a body sentence inside a module when it has no module of its own.
func WrapSentence(s string) string {
	if strings.Contains(s, "MODULE") {
		return strings.ReplaceAll(s, "====", "\n====") + "\n"
	}
	return "---- MODULE S ----\n" + s + "\n====\n"
}

// Mutations of a sentence's tokens: each one dropped, each one doubled, and each
// replaced by every token in others. They reach the syntax-error arms.
func MutateSentence(s string, others []string, emit func(string)) {
	toks := strings.Fields(s)
	for i := range toks {
		a := strings.Join(toks[:i], " ")
		b := strings.Join(toks[i+1:], " ")
		emit(a + " " + b)
		emit(a + " " + toks[i] + " " + toks[i] + " " + b)
		for _, o := range others {
			emit(a + " " + o + " " + b)
		}
	}
}

// dictRaw undoes JavaCC's string escapes.
func dictRaw(k string) string {
	var raw []byte
	for i := 0; i < len(k); i++ {
		if k[i] != '\\' || i+1 == len(k) {
			raw = append(raw, k[i])
			continue
		}
		i++
		switch k[i] {
		case 'n':
			raw = append(raw, '\n')
		case 't':
			raw = append(raw, '\t')
		case 'r':
			raw = append(raw, '\r')
		case 'f':
			raw = append(raw, '\f')
		case 'u':
			if i+4 < len(k) {
				if v, err := strconv.ParseUint(k[i+1:i+5], 16, 32); err == nil {
					raw = append(raw, string(rune(v))...)
					i += 4
					continue
				}
			}
			raw = append(raw, 'u')
		default:
			raw = append(raw, k[i])
		}
	}
	return string(raw)
}
