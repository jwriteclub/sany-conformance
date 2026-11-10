// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/jwriteclub/prudence"
)

// AcquisitionOverride is one recorded per-case value for prudence's module-acquisition
// visit budget (Options.MaxAcquisitionVisits). The case is scored normally (match,
// mismatch, ...); the value is written in its row and counted in the report.
type AcquisitionOverride struct {
	Set    string // set as the report names it (manifest directory name)
	ID     string
	Root   string // the manifest's root_file, checked so a stale entry fails loudly
	Visits int    // Options.MaxAcquisitionVisits for this case
	Reason string
}

// acquisitionOverrides is the whole table. Add a case only with a recorded reason;
// there is no other per-case setting.
var acquisitionOverrides = []AcquisitionOverride{
	{Set: "tlaplus-test-models", ID: "B-test-model-bug156-te-9a905e", Root: "TE.tla", Visits: 3000000,
		Reason: "TE.tla's module acquisition needs more than the default 1,000,000 visits per pass (root reviewer)"},
}

// AcquisitionRecord is what a row carries for a case scored under an override.
// Applied is true when the value was passed to Check as Options.MaxAcquisitionVisits,
// which a live run always does (prudence pinned in prudence.pin); it is kept in the row
// so that a recorded result says so explicitly.
type AcquisitionRecord struct {
	MaxAcquisitionVisits int    `json:"max_acquisition_visits"`
	Applied              bool   `json:"applied"`
	Reason               string `json:"reason"`
}

// applyAcquisitionOverrides validates the table and sets each selected case's record.
// An entry whose case is not selected (a -set, -ids or -limit run) is skipped; an entry
// whose case is selected under a different root file is an error.
func applyAcquisitionOverrides(cases []*Case, table []AcquisitionOverride) error {
	byKey := map[string]AcquisitionOverride{}
	for _, o := range table {
		k := o.Set + "/" + o.ID
		if o.Set == "" || o.ID == "" || o.Root == "" || o.Reason == "" || o.Visits <= 0 {
			return fmt.Errorf("acquisition override %s: set, id, root, reason and a positive value are required", k)
		}
		if _, dup := byKey[k]; dup {
			return fmt.Errorf("acquisition override %s is listed twice", k)
		}
		byKey[k] = o
	}
	for _, c := range cases {
		o, ok := byKey[c.Set+"/"+c.ID]
		if !ok {
			continue
		}
		if c.Root != o.Root {
			return fmt.Errorf("acquisition override %s/%s names root %s, the manifest has %s", o.Set, o.ID, o.Root, c.Root)
		}
		c.Acquisition = &AcquisitionRecord{MaxAcquisitionVisits: o.Visits, Applied: true, Reason: o.Reason}
	}
	return nil
}

// acquisitionVisits is the value RunPrudence passes for a case: 0 (prudence's default)
// unless the case has an override.
func acquisitionVisits(c *Case) int {
	if c.Acquisition == nil {
		return 0
	}
	return c.Acquisition.MaxAcquisitionVisits
}

// setAcquisitionVisits passes a case's value to Check; 0 keeps prudence's default
// (DefaultMaxAcquisitionVisits).
func setAcquisitionVisits(req *prudence.Request, visits int) {
	req.Options.MaxAcquisitionVisits = visits
}

// acquisitionFootnote lists the cases scored under an override, for the report.
func acquisitionFootnote(cases []*Case) string {
	var lines []string
	for _, c := range cases {
		if a := c.Acquisition; a != nil {
			state := "applied"
			if !a.Applied {
				state = "recorded as NOT applied: Check ran with its default budget"
			}
			lines = append(lines, fmt.Sprintf("- `%s/%s`: `Options.MaxAcquisitionVisits` = %d (%s; scored %s). %s", c.Set, c.ID, a.MaxAcquisitionVisits, state, c.Result, a.Reason))
		}
	}
	sort.Strings(lines)
	s := fmt.Sprintf("**Acquisition-visit overrides**: %d case(s) scored under a recorded per-case `Options.MaxAcquisitionVisits` (not an outcome; each is counted in its normal column above).", len(lines))
	if len(lines) == 0 {
		return s + "\n\n"
	}
	return s + "\n\n" + strings.Join(lines, "\n") + "\n\n"
}
