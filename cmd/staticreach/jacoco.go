// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// MethodCov holds a JaCoCo method's counters.
type MethodCov struct {
	Class, Name, Desc string
	Source            string // package/File.java
	FirstLine         int
	BranchTotal       int
	BranchHit         int
	InsnHit           int
}

// Key matches Method.Key.
func (m *MethodCov) Key() string { return m.Class + "." + m.Name + ":" + m.Desc }

// LineCov holds a JaCoCo source line's counters.
type LineCov struct {
	InsnMissed, InsnHit, BranchMissed, BranchHit int
}

// Coverage is one JaCoCo XML report.
type Coverage struct {
	Path    string
	Methods map[string]*MethodCov
	Keys    []string // document order
	Lines   map[string]map[int]LineCov
}

// ReadCoverage parses a JaCoCo XML report.
func ReadCoverage(path string) (*Coverage, error) {
	f, err := openData(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	cov := &Coverage{Path: path, Methods: map[string]*MethodCov{}, Lines: map[string]map[int]LineCov{}}
	d := xml.NewDecoder(f)
	d.Strict = true
	var pkg, class, source, sourceFile string
	var method *MethodCov
	for {
		tok, err := d.Token()
		if err == io.EOF {
			return cov, nil
		}
		if err != nil {
			return nil, fmt.Errorf("%s: %w", path, err)
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "package":
				pkg = attr(t, "name")
			case "class":
				class = attr(t, "name")
				source = pkg + "/" + attr(t, "sourcefilename")
			case "method":
				line, _ := strconv.Atoi(attr(t, "line"))
				method = &MethodCov{Class: class, Name: attr(t, "name"), Desc: attr(t, "desc"), Source: source, FirstLine: line}
				cov.Methods[method.Key()] = method
				cov.Keys = append(cov.Keys, method.Key())
			case "counter":
				if method == nil {
					continue
				}
				missed, _ := strconv.Atoi(attr(t, "missed"))
				covered, _ := strconv.Atoi(attr(t, "covered"))
				switch attr(t, "type") {
				case "BRANCH":
					method.BranchTotal = missed + covered
					method.BranchHit = covered
				case "INSTRUCTION":
					method.InsnHit = covered
				}
			case "sourcefile":
				sourceFile = pkg + "/" + attr(t, "name")
				cov.Lines[sourceFile] = map[int]LineCov{}
			case "line":
				nr, _ := strconv.Atoi(attr(t, "nr"))
				mi, _ := strconv.Atoi(attr(t, "mi"))
				ci, _ := strconv.Atoi(attr(t, "ci"))
				mb, _ := strconv.Atoi(attr(t, "mb"))
				cb, _ := strconv.Atoi(attr(t, "cb"))
				cov.Lines[sourceFile][nr] = LineCov{InsnMissed: mi, InsnHit: ci, BranchMissed: mb, BranchHit: cb}
			}
		case xml.EndElement:
			if t.Name.Local == "method" {
				method = nil
			}
		}
	}
}

func attr(t xml.StartElement, name string) string {
	for _, a := range t.Attr {
		if a.Name.Local == name {
			return a.Value
		}
	}
	return ""
}

// gzFile closes both layers.
type gzFile struct {
	*gzip.Reader
	f *os.File
}

func (g gzFile) Close() error { g.Reader.Close(); return g.f.Close() }

// openData opens a report, unpacking one that ends in .gz.
func openData(path string) (io.ReadCloser, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	if !strings.HasSuffix(path, ".gz") {
		return f, nil
	}
	z, err := gzip.NewReader(f)
	if err != nil {
		f.Close()
		return nil, err
	}
	return gzFile{z, f}, nil
}
