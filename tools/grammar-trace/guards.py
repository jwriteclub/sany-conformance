#!/usr/bin/env python3
# Copyright 2026 Christopher O'Connell
# SPDX-License-Identifier: Apache-2.0
"""For each syntax-error arm of the generated parser that is a switch default,
compare the arm's case labels with the tokens that can get to it, read from what
guards it. Reads generated source only; runs nothing.

FIRST of a lookahead routine is over-approximated: semantic predicates are taken
as true, so a guard set inside the case labels holds whatever the predicates do.
Output: TSV of line, production, guard, tokens let through, tokens with no label.

usage: guards.py TLAplusParser.java TLAplusParserConstants.java
"""
import re
import sys

src = open(sys.argv[1], encoding="utf-8").read().split("\n")
cons = {}
for m in re.finditer(r"int (\w+) = (\d+);", open(sys.argv[2], encoding="utf-8").read()):
    cons.setdefault(int(m.group(2)), m.group(1))

routines = {}
i = 0
head = re.compile(r"final private boolean (jj_3(?:R)?_\d+)\(\) \{")
while i < len(src):
    m = head.search(src[i])
    if m:
        j = i + 1
        while src[j] != "  }":
            j += 1
        routines[m.group(1)] = [s.strip() for s in src[i + 1 : j]]
        i = j
    i += 1

item = re.compile(r"(?:!jj_semLA \|\| )?(jj_scan_token\((\w+)\)|(jj_3(?:R)?_\d+)\(\))")
memo = {}


def name_of(a):
    return cons[int(a)] if a.isdigit() else a


def first(r):
    """Returns (set of token names, nullable)."""
    if r in memo:
        return memo[r]
    memo[r] = (set(), True)  # cycles: lookahead routines of left-free grammars have none at FIRST
    out = set()
    nullable = True
    alts = []
    for t in routines[r]:
        m = re.match(r"if \((.*)\) (return true;|\{|jj_scanpos = xsp;|break;)$", t)
        if not m:
            continue
        im = item.fullmatch(m.group(1))
        if not im:
            continue
        if im.group(2):
            f, n = {name_of(im.group(2))}, False
        else:
            f, n = first(im.group(3))
        out |= f
        tail = m.group(2)
        if tail == "{":
            alts.append(n)
            continue
        if tail == "return true;":
            if all(not a for a in alts) and not n:
                nullable = False
                break
            alts = []
            continue
        # Optional or loop body: nothing consumed for sure.
        if alts:
            alts.append(n)
            continue
    memo[r] = (out, nullable)
    return memo[r]


def method_of(n):
    for k in range(n, 0, -1):
        m = re.search(r"final (?:public|private) \S+ (\w+)\(", src[k - 1])
        if m:
            return m.group(1)
    return "?"


def switch_labels(default_line):
    """Labels of the switch that owns the default at default_line."""
    depth = 0
    k = default_line - 1
    labels = []
    while k > 0:
        k -= 1
        s = src[k]
        depth += s.count("}") - s.count("{")
        if depth < 0:
            if "switch" in s:
                return k + 1, labels
            depth = 0  # left a nested block of an enclosing arm is not expected here
            raise SystemExit(f"no switch for {default_line}")
        if depth == 0:
            m = re.match(r"\s*case (\w+):", s)
            if m:
                labels.append(m.group(1))
    raise SystemExit(f"no switch for {default_line}")



def call_sites(prod):
    end = next(i for i, l in enumerate(src) if "boolean jj_2_1(int" in l)
    return [
        i + 1
        for i in range(end)
        if re.search(r"\b" + prod + r"\(", src[i]) and "final public" not in src[i] and not src[i].strip().startswith(("*", "//", "/*"))
    ]


def immediate_guard(line):
    """What sits right above a call: a run of case labels, or a jj_2 test."""
    k = line - 2
    while k > 0:
        t = src[k].strip()
        if t.startswith(("/*", "*", "//")) or t == "":
            k -= 1
            continue
        break
    t = src[k].strip()
    if re.match(r"case \w+:", t):
        labels = []
        while True:
            m = re.match(r"\s*case (\w+):", src[k])
            if m:
                labels.append(m.group(1))
                k -= 1
                continue
            s_ = src[k].strip()
            if s_.startswith(("/*", "*", "//")):
                k -= 1
                continue
            break
        return "case", labels
    m = re.search(r"jj_2_(\d+)\(", t)
    if m:
        return "jj", m.group(1)
    return "other", t


# Productions that are only called through one wrapper.
WRAPPER = {
    "InfixOpToken": ["InfixOp", "InfixLHS"],
    "PostfixOpToken": ["PostfixOp", "PostfixLHS"],
    "PrefixOpToken": ["PrefixOp"],
    "NEPrefixOpToken": ["NonExpPrefixOp", "PrefixLHS"],
}


# Switches past the start of a production: the jj_2 test or case arm just outside.
INNER = {6051: 54, 6367: 62, 4622: 38, 7149: 73, 3223: 18, 3316: 20, 3526: 25, 6935: 57}
# Nested on the same token: 0-based start and end of the outer arm's labels.
OUTER_LABELS = {4927: (4892, 4894), 5042: (4982, 4984), 6631: (6621, 6623), 1496: (1480, 1482)}


def arm_rows():
    rows = []
    for n, l in enumerate(src, 1):
        if not (re.match(r"\s*default:\s*$", l) and "jj_consume_token(-1)" in " ".join(src[n : n + 3])):
            continue
        prod = method_of(n)
        sw, labels = switch_labels(n)
        rows.append((n, prod, sw, set(labels)))
    return rows


def main():
    print("error_arm\tproduction\tcase_labels\tentered_from\tguard\tguard_tokens\tguard_tokens_without_a_label")
    wanted = set(sys.argv[3].split(","))
    for n, prod, sw, labels in arm_rows():
        if str(n) not in wanted:
            continue
        sites = []
        # The switch at the start of the production: its callers guard it. Later
        # switches are read by hand, marked inner.
        if n in INNER:
            g = INNER[n]
            g_set = first(f"jj_3_{g}")[0]
            extra = sorted(g_set - labels)
            print(f"{n}\t{prod}\t{len(labels)}\tinner switch@{sw}\tjj_2_{g}\t{len(g_set)}\t{' '.join(extra)}")
            continue
        if n in OUTER_LABELS:
            k = OUTER_LABELS[n]
            kind, g = "case", [m for m in re.findall(r"case (\w+):", " ".join(src[k[0] : k[1]]))]
            extra = sorted(set(g) - labels)
            print(f"{n}\t{prod}\t{len(labels)}\tnested switch@{sw}\tcase labels at :{k[0] + 1}\t{len(set(g))}\t{' '.join(extra)}")
            continue
        names = WRAPPER.get(prod, [prod])
        for name in names:
            sites += [(name, c) for c in call_sites(name)]
        if not sites:
            print(f"{n}\t{prod}\t{len(labels)}\t(no caller)\t\t\t")
        for name, c in sites:
            kind, g = immediate_guard(c)
            if kind == "case":
                g_set = set(g)
                gdesc = "case labels"
            elif kind == "jj":
                g_set = first(f"jj_3_{g}")[0]
                gdesc = f"jj_2_{g}"
            else:
                g_set = set()
                gdesc = "read: " + g[:60]
            extra = sorted(g_set - labels)
            print(f"{n}\t{prod}\t{len(labels)}\t{name}@{c}\t{gdesc}\t{len(g_set)}\t{' '.join(extra)}")


main()
