// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import "github.com/jwriteclub/prudence"

// unitOf reads prudence's Diagnostic.Unit (rule 3e), a direct copy of its module,
// line and column: nil when prudence reports none. Prudence sets it only for semantic
// and level diagnostics located in a monolith-extracted module (prudence 4b0c8ce,
// api.go DiagnosticUnit). A variable so that tests can stand in for prudence.
var unitOf = func(d prudence.Diagnostic) *reportedUnit {
	if d.Unit == nil {
		return nil
	}
	return &reportedUnit{Module: d.Unit.Module, Line: d.Unit.Line, Col: d.Unit.Column}
}
