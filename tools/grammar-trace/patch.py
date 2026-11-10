#!/usr/bin/env python3
# Copyright 2026 Christopher O'Connell
# SPDX-License-Identifier: Apache-2.0
"""Patch a copy of the generated TLAplusParser.java so that every LOOKAHEAD
decision and every syntax-error arm reports what it did. Line numbers do not
change, so events name lines of the original file. JaCoCo cannot see these: its
probe sits after a call, and the false side of `if (jj_2_k(1)) ... else
{ jj_consume_token(-1); throw ... }` always throws.

usage: patch.py TLAplusParser.java out/TLAplusParser.java
"""
import re
import sys

src = open(sys.argv[1], encoding="utf-8", newline="").read().split("\n")
# Productions end where the lookahead helpers start.
end = next(i for i, l in enumerate(src) if "boolean jj_2_1(int" in l)
dec = re.compile(r"^(\s*(?:\} else )?if \()(jj_2_\d+\(\d+\))(?: && \((.*)\))?\) \{$")
nd = ne = 0
out = []
for i, l in enumerate(src):
    n = i + 1
    if i < end:
        m = dec.match(l)
        if m:
            nd += 1
            if "jj_2_55(" in l or "jj_2_63(" in l:
                # OpenStart's two sites: which of its eight tokens, and the predicate.
                l = f"{m.group(1)}T.o({n}, getToken(1).kind, {m.group(2)}) && T.q({n}, getToken(1).kind, ({m.group(3)}))) {{"
            elif m.group(3) is not None:
                l = f"{m.group(1)}T.j({n}, {m.group(2)}) && T.p({n}, ({m.group(3)}))) {{"
            else:
                l = f"{m.group(1)}T.j({n}, {m.group(2)})) {{"
        elif "jj_2_" in l:
            sys.exit(f"line {n}: decision of an unknown shape: {l}")
        if "jj_consume_token(-1);" in l:
            ne += 1
            l = l.replace("jj_consume_token(-1);", f"T.e({n}); jj_consume_token(-1);")
    out.append(l)
open(sys.argv[2], "w", encoding="utf-8", newline="").write("\n".join(out))
print(f"patched {nd} decisions, {ne} error arms")
