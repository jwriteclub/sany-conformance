// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

// Command sanyharness records what the plain SANY command line does not print: the
// error and warning codes and their locations, through tools/harness/SanyProbe.java
// over SANY's public API, and hashes of the syntax tree and of the semantic graph,
// the latter from the unmodified tla2sany.xml.XMLExporter. Everything it writes is
// labelled "harness" in the manifests. JVM flags are the defaults, except that each
// JVM gets a private java.io.tmpdir.
//
//	sanyharness -jar J -probe DIR -cases cases.jsonl -out out.jsonl [-java J] [-j N]
//
// Each input line is {"id","path","root_file"}.
package main

import (
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"sany-conformance/internal/canon"
)

type input struct {
	ID       string `json:"id"`
	Path     string `json:"path"`
	RootFile string `json:"root_file"`
}

type message struct {
	Code     int    `json:"code"`
	Severity string `json:"severity"`
	Module   string `json:"module"`
	BL       int    `json:"bl"`
	BC       int    `json:"bc"`
}

type probeResult struct {
	Exit     int       `json:"exit"`
	SyntaxOK bool      `json:"syntax_ok"`
	Crash    *string   `json:"crash"`
	Parse    []message `json:"parse_messages"`
	Semantic []message `json:"semantic_messages"`
	TreeOK   bool      `json:"tree_written"`
	TreeErr  *string   `json:"tree_error"`
}

// Row is one case's harness column.
type Row struct {
	ID            string `json:"id"`
	Status        string `json:"status"` // ok, crash, timeout, infra
	Detail        string `json:"detail,omitempty"`
	Verdict       string `json:"verdict,omitempty"`
	SyntaxOK      *bool  `json:"syntax_ok,omitempty"`
	ErrorCodes    []int  `json:"error_codes"`
	WarningCodes  []int  `json:"warning_codes"`
	FirstError    string `json:"first_error_location,omitempty"`
	SyntaxTree    string `json:"syntax_tree_sha256,omitempty"`
	SemanticGraph string `json:"semantic_graph_sha256,omitempty"`
	GraphState    string `json:"semantic_graph_state,omitempty"`
}

func main() {
	jar := flag.String("jar", "", "tla2tools.jar")
	probe := flag.String("probe", "", "directory with the compiled SanyProbe")
	java := flag.String("java", "java", "java binary")
	cases := flag.String("cases", "", "JSON lines of id, path, root_file")
	out := flag.String("out", "", "output JSON lines")
	jobs := flag.Int("j", 4, "cases at once")
	timeout := flag.Duration("timeout", 300*time.Second, "per JVM")
	work := flag.String("work", os.TempDir(), "directory for temporary copies")
	flag.Parse()
	if err := run(*jar, *probe, *java, *cases, *out, *work, *jobs, *timeout); err != nil {
		fmt.Fprintln(os.Stderr, "sanyharness:", err)
		os.Exit(2)
	}
}

func run(jar, probe, java, cases, out, work string, jobs int, timeout time.Duration) error {
	f, err := os.Open(cases)
	if err != nil {
		return err
	}
	defer f.Close()
	var in []input
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 0, 1<<20), 1<<24)
	for sc.Scan() {
		var c input
		if err := json.Unmarshal(sc.Bytes(), &c); err != nil {
			return err
		}
		in = append(in, c)
	}
	if err := sc.Err(); err != nil {
		return err
	}
	base, err := os.MkdirTemp(work, "harness-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(base)
	rows := make([]Row, len(in))
	next := make(chan int)
	var wg sync.WaitGroup
	var mu sync.Mutex
	done := 0
	for w := 0; w < jobs; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range next {
				rows[i] = one(in[i], i, jar, probe, java, base, timeout)
				mu.Lock()
				done++
				if done%500 == 0 {
					fmt.Fprintf(os.Stderr, "%d of %d\n", done, len(in))
				}
				mu.Unlock()
			}
		}()
	}
	for i := range in {
		next <- i
	}
	close(next)
	wg.Wait()
	o, err := os.Create(out)
	if err != nil {
		return err
	}
	defer o.Close()
	enc := json.NewEncoder(o)
	enc.SetEscapeHTML(false)
	for i := range rows {
		if err := enc.Encode(&rows[i]); err != nil {
			return err
		}
	}
	return nil
}

// copyTree copies a case directory, following symbolic links to library modules.
func copyTree(from, to string) error {
	es, err := os.ReadDir(from)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(to, 0o755); err != nil {
		return err
	}
	for _, e := range es {
		b, err := os.ReadFile(filepath.Join(from, e.Name()))
		if err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(to, e.Name()), b, 0o644); err != nil {
			return err
		}
	}
	return nil
}

type proc struct {
	exit     int
	timedOut bool
	stdout   string
	stderr   string
	err      error
}

func spawn(dir string, timeout time.Duration, tag string, argv ...string) proc {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	cmd := exec.Command(argv[0], argv[1:]...)
	cmd.Dir = dir
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	so, se := filepath.Join(filepath.Dir(dir), tag+".out"), filepath.Join(filepath.Dir(dir), tag+".err")
	fo, err := os.Create(so)
	if err != nil {
		return proc{err: err}
	}
	fe, err := os.Create(se)
	if err != nil {
		fo.Close()
		return proc{err: err}
	}
	cmd.Stdout, cmd.Stderr = fo, fe
	if err := cmd.Start(); err != nil {
		fo.Close()
		fe.Close()
		return proc{err: err}
	}
	wait := make(chan error, 1)
	go func() { wait <- cmd.Wait() }()
	p := proc{stdout: so, stderr: se}
	select {
	case <-wait:
	case <-ctx.Done():
		p.timedOut = true
		_ = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		<-wait
	}
	fo.Close()
	fe.Close()
	p.exit = cmd.ProcessState.ExitCode()
	return p
}

func sha(b []byte) string {
	h := sha256.Sum256(b)
	return hex.EncodeToString(h[:])
}

func head(path string, n int64) string {
	f, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer f.Close()
	b, _ := io.ReadAll(io.LimitReader(f, n))
	return string(b)
}

var (
	reInternal = regexp.MustCompile(`(?m)^(java\.lang\.[A-Za-z]+(Exception|Error)\b.*|tla2sany\.[A-Za-z.]*Exception\b.*)$`)
	reEscaped  = regexp.MustCompile(`(?m)^Exception in thread .*$`)
)

func one(c input, n int, jar, probe, java, base string, timeout time.Duration) Row {
	r := Row{ID: c.ID, Status: "ok", ErrorCodes: []int{}, WarningCodes: []int{}}
	d := filepath.Join(base, "r"+strconv.Itoa(n))
	defer os.RemoveAll(d)
	run := filepath.Join(d, "case")
	tmp := filepath.Join(d, "tmp")
	if err := os.MkdirAll(tmp, 0o755); err != nil {
		r.Status, r.Detail = "infra", err.Error()
		return r
	}
	if err := copyTree(c.Path, run); err != nil {
		r.Status, r.Detail = "infra", err.Error()
		return r
	}
	res, tree := filepath.Join(d, "probe.json"), filepath.Join(d, "tree.txt")
	p := spawn(run, timeout, "probe", java, "-Djava.io.tmpdir="+tmp, "-cp", jar+":"+probe, "SanyProbe", res, tree, c.RootFile)
	if p.err != nil {
		r.Status, r.Detail = "infra", p.err.Error()
		return r
	}
	if p.timedOut {
		r.Status, r.Detail = "timeout", "probe"
		return r
	}
	raw, err := os.ReadFile(res)
	if err != nil || p.exit != 0 {
		r.Status = "crash"
		r.Detail = fmt.Sprintf("probe exit %d: %s", p.exit, clip(head(p.stderr, 1<<14)+head(p.stdout, 1<<14), 600))
		return r
	}
	var po probeResult
	if err := json.Unmarshal(raw, &po); err != nil {
		r.Status, r.Detail = "infra", "probe output: "+err.Error()
		return r
	}
	errs, warns := map[int]bool{}, map[int]bool{}
	for _, list := range [][]message{po.Parse, po.Semantic} {
		for _, m := range list {
			if m.Severity == "ERROR" {
				if !errs[m.Code] && len(errs) == 0 {
					r.FirstError = fmt.Sprintf("%s:%d:%d", m.Module, m.BL, m.BC)
				}
				errs[m.Code] = true
			} else {
				warns[m.Code] = true
			}
		}
	}
	r.ErrorCodes, r.WarningCodes = sorted(errs), sorted(warns)
	stdout := head(p.stdout, 1<<26)
	if po.Crash != nil {
		r.Status, r.Detail = "crash", clip(*po.Crash, 600)
	} else if m := reInternal.FindString(stdout); m != "" {
		r.Status, r.Detail = "crash", "caught inside SANY: "+clip(m, 600)
	}
	if po.TreeOK && po.TreeErr == nil {
		if b, err := os.ReadFile(tree); err == nil {
			r.SyntaxTree = sha(b)
		}
	}
	if r.Status == "crash" {
		r.Verdict = "reject"
		return r
	}
	ok := po.SyntaxOK
	r.SyntaxOK = &ok
	if po.Exit != 0 {
		r.Verdict, r.GraphState = "reject", "not-exported:rejected"
		return r
	}
	r.Verdict = "accept"
	x := spawn(run, timeout, "xml", java, "-Djava.io.tmpdir="+tmp, "-cp", jar, "tla2sany.xml.XMLExporter", "-o", "-t", c.RootFile)
	switch {
	case x.err != nil:
		r.GraphState = "infra:" + x.err.Error()
	case x.timedOut:
		r.GraphState = "timeout"
	case x.exit != 0:
		text := head(x.stderr, 1<<16)
		if m := reEscaped.FindString(text); m != "" {
			r.GraphState = "crash:" + clip(m, 200)
		} else {
			r.GraphState = fmt.Sprintf("export-failed:exit=%d", x.exit)
		}
	default:
		b, err := os.ReadFile(x.stdout)
		if err != nil {
			r.GraphState = "infra:" + err.Error()
			break
		}
		cn, st, err := canon.CanonicalXML(b)
		if err != nil {
			r.GraphState = "canon-error:" + clip(err.Error(), 200)
			break
		}
		r.SemanticGraph = sha(cn)
		r.GraphState = fmt.Sprintf("ok:entries=%d,refs=%d,unreachable=%d,dangling=%d", st.Entries, st.Refs, st.Unreachable, st.Dangling)
	}
	return r
}

func sorted(m map[int]bool) []int {
	out := make([]int, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Ints(out)
	return out
}

func clip(s string, n int) string {
	s = strings.TrimSpace(s)
	if len(s) > n {
		return s[:n] + "..."
	}
	return s
}
