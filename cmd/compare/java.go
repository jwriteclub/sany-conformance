// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// RunJava runs stock SANY the way the manifests record it, in a private copy of the
// case directory (symlinks into libs/ followed), and returns its stdout text.
func RunJava(java, jar, caseDir, root, scratch string, timeout time.Duration) (Norm, error) {
	work, err := os.MkdirTemp(scratch, "java-")
	if err != nil {
		return Norm{}, err
	}
	defer os.RemoveAll(work)
	if out, err := exec.Command("cp", "-RL", caseDir+"/.", work).CombinedOutput(); err != nil {
		return Norm{}, fmt.Errorf("copy case: %v: %s", err, out)
	}
	tmp, _ := os.MkdirTemp(work, "tmp-")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, java, "-Djava.io.tmpdir="+tmp, "-cp", jar, "tla2sany.SANY", root)
	cmd.Dir = work
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	runErr := cmd.Run()
	if ctx.Err() != nil {
		return Norm{Class: ClassCrash, Note: "timeout"}, nil
	}
	if strings.Contains(stderr.String(), "Exception in thread") {
		return Norm{Class: ClassCrash, Note: "SANY threw (stderr)"}, nil
	}
	if _, isExit := runErr.(*exec.ExitError); runErr != nil && !isExit {
		return Norm{}, runErr // could not start java at all: a harness failure
	}
	if !strings.Contains(stdout.String(), "SANY") {
		// SANY always prints its banner. Without it Java did not run SANY (no JRE, bad jar):
		// that is the harness failing, never an "ok" answer.
		return Norm{}, fmt.Errorf("java did not run SANY on %s: %v: %s%s", root, runErr, stdout.String(), stderr.String())
	}
	return NormalizeSany(stdout.String()), nil
}
