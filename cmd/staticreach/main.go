// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

// Command staticreach builds a candidate coverage map for SANY from the pinned
// class files and JaCoCo reports: every method that static reachability cannot
// rule out for a named set of entry points, joined to branch counters. No model,
// no network, no clock: the same input gives the same bytes.
//
//	staticreach            write the map and the report, then check soundness
//	staticreach check      only check soundness; exits 1 on an unexplained violation
//	staticreach view       print every scope with the jj_ntk ternary treated as
//	                       constant; a reporting view that writes nothing
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "staticreach: "+format+"\n", args...)
	os.Exit(2)
}

// Paths locates the inputs. Classes and sources come from the pinned tlaplus commit,
// the coverage reports from data/staticreach, which were measured by the suite.
type Paths struct {
	Root     string
	Tools    string // tlatools/org.lamport.tlatools of the pinned commit: src/
	ClassDir string // class files of the pinned build
	JDKHome  string
}

func (p Paths) abs(rel string) string { return filepath.Join(p.Root, rel) }

// Classes is the pinned build's class directory.
func (p Paths) Classes() string { return p.ClassDir }

// Sources is the pinned build's source directory.
func (p Paths) Sources() string { return filepath.Join(p.Tools, "src") }

// JDK is the JDK's home.
func (p Paths) JDK() string { return p.JDKHome }

// SpecsXML is the report for corpus plus kept tiers.
func (p Paths) SpecsXML() string { return p.abs("data/staticreach/specs.xml.gz") }

// UnionXML is the report for specs plus upstream's tests.
func (p Paths) UnionXML() string { return p.abs("data/staticreach/union.xml.gz") }

func main() {
	root := flag.String("root", ".", "repository root")
	scopes := flag.String("scopes", "", "scopes file; default is the embedded cmd/staticreach/scopes.conf")
	outDir := flag.String("out", ".work/staticreach", "directory for the map and the report; relative paths are under the root")
	classes := flag.String("classes", ".work/classes", "class files of the pinned build, unpacked from its jar")
	tools := flag.String("tools", ".work/external/tlaplus/tlatools/org.lamport.tlatools", "the pinned tlatools tree; only src/ is read")
	jdk := flag.String("jdk", os.Getenv("JAVA_HOME"), "JDK home, for the signatures of JDK classes")
	flag.Parse()
	abs, err := filepath.Abs(*root)
	if err != nil {
		fatalf("%v", err)
	}
	p := Paths{Root: abs, ClassDir: under(abs, *classes), Tools: under(abs, *tools), JDKHome: *jdk}
	cmd := "all"
	if flag.NArg() > 0 {
		cmd = flag.Arg(0)
	}
	cfg, err := loadConfig(*scopes)
	if err != nil {
		fatalf("%v", err)
	}
	switch cmd {
	case "check":
		a, err := NewAnalysis(p, cfg)
		if err != nil {
			fatalf("%v", err)
		}
		defer a.Close()
		rep := a.CheckSoundness()
		fmt.Print(rep.Text())
		if rep.Unexplained > 0 {
			fmt.Fprintf(os.Stderr, "SOUNDNESS FAILED: %d unexplained violations\n", rep.Unexplained)
			os.Exit(1)
		}
		fmt.Println("soundness: ok")
	case "view":
		a, err := NewAnalysis(p, cfg)
		if err != nil {
			fatalf("%v", err)
		}
		defer a.Close()
		a.NtkAll = true
		fmt.Print(a.viewText())
	case "all":
		a, err := NewAnalysis(p, cfg)
		if err != nil {
			fatalf("%v", err)
		}
		defer a.Close()
		rep := a.CheckSoundness()
		dir := *outDir
		if !filepath.IsAbs(dir) {
			dir = p.abs(dir)
		}
		if err := a.WriteOutputs(dir, rep); err != nil {
			fatalf("%v", err)
		}
		fmt.Print(rep.Text())
		if rep.Unexplained > 0 {
			fmt.Fprintf(os.Stderr, "SOUNDNESS FAILED: %d unexplained violations\n", rep.Unexplained)
			os.Exit(1)
		}
		fmt.Println("soundness: ok")
	default:
		fatalf("unknown command %q", cmd)
	}
}

func under(root, p string) string {
	if filepath.IsAbs(p) {
		return p
	}
	return filepath.Join(root, p)
}
