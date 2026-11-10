// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bufio"
	"compress/gzip"
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// Paths locates this tool's data under the repository root.
type Paths struct {
	Root string
}

func (p Paths) abs(rel string) string { return filepath.Join(p.Root, rel) }

// IndexFile is the committed, compressed index.
func (p Paths) IndexFile() string { return p.abs("data/casemap-index.gob.gz") }

// Save writes the index. gzip's header carries no time here, so equal input gives equal bytes.
func (ix *Index) Save(path string) error {
	tmp := path + ".tmp"
	f, err := os.Create(tmp)
	if err != nil {
		return err
	}
	z, err := gzip.NewWriterLevel(f, gzip.BestCompression)
	if err != nil {
		f.Close()
		return err
	}
	w := bufio.NewWriterSize(z, 1<<20)
	if err := gob.NewEncoder(w).Encode(ix); err != nil {
		f.Close()
		return err
	}
	if err := w.Flush(); err != nil {
		f.Close()
		return err
	}
	if err := z.Close(); err != nil {
		f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}

// LoadIndex reads the committed index.
func LoadIndex(p Paths) (*Index, error) {
	f, err := os.Open(p.IndexFile())
	if err != nil {
		return nil, err
	}
	defer f.Close()
	z, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	var ix Index
	if err := gob.NewDecoder(bufio.NewReaderSize(z, 1<<20)).Decode(&ix); err != nil {
		return nil, err
	}
	return &ix, nil
}

func hasBit(b []byte, i int) bool { return b[i/8]&(1<<(i%8)) != 0 }

func clip(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}

// patternParts is the union of the probes' images.
func (ix *Index) patternParts(pat *Pattern) []int32 {
	cls := &ix.Classes[pat.Class]
	seen := map[int32]bool{}
	var out []int32
	for i := 0; i < int(cls.NProbes); i++ {
		if !hasBit(pat.Bits, i) {
			continue
		}
		for _, part := range cls.Images[i] {
			if !seen[part] {
				seen[part] = true
				out = append(out, part)
			}
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
	return out
}

// recount recomputes each part's reach and groups from the probe arrays, and each
// token's, after cases changed set.
func (ix *Index) recount() error {
	groupBit := map[string]uint64{}
	for i, g := range ix.Groups {
		groupBit[g] = 1 << i
	}
	if len(ix.Groups) > 64 {
		return fmt.Errorf("more than 64 sets")
	}
	for i := range ix.Parts {
		ix.Parts[i].Reach, ix.Parts[i].Groups, ix.Parts[i].First = 0, 0, -1
	}
	for pn := range ix.Patterns {
		pat := &ix.Patterns[pn]
		var mask uint64
		first := int32(-1)
		for _, ci := range pat.Cases {
			mask |= groupBit[ix.Cases[ci].Group]
			if first < 0 || ci < first {
				first = ci
			}
		}
		n := int32(len(pat.Cases))
		for _, part := range ix.patternParts(pat) {
			pt := &ix.Parts[part]
			pt.Reach += n
			pt.Groups |= mask
			if pt.First < 0 || first < pt.First {
				pt.First = first
			}
		}
	}
	for k := range ix.Tokens {
		ix.Tokens[k].Reach, ix.Tokens[k].Groups, ix.Tokens[k].First = 0, 0, -1
	}
	for i := range ix.Cases {
		bit := groupBit[ix.Cases[i].Group]
		for _, k := range ix.CaseTokens[i] {
			t := &ix.Tokens[k]
			t.Reach++
			t.Groups |= bit
			if t.First < 0 {
				t.First = int32(i)
			}
		}
	}
	return nil
}
