// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

// File is one file of a case; Path is relative to the case directory.
type File struct {
	Path   string `json:"path"`
	SHA256 string `json:"sha256"`
}

// Expect is the expectation recorded upstream for a tier A case.
type Expect struct {
	Stage    string   `json:"stage"`              // syntax | semantic
	Result   string   `json:"result"`             // accept | reject
	Code     int      `json:"code,omitempty"`     // SANY error/warning code (semantic error corpus)
	Severity string   `json:"severity,omitempty"` // E | W
	Messages []string `json:"messages,omitempty"` // substrings that must appear in SANY output
	Basis    string   `json:"basis"`              // how upstream records it
	Upstream string   `json:"upstream,omitempty"` // ok | known_failure | skipped | ignored
	Note     string   `json:"note,omitempty"`
	Asserts  int      `json:"assertions,omitempty"` // RefersTo/IsLevel assertions embedded in the module
}

// Mutant describes how a tier C case was derived.
type Mutant struct {
	Parent string `json:"parent"`
	Rule   string `json:"rule"`
	Seed   int64  `json:"seed"`
	Detail string `json:"detail"`
	Intent string `json:"intent"` // hint only, never an assertion
}

// Dep is a module a case needs that is not copied into the case directory.
type Dep struct {
	Name string `json:"name"`
	Lib  string `json:"lib"` // std | CommunityModules | tlapm-library
}

// Case is one manifest record.
type Case struct {
	ID         string   `json:"id"`
	Tier       string   `json:"tier"`
	Source     string   `json:"source"`
	Commit     string   `json:"commit"`
	Origin     string   `json:"origin"` // path (and test name) inside the source repo
	Name       string   `json:"name,omitempty"`
	Root       string   `json:"root_module"`
	RootFile   string   `json:"root_file"`
	Files      []File   `json:"files"`
	Aux        []File   `json:"aux,omitempty"` // not fed to SANY (expected.sexp)
	Hash       string   `json:"hash"`
	RootHash   string   `json:"root_hash"`
	Expect     *Expect  `json:"expect,omitempty"`
	ExtDeps    []Dep    `json:"external_deps,omitempty"`
	Unresolved []string `json:"unresolved,omitempty"`
	Flags      []string `json:"flags,omitempty"`
	DupOf      string   `json:"dup_of,omitempty"`
	Mut        *Mutant  `json:"mutant,omitempty"`

	content map[string][]byte // path -> bytes, written to cases/<id>/
	aux     map[string][]byte
}

func sum(b []byte) string {
	h := sha256.Sum256(b)
	return hex.EncodeToString(h[:])
}

// finalize computes file hashes and the case hash. The case hash covers the
// sorted (path, bytes) set of Files, so a renamed copy of the same closure
// collides only when paths match too.
func (c *Case) finalize() {
	c.Files = c.Files[:0]
	for p, b := range c.content {
		c.Files = append(c.Files, File{Path: p, SHA256: sum(b)})
	}
	sort.Slice(c.Files, func(i, j int) bool { return c.Files[i].Path < c.Files[j].Path })
	var sb strings.Builder
	for _, f := range c.Files {
		fmt.Fprintf(&sb, "%s\x00%s\n", f.Path, f.SHA256)
	}
	c.Hash = sum([]byte(sb.String()))
	c.RootHash = sum(c.content[c.RootFile])
	c.Aux = nil
	for p, b := range c.aux {
		c.Aux = append(c.Aux, File{Path: p, SHA256: sum(b)})
	}
	sort.Slice(c.Aux, func(i, j int) bool { return c.Aux[i].Path < c.Aux[j].Path })
	sort.Strings(c.Flags)
}
