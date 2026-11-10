// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

// Command casemap answers which parts of SANY's grammar and code each case exercises.
// It reads data/casemap-index.gob.gz and needs no JVM, no network and no other
// repository. "Reached" here means JaCoCo saw the code run, not that any output was
// checked.
//
//	casemap case <id>       case to parts
//	casemap who <part>      part to cases: production, Production#n, Class.method, file:line, token:NAME
//	casemap sources         per set: cases, what it reaches, what only it reaches
//	casemap grammar         grammar elements: totals, and what could not be mapped
//	casemap select          recompute the selection and every case's why
//	casemap prove           show that the selection reaches what the whole suite reaches
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "casemap: "+format+"\n", args...)
	os.Exit(2)
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: casemap [-root dir] case|who|sources|grammar|select|prove ...")
	os.Exit(2)
}

func main() {
	root := flag.String("root", ".", "repository root")
	flag.Usage = usage
	flag.Parse()
	abs, err := filepath.Abs(*root)
	if err != nil {
		fatalf("%v", err)
	}
	p := Paths{Root: abs}
	if flag.NArg() == 0 {
		usage()
	}
	args := flag.Args()[1:]
	switch flag.Arg(0) {
	case "import":
		err = cmdImport(p, args)
	case "select":
		err = cmdSelect(p, args)
	case "variant-check":
		err = cmdVariantCheck(p, args)
	case "variant-apply":
		err = cmdVariantApply(p, args)
	case "prove":
		err = cmdProve(p, args)
	case "case":
		err = cmdCase(p, args)
	case "who":
		err = cmdWho(p, args)
	case "sources":
		err = cmdSources(p, args)
	case "grammar":
		err = cmdGrammar(p, args)
	default:
		usage()
	}
	if err != nil {
		fatalf("%v", err)
	}
}
