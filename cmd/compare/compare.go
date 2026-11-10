// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import "fmt"

// Compare decides one case. The rule, in full:
//
//	match       same class, same number of errors, and for each error the same line,
//	            column and (whitespace-collapsed) message text
//	incomplete  prudence says it has not finished (class "incomplete") and SANY got
//	            past parsing, so prudence has not been shown wrong, only silent
//	mismatch    anything else, including prudence being incomplete on a case SANY
//	            rejects at parsing (syntax is implemented, so that is a miss);
//	            (SANY's "missing-module" is not syntax, so it counts as incomplete.)
//	            Kind says how: class, position (same class, different places) or
//	            text (same places, different words)
//	oracle-unavailable
//	            SANY itself crashed (recorded or live), so there is no verdict to compare
//	            with; never a match, whatever prudence said
//	timeout     the per-case deadline passed before prudence answered (its own outcome,
//	            never a match, whatever SANY said)
//	error       prudence panicked or errored where SANY did not crash
//	match (order-normalised)
//	            would be a match, except that one or more errors differ only in the
//	            order of the names in a 4332 extra-parameter list (rule 8,
//	            LabelExtraOrderEqual); every other class, count, position, module and
//	            text agrees. Counted on its own, never as a match. The returned Policy
//	            names the rule and keeps both raw texts of each such error.
func Compare(sany, pru Norm) (result, kind string, policy *Policy) {
	switch {
	case sany.Class == ClassCrash:
		return "oracle-unavailable", "", nil // SANY gave no verdict; nothing to match against
	case pru.Class == ClassTimeout:
		return "timeout", "", nil
	case pru.Class == ClassCrash && sany.Class != ClassCrash:
		return "error", "prudence", nil
	case pru.Class == ClassIncomp && sany.Class == ClassParse:
		return "mismatch", "class", nil
	case pru.Class == ClassIncomp:
		return "incomplete", "", nil
	case sany.Class != pru.Class:
		return "mismatch", "class", nil
	}
	// SANY errors with no position (missing module, cycle, 4003, 4221, BUILTINS, the
	// never-set sentinel) are matched to a prudence error with the same text, wherever
	// prudence puts it (rule 5); the rest are compared in order.
	var withPos []Diag
	left := append([]Diag(nil), pru.Diags...)
	unmatched := 0
	textDiff := false
	for _, a := range sany.Diags {
		if a.Line != 0 {
			withPos = append(withPos, a)
			continue
		}
		found := false
		for i, b := range left {
			if b.Msg == a.Msg {
				left = append(left[:i], left[i+1:]...)
				found = true
				break
			}
		}
		if !found {
			unmatched++
		}
	}
	// A no-position SANY error with no same-text counterpart is paired with a prudence
	// error that no positioned SANY error claims, so a wording difference is a text
	// mismatch, not a count mismatch.
	for ; unmatched > 0; unmatched-- {
		paired := false
		for i, b := range left {
			claimed := false
			for _, a := range withPos {
				if a.Line == b.Line && a.Col == b.Col {
					claimed = true
					break
				}
			}
			if !claimed {
				left = append(left[:i], left[i+1:]...)
				paired = true
				break
			}
		}
		if !paired {
			break
		}
		textDiff = true
	}
	if unmatched > 0 || len(withPos) != len(left) {
		return "mismatch", "position", nil
	}
	text := textDiff
	var diffs []PolicyDiff
	for i := range withPos {
		a, b := withPos[i], left[i]
		if a.Line != b.Line || a.Col != b.Col {
			return "mismatch", "position", nil
		}
		if a.Module != "" && b.Module != "" && a.Module != b.Module {
			return "mismatch", "position", nil // same place, different module
		}
		if a.Msg != b.Msg {
			if LabelExtraOrderEqual(a.Msg, b.Msg) {
				diffs = append(diffs, PolicyDiff{Index: i, Module: a.Module, Line: a.Line, Col: a.Col,
					Sany: a.Msg, Prudence: b.Msg, SanySHA256: sha(a.Msg), PrudenceSHA256: sha(b.Msg)})
			} else {
				text = true
			}
		}
	}
	if text {
		return "mismatch", "text", nil
	}
	if len(diffs) > 0 {
		return ResultOrderNormalised, "", &Policy{Name: PolicyLabelExtraOrder, Diffs: diffs}
	}
	return "match", "", nil
}

// Show renders a Norm on a few lines for the report.
func Show(n Norm) string {
	s := n.Class
	if n.Note != "" {
		s += " (" + n.Note + ")"
	}
	if n.Completion != "" {
		s += " [completion " + n.Completion + ", level " + n.Level + "]"
	}
	for _, d := range n.Diags {
		if d.Unit {
			s += fmt.Sprintf("\n%s %d:%d (unit)  %s", d.Module, d.Line, d.Col, d.Msg)
		} else if d.Module != "" {
			s += fmt.Sprintf("\n%s %d:%d  %s", d.Module, d.Line, d.Col, d.Msg)
		} else {
			s += fmt.Sprintf("\n%d:%d  %s", d.Line, d.Col, d.Msg)
		}
	}
	return s
}
