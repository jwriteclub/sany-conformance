// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package canon

import (
	"strings"
	"testing"
)

func mustCanon(t *testing.T, raw string) (string, GraphStats) {
	t.Helper()
	b, st, err := CanonicalXML([]byte(raw))
	if err != nil {
		t.Fatalf("CanonicalXML: %v", err)
	}
	return string(b), st
}

const graphA = `<?xml version="1.0" encoding="UTF-8" standalone="no"?><modules><RootModule>m</RootModule>` +
	`<context>` +
	`<entry><UID>7</UID><BuiltInKind><uniquename>$And</uniquename><level>0</level></BuiltInKind></entry>` +
	`<entry><UID>30</UID><UserDefinedOpKind><uniquename>x</uniquename><level>1</level><body><BuiltInKindRef><UID>7</UID></BuiltInKindRef></body></UserDefinedOpKind></entry>` +
	`<entry><UID>2</UID><ModuleNode><uniquename>m</uniquename><definitions><UserDefinedOpKindRef><UID>30</UID></UserDefinedOpKindRef></definitions></ModuleNode></entry>` +
	`</context><ModuleNodeRef><UID>2</UID></ModuleNodeRef></modules>`

// The same graph with every UID renumbered and the context reordered.
const graphARelabelled = `<modules><RootModule>m</RootModule>` +
	`<context>` +
	`<entry><UID>900</UID><ModuleNode><uniquename>m</uniquename><definitions><UserDefinedOpKindRef><UID>901</UID></UserDefinedOpKindRef></definitions></ModuleNode></entry>` +
	`<entry><UID>902</UID><BuiltInKind><uniquename>$And</uniquename><level>0</level></BuiltInKind></entry>` +
	`<entry><UID>901</UID><UserDefinedOpKind><uniquename>x</uniquename><level>1</level><body><BuiltInKindRef><UID>902</UID></BuiltInKindRef></body></UserDefinedOpKind></entry>` +
	`</context><ModuleNodeRef><UID>900</UID></ModuleNodeRef></modules>`

func TestCanonicalIgnoresUIDNumbersAndContextOrder(t *testing.T) {
	a, sa := mustCanon(t, graphA)
	b, _ := mustCanon(t, graphARelabelled)
	if a != b {
		t.Fatalf("relabelled graph differs:\n%s\n---\n%s", a, b)
	}
	if sa.Entries != 3 || sa.Unreachable != 0 || sa.Dangling != 0 {
		t.Fatalf("stats %+v", sa)
	}
	if strings.Contains(a, "30") || strings.Contains(a, "UID 7") {
		t.Fatalf("raw uid survived:\n%s", a)
	}
}

func TestCanonicalKeepsRealDifferences(t *testing.T) {
	base, _ := mustCanon(t, graphA)
	cases := map[string]struct{ from, to string }{
		"level":        {"<level>1</level>", "<level>2</level>"},
		"name":         {"<uniquename>x</uniquename>", "<uniquename>y</uniquename>"},
		"name space":   {"<uniquename>x</uniquename>", "<uniquename>x </uniquename>"},
		"root":         {"<RootModule>m</RootModule>", "<RootModule>n</RootModule>"},
		"ref target":   {"<body><BuiltInKindRef><UID>7</UID>", "<body><BuiltInKindRef><UID>30</UID>"},
		"ref kind":     {"<body><BuiltInKindRef><UID>7</UID></BuiltInKindRef>", "<body><UserDefinedOpKindRef><UID>7</UID></UserDefinedOpKindRef>"},
		"element name": {"<definitions>", "<definitionz>"},
	}
	for name, c := range cases {
		mutated := strings.Replace(graphA, c.from, c.to, 1)
		if name == "element name" {
			mutated = strings.Replace(mutated, "</definitions>", "</definitionz>", 1)
		}
		if mutated == graphA {
			t.Fatalf("%s: mutation did not apply", name)
		}
		got, _ := mustCanon(t, mutated)
		if got == base {
			t.Errorf("%s: real difference was normalised away", name)
		}
	}
}

func TestCanonicalKeepsSiblingOrder(t *testing.T) {
	a := `<modules><context/><A><B>1</B><B>2</B></A></modules>`
	b := `<modules><context/><A><B>2</B><B>1</B></A></modules>`
	ca, _ := mustCanon(t, a)
	cb, _ := mustCanon(t, b)
	if ca == cb {
		t.Fatal("sibling order was normalised away")
	}
}

func TestCanonicalCDATAEqualsEscapedText(t *testing.T) {
	a, _ := mustCanon(t, `<modules><context/><c><![CDATA[a<b & c]]></c></modules>`)
	b, _ := mustCanon(t, `<modules><context/><c>a&lt;b &amp; c</c></modules>`)
	if a != b {
		t.Fatalf("CDATA and escaped text differ:\n%s\n%s", a, b)
	}
}

func TestCanonicalLeafWhitespaceIsKept(t *testing.T) {
	a, _ := mustCanon(t, `<modules><context/><s>  </s></modules>`)
	b, _ := mustCanon(t, `<modules><context/><s></s></modules>`)
	if a == b {
		t.Fatal("whitespace in a leaf was removed")
	}
	if !strings.Contains(a, `s = "  "`) {
		t.Fatalf("leaf text lost: %s", a)
	}
}

func TestCanonicalIndentationIsSerialisationOnly(t *testing.T) {
	a, _ := mustCanon(t, `<modules><context/><a><b>1</b></a></modules>`)
	b, _ := mustCanon(t, "<modules>\n  <context/>\n  <a>\n    <b>1</b>\n  </a>\n</modules>\n")
	if a != b {
		t.Fatalf("pretty printing changed the result:\n%s\n%s", a, b)
	}
}

func TestCanonicalReportsUnreachableAndDangling(t *testing.T) {
	raw := `<modules><context>` +
		`<entry><UID>1</UID><A/></entry>` +
		`<entry><UID>2</UID><B><XRef><UID>99</UID></XRef></B></entry>` +
		`</context><Ref><UID>1</UID></Ref></modules>`
	c, st := mustCanon(t, raw)
	if st.Unreachable != 1 || st.Dangling != 1 {
		t.Fatalf("stats %+v", st)
	}
	if !strings.Contains(c, "B") {
		t.Fatalf("unreachable entry dropped:\n%s", c)
	}
}

func TestCanonicalRejectsDuplicateKeys(t *testing.T) {
	raw := `<modules><context><entry><UID>1</UID><A/></entry><entry><UID>1</UID><B/></entry></context></modules>`
	if _, _, err := CanonicalXML([]byte(raw)); err == nil {
		t.Fatal("duplicate UID accepted")
	}
}
