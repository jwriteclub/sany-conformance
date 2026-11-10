// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bufio"
	"encoding/gob"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

// cmdImport regroups the index the measurement produced into this repository's
// sets and writes the committed copy. Only a maintainer with that index runs it; the
// queries read the committed copy and need nothing else.
func cmdImport(p Paths, args []string) error {
	fs := flag.NewFlagSet("import", flag.ExitOnError)
	from := fs.String("from", "", "index.gob as cmd/casemap build wrote it")
	dump := fs.String("dump", "", "also write every case as JSON lines, with its directory there")
	_ = fs.Parse(args)
	f, err := os.Open(*from)
	if err != nil {
		return err
	}
	defer f.Close()
	var ix Index
	if err := gob.NewDecoder(bufio.NewReaderSize(f, 1<<20)).Decode(&ix); err != nil {
		return err
	}
	ix.Groups = nil
	for _, g := range Groups {
		ix.Groups = append(ix.Groups, g.Name)
	}
	var out *json.Encoder
	if *dump != "" {
		d, err := os.Create(*dump)
		if err != nil {
			return err
		}
		defer d.Close()
		out = json.NewEncoder(d)
		out.SetEscapeHTML(false)
	}
	for i := range ix.Cases {
		c := &ix.Cases[i]
		c.Group = GroupOf(c.Source, c.Parent)
		if c.Group == "" {
			return fmt.Errorf("%s: source %q matches no set", c.ID, c.Source)
		}
		if out != nil {
			if err := out.Encode(c); err != nil {
				return err
			}
		}
		c.Dir = ""
	}
	if err := ix.recount(); err != nil {
		return err
	}
	return ix.Save(p.IndexFile())
}
