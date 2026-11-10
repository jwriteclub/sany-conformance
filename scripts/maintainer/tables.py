#!/usr/bin/env python3
# Copyright 2026 Christopher O'Connell
# SPDX-License-Identifier: Apache-2.0
"""Print the tables README.md and FINDINGS.md quote, from the committed manifests and data."""
import collections
import glob
import gzip
import json
import os
import re
import sys

root = sys.argv[1] if len(sys.argv) > 1 else "."
rows = {}
for m in glob.glob(root + "/sets/**/manifest.jsonl", recursive=True) + glob.glob(root + "/deep-cuts/**/manifest.jsonl", recursive=True):
    rel = os.path.relpath(os.path.dirname(m), root)
    with open(m, encoding="utf-8") as f:
        f.readline()
        for line in f:
            r = json.loads(line)
            r["dir"] = rel
            rows[r["id"]] = r
print("rows", len(rows))

print("\n## Verdicts per place")
c = collections.Counter()
for r in rows.values():
    c[(r["dir"].split("/")[0], r["expected"]["status"], r["expected"]["verdict"])] += 1
for k, v in sorted(c.items(), key=str):
    print(k, v)

print("\n## Crashes by signature")
sig = collections.defaultdict(list)
for r in rows.values():
    e = r["expected"]
    if e["status"] == "crash":
        cr = e["crash"]
        msg = re.sub(r"\d+", "N", re.sub(r'^(Exception in thread "main" |tla2sany\.drivers\.FrontEndException: )', "", cr["line"]))
        size = sum(os.path.getsize(os.path.join(root, r["dir"], "cases", r["id"], f["path"])) for f in r["files"] if f["role"] != "library") if os.path.isdir(os.path.join(root, r["dir"], "cases", r["id"])) else 10**9
        sig[(cr["exception"], cr["frame"], cr["kind"], msg[:90])].append((size, r["id"], r["dir"]))
for k, v in sorted(sig.items(), key=lambda x: (-len(x[1]), x[0][0])):
    v.sort()
    print("| `%s` | `%s` | %s | %d | `%s/cases/%s` | %d |" % (k[0], k[1] or "-", k[2], len(v), v[0][2], v[0][1], v[0][0]))
    print("   msg:", k[3])

print("\n## tlx disagreements")
sigs = collections.defaultdict(list)
with open(root + "/data/elements-extra.tsv") as f:
    for l in f:
        if l.startswith("#"):
            continue
        i, e = l.rstrip("\n").split("\t", 1)
        if e.startswith("tlx "):
            sigs[e[4:]].append(i)
d = collections.Counter()
for s, ids in sigs.items():
    d[s.split(" code=")[0]] += len(ids)
print(len(sigs), "signatures", d)
for s, ids in sorted(sigs.items(), key=lambda x: -len(x[1]))[:60]:
    ids2 = sorted((i for i in ids if i in rows), key=lambda i: (sum(1 for _ in rows[i]["files"]), rows[i]["id"]))
    print(len(ids), s, ids2[:2])

print("\n## Harness graph states")
g = collections.Counter()
for r in rows.values():
    st = r.get("harness", {}).get("semantic_graph_state", "")
    g[st.split(":")[0] + (":" + st.split(":")[1][:60] if st.startswith(("crash", "export")) else "")] += 1
print(g)
print("codes distinct", len({c for r in rows.values() for c in r.get("harness", {}).get("error_codes", [])}), len({c for r in rows.values() for c in r.get("harness", {}).get("warning_codes", [])}))
for k in ("export-failed", "crash"):
    ex = [r["id"] for r in rows.values() if r.get("harness", {}).get("semantic_graph_state", "").startswith(k)]
    print(k, len(ex), ex[:8])
print("\n## Dependence")
dep = collections.Counter()
for r in rows.values():
    for t in r.get("depends_on", []):
        dep[t] += 1
print(dep)
print("\n## Exit codes")
print(collections.Counter((r["expected"]["verdict"], r["expected"].get("exit_plain"), r["expected"].get("exit_error_codes")) for r in rows.values()))
print("\n## Licences")
lc = collections.Counter((r["licence"]["spdx"], r["licence"]["file"]) for r in rows.values())
for k, v in sorted(lc.items(), key=str):
    print(k, v)
