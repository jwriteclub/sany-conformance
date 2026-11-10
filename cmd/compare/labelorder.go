// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"regexp"
	"strings"
)

// ResultOrderNormalised is the outcome of a case that agrees in every respect
// except the order of the names in one or more LABEL_PARAMETER_UNNECESSARY (4332)
// extra-parameter lists (rule 8). It is never counted as a match.
const ResultOrderNormalised = "match (order-normalised)"

// PolicyLabelExtraOrder names the one comparison exception there is (rule 8).
const PolicyLabelExtraOrder = "label-extra-parameter-multiset (coord352)"

// reLabelExtra is the single 4332 template the exception covers, as the message reads
// after rule 4. SANY builds it in tla2sany/semantic/Generator.java, formalParamsEqual
// (pinned dc6a944, lines 7092-7099):
//
//	String res = "Label " + ln.getName().toString() + " declares extra parameter(s)  ";
//	while (iter.hasNext()) { ... res = res + nd.getName().toString() + "  "; }
//	errors.addMessage(ErrorCode.LABEL_PARAMETER_UNNECESSARY, ln.stn.getLocation(), "%s", res);
//
// iter walks the residual HashSet<FormalParamNode>, so the names come out in hash
// order; the loop only appends names to the string, and formalParamsEqual's return
// value is already false and is thrown away by its one caller (Generator.java:6933).
// Label and parameter names are TLA+ identifiers. The other 4332 message ("Illegal
// parameter %s of label `%s'.", Generator.java:6908-6914) does not match.
var reLabelExtra = regexp.MustCompile(`^(Label [A-Za-z0-9_]+ declares extra parameter\(s\) )([A-Za-z0-9_]+(?: [A-Za-z0-9_]+)*)$`)

// labelExtraParts splits a 4332 extra-parameter message into the fixed text before
// the list (which carries the label name) and the names, in printed order.
func labelExtraParts(msg string) (prefix string, names []string, ok bool) {
	m := reLabelExtra.FindStringSubmatch(msg)
	if m == nil {
		return "", nil, false
	}
	names = strings.Split(m[2], " ")
	return m[1], names, true
}

// LabelExtraOrderEqual reports whether two different message texts are the same
// 4332 extra-parameter message up to the order of its names: both match the template,
// the text before the list is identical (same label), each list names distinct
// parameters, and the two lists hold the same names with the same multiplicity, so
// neither is truncated, extended or duplicated. Everything else about the error
// (position, module, place in the list, the other errors) is checked by Compare, as
// for any error.
func LabelExtraOrderEqual(sany, pru string) bool {
	if sany == pru {
		return false // nothing to normalise; an exact match needs no exception
	}
	ps, ns, ok := labelExtraParts(sany)
	if !ok {
		return false
	}
	pp, np, ok := labelExtraParts(pru)
	if !ok || ps != pp || len(ns) != len(np) {
		return false
	}
	count := map[string]int{}
	for _, n := range ns {
		count[n]++
	}
	// Every prudence name must be a SANY name that occurs once (so a name SANY repeats
	// fails here), and no prudence name may repeat; with equal lengths that makes the
	// two lists the same set of distinct names.
	seen := map[string]bool{}
	for _, n := range np {
		if seen[n] || count[n] != 1 {
			return false
		}
		seen[n] = true
	}
	return true
}

// PolicyDiff is one error that matched only under rule 8, with both raw texts kept.
type PolicyDiff struct {
	Index          int    `json:"index"` // among the positioned errors, in compared order
	Module         string `json:"module,omitempty"`
	Line           int    `json:"line"`
	Col            int    `json:"col"`
	Sany           string `json:"sany"`
	Prudence       string `json:"prudence"`
	SanySHA256     string `json:"sany_sha256"`
	PrudenceSHA256 string `json:"prudence_sha256"`
}

// Policy records that a case's outcome needed rule 8, and where.
type Policy struct {
	Name  string       `json:"name"`
	Diffs []PolicyDiff `json:"diffs"`
}

func sha(s string) string {
	h := sha256.Sum256([]byte(s))
	return hex.EncodeToString(h[:])
}
