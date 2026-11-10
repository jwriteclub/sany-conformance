#!/usr/bin/env python3
# Copyright 2026 Christopher O'Connell
# SPDX-License-Identifier: Apache-2.0
"""Compare the stock runs and write results/. One-time, for a maintainer holding the
raw runs. Every difference gets a row; nothing is summarised away.

usage: analyse_runs.py --repo R --runs DIR --cases cases.json --baseline B
"""
import argparse
import base64
import collections
import gzip
import hashlib
import importlib.machinery
import importlib.util
import json
import os
import sys

HEAD = "# Copyright 2026 Christopher O'Connell\n# SPDX-License-Identifier: Apache-2.0\n"


def verify_module(repo):
    path = os.path.join(repo, "scripts", "verify")
    loader = importlib.machinery.SourceFileLoader("verify", path)
    spec = importlib.util.spec_from_loader("verify", loader)
    m = importlib.util.module_from_spec(spec)
    sys.dont_write_bytecode = True
    loader.exec_module(m)
    return m


def load_run(path):
    out = {}
    env = None
    with open(path, encoding="utf-8") as f:
        for line in f:
            r = json.loads(line)
            if "environment" in r:
                env = r["environment"]
            else:
                out[r["id"]] = r
    return env, out


def reclassify(V, run):
    """Apply the current reading of SANY's text, so that every run is read one way."""
    for r in run.values():
        if r["status"] in ("ok", "crash"):
            for k in ("crash",):
                r.pop(k, None)
            r.update(V.classify(text(r, "stdout"), text(r, "stderr")))
            if r.get("exit_error_codes") is not None and r["status"] == "ok":
                accepted = r["exit_error_codes"] == 0
                if accepted != (r["verdict"] == "accept"):
                    raise SystemExit("%s: text and -error-codes disagree" % r["id"])
    return run


def text(r, key):
    if key in r:
        return r[key]
    if key + "_base64" in r:
        return base64.b64decode(r[key + "_base64"]).decode("utf-8", "surrogateescape")
    return ""


def compare(V, a, b):
    """Field-level differences between two observations of one case."""
    out = []
    for f in ("status", "verdict", "first_position"):
        if a.get(f) != b.get(f):
            out.append((f, a.get(f), b.get(f)))
    ca = V.crash_of(text(a, "stdout"), text(a, "stderr")) or {}
    cb = V.crash_of(text(b, "stdout"), text(b, "stderr")) or {}
    if (ca.get("exception"), ca.get("frame")) != (cb.get("exception"), cb.get("frame")):
        out.append(("crash", "%s@%s" % (ca.get("exception"), ca.get("frame")), "%s@%s" % (cb.get("exception"), cb.get("frame"))))
    so_a, so_b = text(a, "stdout"), text(b, "stdout")
    if so_a != so_b:
        kind = "stdout-progress-lines-only" if V.soft_text(so_a) == V.soft_text(so_b) else "stdout-messages"
        out.append((kind, hashlib.sha256(so_a.encode("utf-8", "surrogateescape")).hexdigest()[:12], hashlib.sha256(so_b.encode("utf-8", "surrogateescape")).hexdigest()[:12]))
    se_a, se_b = text(a, "stderr"), text(b, "stderr")
    if se_a != se_b:
        out.append(("stderr", hashlib.sha256(se_a.encode("utf-8", "surrogateescape")).hexdigest()[:12], hashlib.sha256(se_b.encode("utf-8", "surrogateescape")).hexdigest()[:12]))
    return out


def write_diffs(path, title, rows, sets):
    with gzip.GzipFile(path, "wb", 9, mtime=0) if path.endswith(".gz") else open(path, "wb") as f:
        out = HEAD + "# " + title + "\nset\tid\tfield\treference\tthis_run\n"
        for i, field, x, y in rows:
            out += "%s\t%s\t%s\t%s\t%s\n" % (sets.get(i, ""), i, field, clean(x), clean(y))
        f.write(out.encode("utf-8", "replace"))


def clean(v):
    return str(v).replace("\t", " ").replace("\n", " ")[:300]


def main():
    ap = argparse.ArgumentParser()
    for name in ("repo", "runs", "cases", "baseline"):
        ap.add_argument("--" + name, required=True)
    a = ap.parse_args()
    V = verify_module(a.repo)
    res = a.repo + "/results"
    os.makedirs(res, exist_ok=True)
    cases = [json.loads(l) for l in open(a.cases)]
    sets = {c["ID"]: c["Group"] for c in cases}
    env1, r1 = load_run(a.runs + "/r1.jsonl")
    reclassify(V, r1)
    summary = {"reference": env1}
    depends = collections.defaultdict(set)
    TAGS = {"r2": "run-to-run", "mac1": "platform", "c1": "jvm-flag-c1", "xss512k": "stack-512k", "xss16m": "stack-16m", "posix": "charset-posix", "latin1": "charset-latin1"}
    for name, title in (
        ("r2", "Second run of the reference: dc6a944, Linux, stock. Order reversed."),
        ("mac1", "The same jar on macOS arm64, APFS, default flags."),
        ("rel174", "The latest official release, v1.7.4, on the reference machine. java.io.tmpdir per JVM, since this jar unpacks standard modules into a shared directory."),
        ("rel180", "The v1.8.0 pre-release of 2026-11-03, on the reference machine."),
        ("c1", "The reference with -XX:TieredStopAtLevel=1, the flag the earlier harness used."),
        ("xss512k", "The reference with -Xss512k."),
        ("xss16m", "The reference with -Xss16m."),
        ("posix", "The reference with LC_ALL=C: default charset ANSI_X3.4-1968."),
        ("latin1", "The reference with -Dfile.encoding=ISO-8859-1."),
        ("ntk", "A jar with every memoised jj_ntk read replaced by jj_ntk()."),
    ):
        p = "%s/%s.jsonl" % (a.runs, name)
        if not os.path.exists(p):
            print("MISSING", name)
            continue
        env, run = load_run(p)
        reclassify(V, run)
        rows = []
        per_field = collections.Counter()
        ids = set()
        for i in sorted(r1):
            if i not in run:
                rows.append((i, "missing", "", ""))
                per_field["missing"] += 1
                ids.add(i)
                continue
            for field, x, y in compare(V, r1[i], run[i]):
                rows.append((i, field, x, y))
                per_field[field] += 1
                ids.add(i)
                if name in TAGS:
                    depends[i].add("%s:%s" % (TAGS[name], "text" if field.startswith(("stdout", "stderr")) else field))
        write_diffs("%s/diff-%s.tsv.gz" % (res, name), title, rows, sets)
        verd = [i for i in ids if i in run and r1[i].get("verdict") != run[i].get("verdict")]
        by_dir = collections.Counter("%s->%s" % (r1[i].get("verdict"), run[i].get("verdict")) for i in verd)
        st = collections.Counter(r.get("status") for r in run.values())
        by_set = collections.Counter(sets[i] for i in ids)
        summary[name] = {"title": title, "environment": env, "cases": len(run), "statuses": st, "cases_differing": len(ids), "by_field": per_field, "verdict_changes": by_dir, "by_set": by_set, "verdict_changed_ids": sorted(verd)[:400]}
        print(name, "differing", len(ids), dict(per_field), dict(by_dir), dict(st))
    with open(res + "/depends-on.tsv", "w") as f:
        f.write(HEAD + "# Every case whose stock result changed with platform, JVM flags, stack size or charset, or between two runs of the reference. Tag: what was changed, then what moved.\nid\tdepends_on\n")
        for i in sorted(depends):
            f.write("%s\t%s\n" % (i, ",".join(sorted(depends[i]))))
    with open(a.runs + "/evidence-extra.tsv", "w") as f:
        for i in sorted(depends):
            real = sorted(t for t in depends[i] if not t.endswith(":text"))
            if real:
                f.write("%s\tits result depends on the environment (%s): see results/depends-on.tsv\n" % (i, ",".join(real)))
    # The harness column against the earlier harness results and the stock verdict.
    h = {}
    for line in open(a.runs + "/harness.jsonl"):
        r = json.loads(line)
        h[r["id"]] = r
    rows = []
    n = collections.Counter()
    for c in cases:
        i = c["ID"]
        s, hr = r1[i], h[i]
        cs = V.crash_of(text(s, "stdout"), text(s, "stderr"))
        if (hr["status"] == "crash") != bool(cs):
            rows.append((i, "crash-status", "stock %s" % s["status"], "harness %s %s" % (hr["status"], hr.get("detail", ""))))
            n["crash-status"] += 1
        if hr["status"] == "ok" and s["status"] == "ok" and hr.get("verdict") != s.get("verdict"):
            rows.append((i, "verdict", s.get("verdict"), hr.get("verdict")))
            n["verdict"] += 1
        # The index: earlier harness, C1 flags, -Dfile.encoding=UTF-8, library path.
        if c.get("L0"):
            old = "crash" if c.get("Outcome") == "crash" else c["L0"]
            new = "crash" if s["status"] == "crash" else s.get("verdict")
            if old != new:
                rows.append((i, "earlier-verdict", old, new))
                n["earlier-verdict"] += 1
            if hr["status"] == "ok":
                e = ",".join(map(str, hr["error_codes"]))
                w = ",".join(map(str, hr["warning_codes"]))
                if (c.get("L1e") or "") != e:
                    rows.append((i, "earlier-error-codes", c.get("L1e"), e))
                    n["earlier-error-codes"] += 1
                if (c.get("L1w") or "") != w:
                    rows.append((i, "earlier-warning-codes", c.get("L1w"), w))
                    n["earlier-warning-codes"] += 1
        elif c.get("Outcome") in ("crash", "hang") and s["status"] == "ok":
            rows.append((i, "earlier-outcome", c.get("Outcome"), s["status"]))
            n["earlier-outcome"] += 1
    # Tiers A to C: every level of the earlier run.
    run1 = a.baseline + "/results/java/run1/results.jsonl"
    seen = 0
    for line in open(run1):
        o = json.loads(line)
        i = o["case"]
        if i not in h:
            continue
        seen += 1
        lv = o.get("levels") or {}
        val = lambda k: (lv.get(k) or {}).get("value")
        hr, s = h[i], r1[i]
        if val("L2") and hr.get("syntax_tree_sha256") != val("L2"):
            rows.append((i, "earlier-syntax-tree", val("L2")[:16], (hr.get("syntax_tree_sha256") or "")[:16]))
            n["earlier-syntax-tree"] += 1
        if val("L3") and (hr.get("semantic_graph_sha256") or hr.get("semantic_graph_state")) != val("L3"):
            rows.append((i, "earlier-semantic-graph", val("L3")[:30], (hr.get("semantic_graph_sha256") or hr.get("semantic_graph_state") or "")[:30]))
            n["earlier-semantic-graph"] += 1
        if val("L1p") and s.get("first_position") != val("L1p") and s["status"] == "ok":
            rows.append((i, "earlier-first-position", val("L1p"), s.get("first_position")))
            n["earlier-first-position"] += 1
    write_diffs(res + "/diff-harness.tsv", "Stock results against the harness: the harness column of this run, and the earlier results that came through the wrapper with -XX:TieredStopAtLevel=1.", rows, sets)
    summary["harness"] = {"by_field": n, "earlier_run_cases": seen, "statuses": collections.Counter(r["status"] for r in h.values()), "graph_states": collections.Counter((r.get("semantic_graph_state") or "").split(":")[0] for r in h.values())}
    print("harness", dict(n), "earlier", seen, summary["harness"]["statuses"], summary["harness"]["graph_states"])
    with open(res + "/stock-summary.json", "w") as f:
        json.dump({"copyright": "Copyright 2026 Christopher O'Connell", "spdx": "Apache-2.0", **summary}, f, indent=1, sort_keys=True)
        f.write("\n")


if __name__ == "__main__":
    main()
