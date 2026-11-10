// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

// The index's types. Field names are the on-disk format: the index is a gob.

// Case is one case of the suite, whatever its origin.
type Case struct {
	ID, Tier, Source, Origin, Commit string
	Group                            string
	RootFile                         string
	Files                            []string
	Aux                              []string
	Hash                             string
	Dir                              string // holds the files
	Size                             int
	Order                            int // keep order, over the whole suite
	InManifest                       bool

	Keep    string // keep reason, or upstream, seeded, not-kept
	Reasons []string
	Model   string
	Target  string
	Aim     string
	Parent  string // provenance parent, or the mutant's parent case
	Rule    string // mutant rule
	Intent  string // mutant intent, a hint only
	Finding string

	Expect   string // what it asserts, in words
	Upstream string // upstream status
	Result   string // upstream's expected accept or reject
	Stage    string

	Libs       []string
	Flags      []string
	Unresolved []string

	Outcome, L0, L1e, L1w string
}

// Part kinds.
const (
	PKMethod uint8 = iota
	PKLine
	PKArm
	PKElem
)

// Part is one thing a case can exercise. Every part belongs to one class, and a case
// has at most one probe array per class, so the cases reaching a part are a disjoint
// sum over that class's probe arrays.
type Part struct {
	Kind   uint8
	Class  int32
	Method int32 // method index in its class, or -1 for a line
	Line   int32
	Ord    int32 // arm: instruction ordinal
	Arm    int32
	Elem   int32
	// InPort says that the part is inside the port scope, ternary excluded.
	InPort bool
	// Ntk marks an arm of the jj_ntk ternary.
	Ntk bool
	// Reach is the number of cases reaching it; Groups a bitmask of their groups;
	// First the first case in keep order, -1 when nothing reaches it.
	Reach  int32
	Groups uint64
	First  int32
}

// ClassInfo is one class in the index.
type ClassInfo struct {
	Name      string
	File      string
	ID        uint64
	NProbes   int32
	Generated bool
	Methods   []MethodInfo
	// Images maps a probe to the parts it covers.
	Images [][]int32
	// Patterns lists this class's distinct probe arrays, by index in Index.Patterns.
	Patterns []int32
}

// MethodInfo is one method in the index.
type MethodInfo struct {
	Name, Desc string
	Part       int32
	InPort     bool
}

// Pattern is a distinct probe array of a class, with the cases that have it.
type Pattern struct {
	Class int32
	Bits  []byte
	Cases []int32
}

// ElemInfo is a mapped grammar element.
type ElemInfo struct {
	Element
	Part   int32
	InPort bool
	// Parts are the method or arms it stands for: any one reached means it ran.
	Parts []int32
}

// TokenStat is one token kind over the suite.
type TokenStat struct {
	Name   string // tokenImage
	Const  string // name in TLAplusParserConstants
	Reach  int32
	Groups uint64
	First  int32
}

// Index is everything the queries read.
type Index struct {
	Fingerprint string
	Groups      []string
	Cases       []Case
	CasePats    [][]int32
	CaseTokens  [][]uint16
	Classes     []ClassInfo
	Parts       []Part
	Patterns    []Pattern
	Elements    []ElemInfo
	Unmapped    []Unmapped
	Tokens      []TokenStat
	Productions []string
}

const (
	EKProduction   = "production"
	EKAlternative  = "alternative"
	EKOptional     = "optional"
	EKOptionalSkip = "optional-skip"
	EKLoopEnter    = "loop-enter"
	EKLoopExit     = "loop-exit"
	EKError        = "syntax-error"
	EKLookTrue     = "lookahead-true"
	EKLookFalse    = "lookahead-false"
)

// IsRule says that a kind counts as a grammar rule in the tables. A LOOKAHEAD
// decision's arms repeat an alternative's or an optional's, so they are listed
// apart and not counted twice.
// IsRule says that a kind counts as a grammar rule in the tables. A LOOKAHEAD
// decision's arms repeat an alternative's or an optional's, so they are listed
// apart and not counted twice.
func IsRule(kind string) bool { return kind != EKLookTrue && kind != EKLookFalse }

// Ref says where an element is in the generated code.
type Ref struct {
	Kind      string // method, switch-arm, switch-default, then, else
	Line, End int
	Labels    []string
	// BodyLine is the first statement of the arm, for the cross-check.
	BodyLine int
}

// Element is one grammar element with its place in the generated parser.
type Element struct {
	Prod string
	Kind string
	Path string // position inside the production, such as 2.1
	Text string
	Ref  Ref
}

// Unmapped is a grammar element with no exact place in the bytecode.
type Unmapped struct {
	Prod, Kind, Text, Reason string
}
