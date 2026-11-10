// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import "strings"

// Category is one of the five kinds of source.
type Category struct {
	Key, Title string
}

// Categories, in the order the table prints them. Key is the directory under sets/.
var Categories = []Category{
	{"tla-project", "From the TLA+ project"},
	{"other-projects", "From other projects"},
	{"ours-mechanical", "Ours, generated mechanically"},
	{"ours-model", "Ours, model-guided"},
	{"ours-inspection", "Ours, by inspection"},
}

// Group is one set of cases: a directory under sets/ and one under deep-cuts/.
type Group struct {
	Name     string
	Category string
	Title    string
	// Upstream sets are kept whole; their reason for being here is that they are upstream's.
	Upstream bool
	// Rules are matched against a case's source: a source, or a prefix ending in "/".
	Rules []string
}

// Groups is the grouping. The first rule to match wins.
var Groups = []*Group{
	{Name: "tlaplus-syntax-corpus", Category: "tla-project", Title: "Unified syntax corpus", Upstream: true, Rules: []string{"tlaplus/syntax-corpus"}},
	{Name: "tlaplus-semantic-corpus", Category: "tla-project", Title: "Semantic corpus", Upstream: true, Rules: []string{"tlaplus/semantic-corpus"}},
	{Name: "tlaplus-semantic-error-corpus", Category: "tla-project", Title: "Semantic error corpus", Upstream: true, Rules: []string{"tlaplus/semantic-error-corpus"}},
	{Name: "tlaplus-inline-java", Category: "tla-project", Title: "Specs written inline in Java tests", Upstream: true, Rules: []string{"tlaplus/inline-java"}},
	{Name: "tlaplus-test-models", Category: "tla-project", Title: "TLC test models", Upstream: true, Rules: []string{"tlaplus/test-model-fixture", "tlaplus/tlatools-test-model"}},
	{Name: "tlaplus-other", Category: "tla-project", Title: "Other modules in the TLA+ tools repository", Upstream: true, Rules: []string{"tlaplus/"}},
	{Name: "tlaplus-examples", Category: "tla-project", Title: "tlaplus/Examples", Upstream: true, Rules: []string{"Examples/"}},
	{Name: "tlaplus-communitymodules", Category: "tla-project", Title: "tlaplus/CommunityModules", Upstream: true, Rules: []string{"CommunityModules/"}},
	{Name: "tlaplus-rfcs", Category: "tla-project", Title: "tlaplus/rfcs language standard", Upstream: true, Rules: []string{"rfcs/"}},
	{Name: "tlapm", Category: "other-projects", Title: "tlapm, the TLA+ proof manager", Upstream: true, Rules: []string{"tlapm/"}},
	{Name: "tlx", Category: "other-projects", Title: "tlx", Upstream: true, Rules: []string{"tlx/"}},
	{Name: "tree-sitter-tlaplus", Category: "other-projects", Title: "tree-sitter-tlaplus", Upstream: true, Rules: []string{"tree-sitter-tlaplus/"}},
	{Name: "mutants", Category: "ours-mechanical", Title: "Seeded mutants of the corpus", Rules: []string{"mutant/"}},
	{Name: "fuzz-jazzer", Category: "ours-mechanical", Title: "Jazzer fuzz outputs", Rules: []string{"climb/jazzer"}},
	{Name: "grammar-sentences", Category: "ours-mechanical", Title: "Sentences derived from the grammar", Rules: []string{"grammar/"}},
	{Name: "model-one-shot", Category: "ours-model", Title: "One-shot model loop", Rules: []string{"climb/model"}},
	{Name: "model-agents", Category: "ours-model", Title: "Exploring agents", Rules: []string{"climb/agent"}},
	{Name: "model-scoring", Category: "ours-model", Title: "Scoring specs", Rules: []string{"climb/score"}},
	{Name: "model-ternary-hunt", Category: "ours-model", Title: "Specs hunting a stale jj_ntk", Rules: []string{"incoming/ternary", "climb/model/ternary"}},
	{Name: "model-openstart", Category: "ours-model", Title: "OpenStart scan modules", Rules: []string{"incoming/openstart", "climb/model/astra-openstart"}},
	{Name: "model-pit-survivors", Category: "ours-model", Title: "Model specs aimed at surviving mutants", Rules: []string{"climb/pit/model"}},
	{Name: "inspection", Category: "ours-inspection", Title: "Written by reading the code", Rules: []string{"climb/hand", "climb/grammar", "climb/inspection", "climb/inspection/", "review/"}},
	{Name: "inspection-pit-survivors", Category: "ours-inspection", Title: "Written by reading a surviving mutant's code", Rules: []string{"climb/pit/review"}},
}

// GroupOf picks a case's set by its source and, for a reviewer's spec that passed
// through the fuzz directory, by its parent. The empty string means no rule matched.
func GroupOf(source, parent string) string {
	if source == "climb/jazzer" && strings.HasPrefix(parent, "hand/") {
		return "inspection"
	}
	// Longer, more specific sources first: climb/model/ternary before climb/model.
	for _, name := range []string{"model-openstart", "model-ternary-hunt"} {
		for _, r := range groupByName(name).Rules {
			if r == source {
				return name
			}
		}
	}
	for _, g := range Groups {
		for _, r := range g.Rules {
			if r == source || (strings.HasSuffix(r, "/") && strings.HasPrefix(source, r)) {
				return g.Name
			}
		}
	}
	return ""
}

// groupByName finds a set.
func groupByName(name string) *Group {
	for _, g := range Groups {
		if g.Name == name {
			return g
		}
	}
	return nil
}
