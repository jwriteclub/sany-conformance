// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const extra = "Label L declares extra parameter(s) "

func sem(diags ...Diag) Norm { return Norm{Class: ClassSemantic, Diags: diags} }

func at(line, col int, msg string) Diag { return Diag{Module: "M", Line: line, Col: col, Msg: msg} }

// base is a typical SANY answer around a 4332 extra-parameter error.
func base(list string) Norm {
	return sem(at(3, 1, "Unknown operator: `x'."), at(5, 9, extra+list), at(7, 2, "Operator Foo already defined or declared."))
}

func check(t *testing.T, name string, sany, pru Norm, wantResult, wantKind string) *Policy {
	t.Helper()
	r, k, p := Compare(sany, pru)
	if r != wantResult || k != wantKind {
		t.Errorf("%s: got (%q, %q), want (%q, %q)", name, r, k, wantResult, wantKind)
	}
	if (p != nil) != (wantResult == ResultOrderNormalised) {
		t.Errorf("%s: policy %v with result %q", name, p, r)
	}
	return p
}

func TestLabelExtraExactMatch(t *testing.T) {
	check(t, "identical", base("a b c"), base("a b c"), "match", "")
	if LabelExtraOrderEqual(extra+"a b", extra+"a b") {
		t.Error("identical texts need no exception")
	}
}

func TestLabelExtraPermittedReorder(t *testing.T) {
	for _, list := range []string{"c a b", "b c a", "c b a"} {
		p := check(t, "reorder "+list, base("a b c"), base(list), ResultOrderNormalised, "")
		if p == nil {
			continue
		}
		if p.Name != PolicyLabelExtraOrder || len(p.Diffs) != 1 {
			t.Fatalf("policy %+v", p)
		}
		d := p.Diffs[0]
		if d.Index != 1 || d.Line != 5 || d.Col != 9 || d.Module != "M" || d.Sany != extra+"a b c" || d.Prudence != extra+list {
			t.Errorf("diff %+v", d)
		}
		if d.SanySHA256 != sha(extra+"a b c") || d.PrudenceSHA256 != sha(extra+list) || d.SanySHA256 == d.PrudenceSHA256 {
			t.Errorf("hashes %+v", d)
		}
	}
	// Two reordered lists in one case, each in its own error.
	two := func(x, y string) Norm {
		return sem(at(5, 9, extra+x), at(9, 4, "Label K declares extra parameter(s) "+y))
	}
	if p := check(t, "two lists", two("a b", "p q r"), two("b a", "r p q"), ResultOrderNormalised, ""); p != nil && len(p.Diffs) != 2 {
		t.Errorf("want 2 diffs, got %+v", p.Diffs)
	}
}

func TestLabelExtraNegativeText(t *testing.T) {
	cases := []struct{ name, sany, pru string }{
		{"missing name", extra + "a b c", extra + "b a"},
		{"truncated to first", extra + "a b c", extra + "a"},
		{"extra name", extra + "a b", extra + "b a c"},
		{"duplicated name, count changes", extra + "a b", extra + "b a a"},
		{"duplicate replaces a name", extra + "a b", extra + "a a"},
		{"duplicated names both sides, same multiset", extra + "a a b", extra + "b a a"},
		{"duplicated on SANY side", extra + "a b a", extra + "b a"},
		{"different set", extra + "a b", extra + "c a"},
		{"different name spelling", extra + "a b", extra + "B a"},
		{"changed label", extra + "a b", "Label K declares extra parameter(s) b a"},
		{"changed prefix word", extra + "a b", "Label L declares extra parameters(s) b a"},
		{"changed prefix punctuation", extra + "a b", "Label L declares extra parameter(s): b a"},
		{"changed suffix", extra + "a b", extra + "b a."},
		{"suffix text after list", extra + "a b", extra + "b a in label"},
		{"spacing outside the list", extra + "a b", "Label L declares extra parameter(s)b a"},
		{"spacing inside the prefix", extra + "a b", "Label L declares extra parameter (s) b a"},
		{"names glued", extra + "a b", extra + "ba"},
		{"other 4332 text", "Illegal parameter a b of label `L'.", "Illegal parameter b a of label `L'."},
		{"other label message", "Label L must contain formal parameter `a b'.", "Label L must contain formal parameter `b a'."},
		{"empty list", extra + "a", "Label L declares extra parameter(s)"},
		{"changed label, longer name", extra + "a b", "Label Lx declares extra parameter(s) b a"},
		{"changed label, case", extra + "a b", "Label l declares extra parameter(s) b a"},
		{"leading text", extra + "a b", "x" + extra + "b a"},
		// The template must be exact on both sides: a shared deviation is not covered.
		{"both sides missing the space", "Label L declares extra parameter(s)a b", "Label L declares extra parameter(s)b a"},
		{"both sides other wording", "Label L declares extra parameters a b", "Label L declares extra parameters b a"},
		{"both sides trailing period", extra + "a b.", extra + "b a."},
		{"both sides non-identifier names", extra + "a' b", extra + "b a'"},
	}
	for _, c := range cases {
		if LabelExtraOrderEqual(c.sany, c.pru) {
			t.Errorf("%s: LabelExtraOrderEqual accepted %q / %q", c.name, c.sany, c.pru)
		}
		check(t, c.name, sem(at(5, 9, c.sany)), sem(at(5, 9, c.pru)), "mismatch", "text")
	}
}

func TestLabelExtraNegativeStructure(t *testing.T) {
	// The reordered list is equal under rule 8, so each failure below is caused by the
	// other difference alone, and is the ordinary outcome for that difference.
	sany := base("a b c")
	moved := func(line, col int, module string) Norm {
		n := base("c b a")
		n.Diags[1].Line, n.Diags[1].Col, n.Diags[1].Module = line, col, module
		return n
	}
	check(t, "changed column", sany, moved(5, 10, "M"), "mismatch", "position")
	check(t, "changed line", sany, moved(6, 9, "M"), "mismatch", "position")
	check(t, "changed module", sany, moved(5, 9, "N"), "mismatch", "position")

	swapped := base("c b a")
	swapped.Diags[0], swapped.Diags[1] = swapped.Diags[1], swapped.Diags[0]
	check(t, "error order", sany, swapped, "mismatch", "position")

	fewer := base("c b a")
	fewer.Diags = fewer.Diags[1:]
	check(t, "error count, one fewer", sany, fewer, "mismatch", "position")
	more := base("c b a")
	more.Diags = append(more.Diags, at(8, 1, extra+"a b c"))
	check(t, "error count, one more", sany, more, "mismatch", "position")

	other := base("c b a")
	other.Diags[2].Msg = "Operator Bar already defined or declared."
	check(t, "other error text", sany, other, "mismatch", "text")

	class := base("c b a")
	class.Class = ClassParse
	check(t, "class", sany, class, "mismatch", "class")
	check(t, "prudence incomplete", sany, Norm{Class: ClassIncomp}, "incomplete", "")

	// An error with no position is matched by exact text only (rule 5): no exception.
	nopos := func(list string) Norm { return sem(Diag{Msg: extra + list}) }
	check(t, "no-position error", nopos("a b"), nopos("b a"), "mismatch", "text")
}

func TestLabelExtraRecordAndReport(t *testing.T) {
	norm := &Case{Set: "s", ID: "norm", Sany: base("a b"), Prudence: base("b a")}
	plain := &Case{Set: "s", ID: "plain", Sany: base("a b"), Prudence: base("a b")}
	diff := &Case{Set: "s", ID: "diff", Sany: base("a b"), Prudence: base("a c")}
	for _, c := range []*Case{norm, plain, diff} {
		c.Result, c.Kind, c.Policy = Compare(c.Sany, c.Prudence)
	}
	for _, c := range []*Case{plain, diff} {
		b, _ := json.Marshal(caseRecord(c))
		if strings.Contains(string(b), "policy") {
			t.Errorf("%s: record carries policy: %s", c.ID, b)
		}
	}
	b, _ := json.Marshal(caseRecord(norm))
	var rec struct {
		Result   string  `json:"result"`
		Policy   *Policy `json:"policy"`
		Sany     Norm    `json:"sany"`
		Prudence Norm    `json:"prudence"`
	}
	if err := json.Unmarshal(b, &rec); err != nil {
		t.Fatal(err)
	}
	if rec.Result != ResultOrderNormalised || rec.Policy == nil || rec.Policy.Diffs[0].Prudence != extra+"b a" {
		t.Errorf("record %s", b)
	}
	if rec.Prudence.Diags[1].Msg != extra+"b a" || rec.Sany.Diags[1].Msg != extra+"a b" {
		t.Errorf("raw sides not kept: %s", b)
	}

	out := t.TempDir()
	if err := writeOutputs(out, []*Case{norm, plain, diff}, []string{"s"}, false, "", ""); err != nil {
		t.Fatal(err)
	}
	report, err := os.ReadFile(filepath.Join(out, "report.md"))
	if err != nil {
		t.Fatal(err)
	}
	r := string(report)
	// cases, match, match (order-normalised), mismatch, ...
	if !strings.Contains(r, "| **all** | 3 | 1 | 1 | 1 | 0 |") {
		t.Errorf("summary row wrong:\n%s", r)
	}
	if !strings.Contains(r, "#### `s/norm`") || !strings.Contains(r, extra+"b a") {
		t.Errorf("order-normalised case not listed")
	}
}
