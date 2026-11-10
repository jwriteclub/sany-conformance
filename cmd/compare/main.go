// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

// Command compare runs every case through SANY's answer and through prudence and
// reports where they agree. See scripts/compare.sh and docs/compare.md, "Compare an
// implementation". Exit status is nonzero only if the harness itself fails.
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

// Case is what we need from a manifest row.
type Case struct {
	Set      string
	ID       string
	Dir      string
	Root     string
	Status   string // recorded: ok or crash
	Stdout   string // recorded SANY stdout
	Verdict  string
	Sany     Norm
	Prudence Norm
	Result   string // match, match (order-normalised), mismatch, incomplete, timeout, oracle-unavailable, error
	Kind     string // for a mismatch: class, position, text
	// Policy is set only when Result is ResultOrderNormalised: the rule that was
	// needed and the raw texts it was needed for.
	Policy *Policy
	// Acquisition is set only for a case in the acquisitionOverrides table (or, in
	// -rescore, a recorded row that carries one). Not an outcome.
	Acquisition *AcquisitionRecord
	// FalseComplete: prudence says Complete with no error where SANY reports one.
	// Independent of Result.
	FalseComplete bool
}

type manifestRow struct {
	ID       string `json:"id"`
	RootFile string `json:"root_file"`
	Expected struct {
		Status string `json:"status"`
		Stdout string `json:"stdout"`
	} `json:"expected"`
}

func main() {
	repo := flag.String("repo", ".", "repository root")
	set := flag.String("set", "", "only this set (directory name, e.g. tlaplus-syntax-corpus)")
	limit := flag.Int("limit", 0, "at most N cases per set, in manifest order (0 = all)")
	jobs := flag.Int("jobs", 1, "cases in parallel")
	out := flag.String("out", "compare-out", "output directory")
	useJava := flag.Bool("java", false, "rerun stock SANY instead of using the recorded output")
	jar := flag.String("jar", "", "tla2tools.jar for -java")
	java := flag.String("javabin", "java", "java binary for -java")
	stage := flag.String("stage", "", `prudence stage: "" (default: semantic and level, as the public Check does), "semantic", or "syntax"`)
	idsFile := flag.String("ids", "", "file with one case id per line: run only these")
	deep := flag.Bool("deep-cuts", false, "also include deep-cuts/")
	timeout := flag.Duration("timeout", 300*time.Second, "per-case timeout, each side")
	flag.BoolVar(&recordInfo, "info", false, "keep prudence's non-error diagnostics (the incomplete reason) in results.jsonl")
	flag.BoolVar(&no3d, "no-3d", false, "turn normalisation rule 3d off (with -rescore: undo it in the given results)")
	rescore := flag.String("rescore", "", "score this results.jsonl again (no prudence or SANY run); -repo gives the case files")
	flag.BoolVar(&sanyFromManifest, "sany-from-manifest", false, "with -rescore: derive the SANY side again from the manifests' recorded output (a recorded-SANY results file only)")
	flag.Parse()

	if sanyFromManifest && *rescore == "" {
		fatal("-sany-from-manifest needs -rescore")
	}
	if *rescore != "" {
		if err := os.MkdirAll(*out, 0o755); err != nil {
			fatal(err.Error())
		}
		if err := rescoreResults(*repo, *rescore, *out, *stage, *deep); err != nil {
			fatal(err.Error())
		}
		return
	}

	if *useJava && *jar == "" {
		fatal("-java needs -jar")
	}
	if err := os.MkdirAll(*out, 0o755); err != nil {
		fatal(err.Error())
	}
	cases, sets, err := loadCases(*repo, *set, *limit, *deep, *idsFile)
	if err != nil {
		fatal(err.Error())
	}
	if len(cases) == 0 {
		fatal("no cases selected")
	}
	if err := applyAcquisitionOverrides(cases, acquisitionOverrides); err != nil {
		fatal(err.Error())
	}
	scratch, err := os.MkdirTemp(*out, "scratch-")
	if err != nil {
		fatal(err.Error())
	}
	defer os.RemoveAll(scratch)

	// progress.jsonl gets each case as it finishes, so a run that is stopped early keeps
	// what it has; results.jsonl and report.md are written at the end in manifest order.
	pf, err := os.Create(filepath.Join(*out, "progress.jsonl"))
	if err != nil {
		fatal(err.Error())
	}
	defer pf.Close()
	progress := json.NewEncoder(pf)
	var harnessErr error
	var mu sync.Mutex
	work := make(chan *Case)
	var wg sync.WaitGroup
	for i := 0; i < *jobs; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := range work {
				// Java side: the recorded answer, or a fresh run.
				if *useJava {
					n, err := RunJava(*java, *jar, c.Dir, c.Root, scratch, *timeout)
					if err != nil {
						mu.Lock()
						harnessErr = err
						mu.Unlock()
						continue
					}
					c.Sany = n
				} else if c.Status == "crash" {
					c.Sany = Norm{Class: ClassCrash, Note: "recorded crash"}
				} else {
					c.Sany = NormalizeSany(c.Stdout)
				}
				c.Prudence = RunPrudence(c.Dir, c.Root, *stage, *timeout, acquisitionVisits(c))
				c.Result, c.Kind, c.Policy = Compare(c.Sany, c.Prudence)
				c.FalseComplete = c.Prudence.Completion == "complete" && c.Prudence.Class == ClassOK && c.Sany.Class != ClassOK && c.Sany.Class != ClassCrash
				mu.Lock()
				if err := progress.Encode(caseRecord(c)); err != nil && harnessErr == nil {
					harnessErr = err
				}
				mu.Unlock()
			}
		}()
	}
	for _, c := range cases {
		work <- c
	}
	close(work)
	wg.Wait()
	if harnessErr != nil {
		fatal("harness failure: " + harnessErr.Error())
	}
	if err := writeOutputs(*out, cases, sets, *useJava, *jar, *stage); err != nil {
		fatal(err.Error())
	}
	if n := abandonedChecks.Load(); n > 0 {
		fmt.Fprintf(os.Stderr, "compare: %d Check calls ignored the deadline and were abandoned (counted as timeouts)\n", n)
	}
}

func fatal(msg string) {
	fmt.Fprintln(os.Stderr, "compare:", msg)
	os.Exit(2)
}

// loadCases lists the manifests of the sets, in a fixed order.
func loadCases(repo, only string, limit int, deep bool, idsFile string) ([]*Case, []string, error) {
	roots := []string{"sets"}
	if deep {
		roots = append(roots, "deep-cuts")
	}
	var manifests []string
	for _, r := range roots {
		m, _ := filepath.Glob(filepath.Join(repo, r, "*", "*", "manifest.jsonl"))
		manifests = append(manifests, m...)
	}
	sort.Strings(manifests)
	var wantIDs map[string]bool
	if idsFile != "" {
		b, err := os.ReadFile(idsFile)
		if err != nil {
			return nil, nil, err
		}
		wantIDs = map[string]bool{}
		for _, l := range strings.Fields(string(b)) {
			wantIDs[l] = true
		}
	}
	var cases []*Case
	var sets []string
	for _, m := range manifests {
		dir := filepath.Dir(m)
		name := filepath.Base(dir)
		if r, _ := filepath.Rel(repo, filepath.Dir(filepath.Dir(dir))); r == "deep-cuts" {
			name = "deep-cuts/" + name
		}
		if only != "" && only != name && only != filepath.Base(dir) {
			continue
		}
		f, err := os.Open(m)
		if err != nil {
			return nil, nil, err
		}
		sc := bufio.NewScanner(f)
		sc.Buffer(make([]byte, 1<<20), 1<<30)
		first, n := true, 0
		for sc.Scan() {
			if first { // line one describes the environment, not a case
				first = false
				continue
			}
			var row manifestRow
			if err := json.Unmarshal(sc.Bytes(), &row); err != nil {
				f.Close()
				return nil, nil, fmt.Errorf("%s: %v", m, err)
			}
			if row.ID == "" {
				continue
			}
			if wantIDs != nil && !wantIDs[row.ID] {
				continue
			}
			if limit > 0 && n >= limit {
				break
			}
			n++
			cases = append(cases, &Case{Set: name, ID: row.ID, Dir: filepath.Join(dir, "cases", row.ID),
				Root: row.RootFile, Status: row.Expected.Status, Stdout: row.Expected.Stdout})
		}
		f.Close()
		if err := sc.Err(); err != nil {
			return nil, nil, err
		}
		if n > 0 {
			sets = append(sets, name)
		}
	}
	return cases, sets, nil
}
