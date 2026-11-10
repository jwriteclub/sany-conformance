#!/usr/bin/env python3
# Copyright 2026 Christopher O'Connell
# SPDX-License-Identifier: Apache-2.0
"""Lay out sets/ and deep-cuts/ with their manifests, READMEs and licence notes.

One-time, for a maintainer holding the measurement's outputs. A reader regenerates
cases with scripts/regenerate and checks them with scripts/verify.

usage: build_tree.py --repo R --tree stage/tree --cases cases.json --runs DIR
                     --support support.tsv --trace trace.tsv --baseline B
"""
import argparse
import collections
import csv
import glob
import hashlib
import importlib.machinery
import importlib.util
import json
import os
import re
import shutil
import sys

sys.dont_write_bytecode = True
HEADMD = "<!-- Copyright 2026 Christopher O'Connell -->\n<!-- SPDX-License-Identifier: Apache-2.0 -->\n"
HEADTSV = "# Copyright 2026 Christopher O'Connell\n# SPDX-License-Identifier: Apache-2.0\n"
REGEN_UP = "`scripts/regenerate` clones `sources.lock` and runs `cmd/extract`, which rewrites these cases byte for byte."

CATS = {
    "tla-project": "From the TLA+ project",
    "other-projects": "From other projects",
    "ours-mechanical": "Ours, generated mechanically",
    "ours-model": "Ours, model-guided",
    "ours-inspection": "Ours, by inspection",
}

# What each set is, where it came from, what it carries, how to regenerate it.
SETS = {
    "tlaplus-syntax-corpus": ("tla-project", "Unified syntax corpus", True,
        "The syntax corpus SANY's own tests run, including the Unicode files. One snippet per case.",
        "`test/tla2sany/corpus` in tlaplus/tlaplus at the pinned commit. Each snippet is wrapped into a module by `cmd/extract`.",
        "An upstream assertion, syntax stage only: accept or reject, and for accepts an expected S-expression tree, kept as `expected.sexp`. Many snippets use undefined names, so an accept there does not mean full SANY accepts, and the expected output here is full SANY's. 12 are in upstream's `expectFailures`: SANY deviates from them, recorded in `upstream_assertion`.",
        REGEN_UP),
    "tlaplus-semantic-corpus": ("tla-project", "Semantic corpus", True,
        "Specs upstream's semantic tests expect to pass with no errors or warnings.",
        "`test/tla2sany/semantic/corpus` in tlaplus/tlaplus.",
        "An upstream assertion: no errors, no warnings, embedded RefersTo and IsLevel assertions hold. `NegativeOpTest` is skipped upstream (tlaplus/tlaplus#1130).",
        REGEN_UP),
    "tlaplus-semantic-error-corpus": ("tla-project", "Semantic error corpus", True,
        "One spec per error or warning code, with the code in the file name.",
        "`test/tla2sany/semantic/error_corpus` in tlaplus/tlaplus.",
        "An upstream assertion: the code in the file name appears in the log, and E means failure. Nothing about the exact text, location or the rest of the code list; those here are SANY's output, recorded.",
        REGEN_UP),
    "tlaplus-inline-java": ("tla-project", "Specs written inline in Java tests", True,
        "The `TestLevelChecking` table and one `ParseErrorTests` entry, written as Java strings upstream and extracted by regular expression.",
        "`test/tla2sany/semantic/TestLevelChecking.java` and `test/tla2sany/parser/ParseErrorTests.java` in tlaplus/tlaplus.",
        "An upstream assertion: parse and semantics succeed with the level check result listed, or a syntax rejection with message text.",
        REGEN_UP),
    "tlaplus-test-models": ("tla-project", "TLC test models", True,
        "Modules under `test-model`. 12 are fixtures of SANY tests; the rest exist for TLC's tests and SANY only happens to read them.",
        "`tlatools/org.lamport.tlatools/test-model` in tlaplus/tlaplus.",
        "The 12 fixtures carry an upstream assertion: 2 reject with message text, 10 accept weakly (XML export exit 0, or no exception), one of them `@Ignore`d. The rest assert nothing about SANY.",
        REGEN_UP),
    "tlaplus-other": ("tla-project", "Other modules in the TLA+ tools repository", True,
        "Every other `.tla` file in tlaplus/tlaplus: `general`, the toolbox, tlatools' spec, source and test resources, and helper modules of the error corpus.",
        "tlaplus/tlaplus, by each row's `source.path`.",
        "Nothing beyond SANY's recorded output.",
        REGEN_UP),
    "tlaplus-examples": ("tla-project", "tlaplus/Examples", True,
        "The community collection of specifications.",
        "`specifications/` in tlaplus/Examples.",
        "Nothing beyond SANY's recorded output.",
        REGEN_UP),
    "tlaplus-communitymodules": ("tla-project", "tlaplus/CommunityModules", True,
        "The community modules, their tests and `.smoke` specs.",
        "`modules/`, `tests/` and `.smoke/` in tlaplus/CommunityModules.",
        "Nothing beyond SANY's recorded output.",
        REGEN_UP),
    "tlaplus-rfcs": ("tla-project", "tlaplus/rfcs language standard", True,
        "`language_standard/tests`, the origin of upstream's syntax, semantic and error corpora, and its syntax corpus. Nearly all duplicates of the corpora above, by content; kept, because this set is upstream's as it stands.",
        "`language_standard/` in tlaplus/rfcs.",
        "Its syntax-corpus rows are marked `other-parser`: another repository's assertion, kept for reference and not taken as an assertion about SANY. The rest carry nothing beyond SANY's recorded output.",
        REGEN_UP),
    "tlapm": ("other-projects", "tlapm, the TLA+ proof manager", True,
        "The proof manager's library, examples, tests and its copy of the syntax corpus. `test/sany/equivalence` compares tlapm's parser with SANY, tolerating SANY failure.",
        "tlaplus/tlapm, by each row's `source.path`.",
        "Its syntax-corpus rows are `other-parser`: tlapm's own test, nothing about SANY. The rest carry nothing beyond SANY's recorded output.",
        REGEN_UP),
    "tlx": ("other-projects", "tlx", True,
        "Specs from the tlx project: its tests, its fuzz directory, and `corpus/external`, tlx's wrapper modules around third-party specs that it does not vendor.",
        "tlaplus-x/tlx, by each row's `source.path`.",
        "Nothing beyond SANY's recorded output. tlx's own verdict is not an assertion about SANY.",
        REGEN_UP),
    "tree-sitter-tlaplus": ("other-projects", "tree-sitter-tlaplus", True,
        "The tree-sitter grammar's test corpus, unwrapped into modules. PlusCal-only material is left out, as SANY does not translate PlusCal.",
        "`test/` in tlaplus-community/tree-sitter-tlaplus.",
        "Its syntax-corpus rows are `other-parser`: a tree-sitter test, nothing about SANY. The rest carry nothing beyond SANY's recorded output.",
        REGEN_UP),
    "mutants": ("ours-mechanical", "Seeded mutants of the corpus", False,
        "3,000 mutants, seed 20261103, of upstream parents: delete a token, swap a delimiter, misalign a junction, use an unknown name, prime in an invariant, bad arity. At most 2 per parent, parents at most 30 KiB. `source` names each parent and rule.",
        "`cmd/extract/mutate.go`. Each is a small edit of its parent and follows its parent's licence.",
        "Nothing beyond SANY's recorded output. `source.intent` (`likely-reject` or `unknown`) is a hint from the rule, not an assertion.",
        "`scripts/regenerate`, which is seeded and deterministic."),
    "fuzz-jazzer": ("ours-mechanical", "Jazzer fuzz outputs", False,
        "What the Jazzer 0.30.0 fuzzer found on SANY, minimised by lines and tokens. Some contain bytes that are not valid UTF-8. 8 held a local file path, which was deleted, with verdict, exit status, codes and coverage measured unchanged; `source.derived_from` names the original.",
        "A harness over SANY seeded as `LICENSE.md` lists. libFuzzer records no parent, so `source.strategy` names only the strategy and unit. `source.trace`, where present, is the file closest by content (`cmd/tracecases`): provenance, not a licence claim.",
        "Nothing beyond SANY's recorded output.",
        "The bytes are here. Rerunning a fuzzer gives different cases."),
    "grammar-sentences": ("ours-mechanical", "Sentences derived from the grammar", False,
        "One short sentence per choice point of `tla+.jj`, in the manner of Purdom, ignoring semantic lookahead and actions, so many are rejected.",
        "`cmd/gensentences` over the pinned `tla+.jj`.",
        "Nothing beyond SANY's recorded output.",
        "`scripts/regenerate`."),
    "model-one-shot": ("ours-model", "One-shot model loop", False,
        "Specs from a loop giving a model the target code, taking one reply, and replaying it. `source.model` and `source.target` name them. Prompts carried nearby specs from the corpus, so a reply can echo them; `source.trace` measures that by content.",
        "Model replies, kept when a replay in a fresh JVM showed a new branch, error-code combination, tlx disagreement or Java crash.",
        "Nothing beyond SANY's recorded output.",
        "The bytes are here. Rerunning needs models and gives different cases."),
    "model-agents": ("ours-model", "Exploring agents", False,
        "Specs from sessions of a model with tools on one block of missed lines.",
        "Agent sessions; each row names the model and target.",
        "Nothing beyond SANY's recorded output.",
        "The bytes are here. Rerunning needs models and gives different cases."),
    "model-scoring": ("ours-model", "Scoring specs", False,
        "Specs a model wrote to back a `yes` verdict while scoring the remaining lines. None took its line.",
        "Scoring sessions; each row names the model and target.",
        "Nothing beyond SANY's recorded output.",
        "The bytes are here. Rerunning needs models and gives different cases."),
    "model-ternary-hunt": ("ours-model", "Specs hunting a stale jj_ntk", False,
        "Empty: models tried to make the memo in `(jj_ntk==-1)?jj_ntk():jj_ntk` go stale in 509 runs and none did, and nothing met the bar to be kept. The directory stays so that the set is accounted for. See FINDINGS.md.",
        "Agent sessions.",
        "Nothing.",
        "Nothing to regenerate."),
    "model-openstart": ("ours-model", "OpenStart scan modules", False,
        "Modules a model (GPT-6 Astra) wrote with no tools and without running them, aimed at the eight alternatives of the `OpenStart` lookahead scan. Evidence: they reach nothing new by coverage, and each has a stated purpose.",
        "One prompt, recorded in each row's provenance.",
        "Nothing beyond SANY's recorded output. The author's expectation is not an assertion.",
        "The bytes are here."),
    "model-pit-survivors": ("ours-model", "Model specs aimed at surviving mutants", False,
        "Specs models wrote, aimed at mutants of SANY that no earlier case noticed, that then killed mutants nothing earlier in the order did. The model is in `source.model`, the target in `source.aim`.",
        "A PIT 1.30.1 mutation run over SANY: the `\\* Aim:` line names the mutant.",
        "Nothing beyond SANY's recorded output.",
        "The bytes are here. Rerunning needs models and gives different cases."),
    "inspection": ("ours-inspection", "Written by reading the code", False,
        "Specs review agents wrote after reading the code, each aimed at something in it. `source.aim` is what it was aimed at, as its author or the review states it.",
        "Code reviews of SANY. Authors are in `source.author`.",
        "Nothing beyond SANY's recorded output.",
        "The bytes are here."),
    "inspection-pit-survivors": ("ours-inspection", "Written by reading a surviving mutant's code", False,
        "Specs agents wrote after reading the code of a mutant no case noticed. Each states its aim in `\\* Aim:` lines. Their file headers say all rights reserved; `LICENSE.md` says how that is read here.",
        "A PIT 1.30.1 mutation run over SANY.",
        "Nothing beyond SANY's recorded output.",
        "The bytes are here."),
}

CHECK_ORDER = list(SETS)


def load_verify(repo):
    path = os.path.join(repo, "scripts", "verify")
    loader = importlib.machinery.SourceFileLoader("verify", path)
    spec = importlib.util.spec_from_loader("verify", loader)
    m = importlib.util.module_from_spec(spec)
    loader.exec_module(m)
    return m


def sha(path):
    with open(path, "rb") as f:
        return hashlib.sha256(f.read()).hexdigest()


def read_licence_cols(baseline):
    out = {}
    for p in glob.glob(baseline + "/suite/*/cases.tsv"):
        with open(p, newline="", encoding="utf-8", errors="replace") as f:
            lines = [l for l in f if not l.startswith("#")]
        for r in csv.DictReader(lines, delimiter="\t", quoting=csv.QUOTE_NONE):
            out[r["id"]] = r["licence"]
    return out


RE_LIC = re.compile(r"^(?:derivative of \S+: )?(\S+) \((.+)\)$")


def parse_licence(s):
    m = RE_LIC.match(s or "")
    return (m.group(1), m.group(2)) if m else (None, None)


def main():
    ap = argparse.ArgumentParser()
    for name in ("repo", "tree", "cases", "runs", "support", "trace", "baseline", "paxosstore"):
        ap.add_argument("--" + name, required=True)
    a = ap.parse_args()
    V = load_verify(a.repo)
    repo = a.repo
    for d in ("sets", "deep-cuts"):
        shutil.rmtree(os.path.join(repo, d), ignore_errors=True)
    cases = [json.loads(l) for l in open(a.cases)]
    byid = {c["ID"]: c for c in cases}
    run = {}
    env = None
    for line in open(a.runs + "/r1.jsonl", encoding="utf-8"):
        r = json.loads(line)
        if "environment" in r:
            env = r["environment"]
        else:
            run[r["id"]] = r
    h = {json.loads(l)["id"]: json.loads(l) for l in open(a.runs + "/harness.jsonl")}
    why = {}
    for l in open(repo + "/data/why.jsonl"):
        r = json.loads(l)
        if r.get("id"):
            why[r["id"]] = r
    support = collections.defaultdict(list)
    for l in open(a.support):
        i, lib, f = l.rstrip("\n").split("\t")
        support[i].append((lib, f))
    trace = {}
    for r in csv.DictReader((l for l in open(a.trace) if not l.startswith("#")), delimiter="\t"):
        trace[r["id"]] = r
    fuzz = {}
    for r in csv.DictReader((l for l in open(repo + "/data/fuzz-trace.tsv") if not l.startswith("#")), delimiter="\t"):
        fuzz[r["id"]] = r
    depends = collections.defaultdict(list)
    dep_path = repo + "/results/depends-on.tsv"
    if os.path.exists(dep_path):
        for l in open(dep_path):
            if l.startswith("#") or l.startswith("id\t"):
                continue
            w = l.rstrip("\n").split("\t")
            depends[w[0]] = w[1].split(",")
    licence_col = read_licence_cols(a.baseline)
    locks = {}
    for l in open(repo + "/sources.lock"):
        if l.startswith("#"):
            continue
        n, url, commit = l.split()[:3]
        locks[n] = (url[:-4] if url.endswith(".git") else url, commit)

    def licence_of(c):
        """(SPDX, licence file under licences/, note) for one case."""
        g, i = c["Group"], c["ID"]
        spdx, file = parse_licence(licence_col.get(i))
        if SETS[g][2] or g == "mutants":
            return spdx or "UNKNOWN", file, None
        if g == "grammar-sentences":
            return "MIT", "tlaplus/LICENSE", "generated from tla+.jj, so read as following it"
        if g == "fuzz-jazzer":
            return "Apache-2.0", "LICENSE", "at set level: the set's LICENSE.md lists what the fuzzer was seeded with"
        return "Apache-2.0", "LICENSE", None

    env_head = dict(env)
    env_head["jar"]["file"] = "tla2tools.jar"
    env_head["tlaplus_commit"] = env["jar"].get("git_revision")
    counts = collections.defaultdict(lambda: collections.Counter())
    manifests = collections.defaultdict(list)
    lic_counts = collections.defaultdict(collections.Counter)
    lic_files = set()
    apache_changed = 0
    for c in sorted(cases, key=lambda x: x["ID"]):
        i, g = c["ID"], c["Group"]
        cat, title, up = SETS[g][:3]
        w = why[i]
        o, hr = run[i], h[i]
        place = w["Place"] if "Place" in w else w["place"]
        counts[g][place] += 1
        crash = V.crash_of(o.get("stdout", ""), o.get("stderr", "")) if o["status"] == "crash" else None
        expected = {"status": o["status"], "verdict": o["verdict"], "first_position": o.get("first_position"), "stage": o.get("stage"), "exit_plain": o.get("exit_plain"), "exit_error_codes": o.get("exit_error_codes")}
        if crash:
            expected["crash"] = crash
            expected["crash_note"] = "counted as rejection at this stage"
        for k in ("stdout", "stderr", "stdout_base64", "stderr_base64"):
            if k in o:
                expected[k] = o[k]
        harness = {k: hr.get(k) for k in ("status", "error_codes", "warning_codes", "syntax_tree_sha256", "semantic_graph_sha256", "semantic_graph_state")}
        harness["first_error_location"] = hr.get("first_error_location")
        harness = {k: v for k, v in harness.items() if v not in (None, "")}
        # Source.
        src = {}
        org = c["Origin"] or ""
        if c["Commit"] and org:
            repo_name, _, rest = org.partition("/")
            path, _, sub = rest.partition(" :: ")
            url, commit = locks[repo_name]
            src = {"repository": url, "commit": commit, "path": path}
            if sub:
                src["section"] = sub
            src["url"] = "%s/blob/%s/%s" % (url, commit, path)
        if g == "mutants":
            src.update(kind="mutant", parent=c["Parent"], rule=c["Rule"], seed=20261103, intent=c["Intent"])
        elif g == "fuzz-jazzer":
            src = {"kind": "jazzer", "strategy": c["Parent"]}
            if c.get("DerivedFrom"):
                src["derived_from"] = c["DerivedFrom"]
                src["change"] = "a local file path in the fuzz output was deleted; stock verdict, exit status, codes and coverage were measured unchanged"
            t = fuzz[i]
            if t["closest"] == "dictionary":
                src["trace"] = {"method": "dictionary", "from": "data/fuzz-dictionary.txt, generated from tla+.jj"}
        elif g == "grammar-sentences":
            src = {"kind": "grammar-sentence", "aim": c["Aim"]}
        elif not up:
            src = {"kind": "model" if cat == "ours-model" else "inspection"}
            if c["Model"]:
                src["model" if cat == "ours-model" else "author"] = c["Model"]
            if c["Target"]:
                src["target"] = c["Target"]
            aim = (c["Aim"] or "")
            if aim and not aim.startswith("no aim"):
                src["aim"] = aim
            if c["Parent"] or org.startswith(("reference/", "incoming/", "suite/")):
                src["written_as"] = c["Parent"] or org
        if not up and i in trace:
            t = trace[i]
            if t["best_parent"] and not (g == "fuzz-jazzer" and fuzz[i]["closest"] == "dictionary"):
                src["trace"] = {"closest_corpus_case": t["best_parent"], "longest_shared_run_bytes": int(t["longest_run"]), "covered_pct": float(t["covered_pct"]), "other_files_tied": int(t["other_files_tied"])}
        spdx, file, note = licence_of(c)
        if "corpus/external/paxosstore/" in org:
            src["wraps"] = {"repository": "https://github.com/Disalg-ICS-NJU/PaxosStore-TLA", "commit": "0188955e7023b4350943a9961d9b83a0d87e9fe9", "path": "experiment/TPaxos", "licence": "licences/PaxosStore-TLA/LICENSE", "note": "the modules this wrapper extends are not in this case, so SANY reports them missing; that is what is recorded"}
        d_case = None
        rel_dir = "%s/%s/%s" % ("sets" if place != "deep-cut" else "deep-cuts", cat, g)
        files = []
        src_dir = os.path.join(a.tree, "cases", i)
        for f in (c["Files"] or []):
            files.append({"path": f, "sha256": sha(os.path.join(src_dir, f)), "role": "root" if f == c["RootFile"] else "support"})
        for f in (c["Aux"] or []):
            files.append({"path": f, "sha256": sha(os.path.join(src_dir, f)), "role": "upstream-expected-tree"})
        for lib, f in sorted(support.get(i, [])):
            files.append({"path": f, "sha256": sha(os.path.join(a.tree, "libs", lib, f)), "role": "library", "library": lib})
        row = {"id": i, "root_file": c["RootFile"], "files": files, "source": src}
        if c.get("Unresolved"):
            row["unresolved_imports"] = c["Unresolved"]
        if c.get("Flags"):
            row["flags"] = c["Flags"]
        row["purpose"] = w.get("reason") or w.get("Reason")
        row["why"] = {k: w[k] for k in ("only_this_case", "beyond_upstream", "sole", "rarest", "keeps", "keeps_more") if w.get(k)}
        if c["Upstream"] or (c["Expect"] and not c["Expect"].startswith("nothing")):
            row["upstream_assertion"] = c["Expect"]
        row["licence"] = {"spdx": spdx, "file": ("licences/" + file) if file and file != "LICENSE" else (file or None)}
        if note:
            row["licence"]["note"] = note
        if file and file != "LICENSE":
            lic_files.add(file)
        lic_counts[(cat, g, place != "deep-cut")][(spdx, file)] += 1
        if depends.get(i):
            row["depends_on"] = depends[i]
        row["expected"] = expected
        row["harness"] = harness
        manifests[rel_dir].append(row)
        d_case = os.path.join(repo, rel_dir, "cases", i)
        os.makedirs(d_case)
        for f in (c["Files"] or []) + (c["Aux"] or []):
            shutil.copyfile(os.path.join(src_dir, f), os.path.join(d_case, f))
        for lib, f in support.get(i, []):
            os.symlink("../../../../../libs/%s/%s" % (lib, f), os.path.join(d_case, f))
        if spdx == "Apache-2.0" and file and file != "LICENSE" and not up:
            apache_changed += 1
            with open(os.path.join(d_case, "CHANGES.txt"), "w") as f:
                f.write(apache_notice(c, src))
    # Manifests.
    head = {
        "manifest": "SANY expected output",
        "copyright": "Copyright 2026 Christopher O'Connell",
        "spdx": "Apache-2.0",
        "reference": env_head,
        "expected_comes_from": "plain stock SANY: the command in reference.command in a directory holding the case, default JVM flags, no wrapper. expected.verdict is confirmed by the exit code of reference.verdict_command. Paths are mapped: $CASE is the case directory, $TMP java.io.tmpdir, $JAR the jar; nothing else changes.",
        "harness_column": "harness.* is not printed by the plain command, so it comes from cmd/sanyharness: tools/harness/SanyProbe.java over SANY's public API for codes and the syntax tree, and the unmodified tla2sany.xml.XMLExporter -o -t for the semantic graph, canonicalised by internal/canon. Default JVM flags except a private java.io.tmpdir.",
        "meaning": "What this SANY did on this input. It says nothing about whether SANY is right.",
    }
    for rel, rows in manifests.items():
        d = os.path.join(repo, rel)
        os.makedirs(d, exist_ok=True)
        h2 = dict(head)
        h2["set"] = rel
        with open(os.path.join(d, "manifest.jsonl"), "w", encoding="utf-8") as f:
            f.write(json.dumps(h2, sort_keys=True, ensure_ascii=False) + "\n")
            for row in rows:
                f.write(json.dumps(row, ensure_ascii=False, sort_keys=True) + "\n")
    # Licence files and libraries.
    for file in sorted(lic_files):
        repo_name, _, rest = file.partition("/")
        src = os.path.join(a.baseline, "external", repo_name, rest)
        dst = os.path.join(repo, "licences", file)
        os.makedirs(os.path.dirname(dst), exist_ok=True)
        shutil.copyfile(src, dst)
    lib_dst = os.path.join(repo, "libs")
    shutil.rmtree(lib_dst, ignore_errors=True)
    shutil.copytree(os.path.join(a.tree, "libs"), lib_dst)
    shutil.copyfile(os.path.join(a.baseline, "external", "CommunityModules", "LICENSE"), os.path.join(lib_dst, "CommunityModules", "LICENSE"))
    with open(os.path.join(lib_dst, "README.md"), "w") as f:
        f.write(HEADMD + "\n# Library modules\n\nModules that some cases import and that SANY does not carry: `CommunityModules` from tlaplus/CommunityModules (MIT) and `tlapm-library` from tlaplus/tlapm's `library` (BSD-2-Clause), at the commits in `sources.lock`, each with its licence. A case that parses one has a link to it in its own directory, so the plain command finds it there. `files` in each manifest row lists them with `role: library`. Note that `tlapm-library/FiniteSets.tla` is not SANY's own: a case linking it reads that copy.\n")
    # What wrappers name upstream: licensed there, not vendored by tlx, not added here.
    upstream_licence = os.path.join(repo, "licences", "PaxosStore-TLA", "LICENSE")
    os.makedirs(os.path.dirname(upstream_licence), exist_ok=True)
    shutil.copyfile(a.paxosstore + "/LICENSE", upstream_licence)
    # READMEs.
    for g, (cat, title, up, what, where, expects, regen) in SETS.items():
        write_set_readme(repo, g, cat, title, up, what, where, expects, regen, counts[g], lic_counts, manifests)
    with open(os.path.join(repo, "deep-cuts", "README.md"), "w") as f:
        f.write(HEADMD + "\n# Deep cuts\n\nCases of ours that the selection rule did not keep: each reaches nothing the selected and upstream cases do not, or only what a smaller case also reaches. They are kept, by source, as extra data for testing implementations, with the same manifest format as `sets/`. `docs/coverage.md` (\"Why each case is here\") says how they were chosen, and each row's `purpose` and `why` say why this one is here. `scripts/verify` runs them unless `--sets` says otherwise.\n")
    summary = {g: dict(counts[g]) for g in SETS}
    with open(repo + "/data/counts.json", "w") as f:
        json.dump({"copyright": "Copyright 2026 Christopher O'Connell", "spdx": "Apache-2.0", "per_set": summary, "apache_change_notices": apache_changed}, f, indent=1, sort_keys=True)
        f.write("\n")
    tot = collections.Counter()
    for g in SETS:
        tot.update(counts[g])
        print("%-32s %s" % (g, dict(counts[g])))
    print("total", dict(tot), "change notices", apache_changed)


def seed_note(repo):
    """What the Jazzer runs were fed, from the measurement in data/fuzz-seed-sources.json."""
    m = json.load(open(repo + "/data/fuzz-seed-sources.json"))
    names = {"LICENSE": "our own model specs"}
    per = collections.OrderedDict()
    for file, spdx, kind, n in m["seed_table"]:
        e = per.setdefault((file, spdx), {"file from": 0, "mutant of": 0, "model spec": 0})
        e[kind] += n
    text = "## What the fuzzer was fed\n\nMeasured from the directories the runs used, not assumed. `data/fuzz-seeds.tsv` lists every seed; `data/fuzz-seed-sources.json` holds these totals.\n\n**Seeds: %d distinct cases, chosen per strategy.** The licences below are the seeds' own, kept in `licences/` so that their notices stay with this repository.\n\n| from | licence | files | mutants of files there | notice |\n|---|---|---:|---:|---|\n" % m["distinct_seeds"]
    for (file, spdx), e in sorted(per.items()):
        if file == "LICENSE":
            text += "| %s | %s | %d | | `../../../LICENSE` |\n" % (names[file], spdx, e["model spec"])
        else:
            text += "| `%s` | %s | %d | %d | `../../../licences/%s` |\n" % (file.rsplit("/", 1)[0], spdx, e["file from"], e["mutant of"], file)
    o = m["other_inputs"]
    text += "\n**Also fed to it:** the dictionary `../../../data/fuzz-dictionary.txt`, generated from `tla+.jj` in tlaplus/tlaplus (MIT); %d earlier kept outputs (%d of this set, %d model specs, %d inspection); %d model replies that got near a target; and up to %d of its own earlier units. `CommunityModules` (MIT) and tlapm's `library` (BSD-2-Clause) were on the library path, and `../../../libs/` holds them with their licences.\n\n**Nothing else.** No seed comes from tlx's `corpus/external`, and none of those %d files mentions PaxosStore. Model specs are ours, written from prompts that carried corpus text.\n" % (o["kept-seeds"]["files"], m["kept_seeds_by_set"].get("fuzz-jazzer", 0), m["kept_seeds_by_set"].get("model-one-shot", 0), m["kept_seeds_by_set"].get("inspection", 0), o["model-near"]["files"], o["retained"]["files"], sum(v["files"] for v in o.values()))
    return text


def apache_notice(c, src):
    rule = c["Rule"] or "fuzzed and minimised"
    return ("Copyright 2026 Christopher O'Connell\nSPDX-License-Identifier: Apache-2.0\n\n"
            "Change notice, Apache License 2.0, section 4(b). The file in this directory is a\n"
            "modified copy of a file in tlaplus/Examples, specifications/byihive, which is under\n"
            "the Apache License 2.0; see licences/Examples/specifications/byihive/LICENSE.\n"
            "It was changed in November 2026 by Christopher O'Connell: %s. The exact bytes are\n"
            "under test, so the notice sits beside the file and not in it.\n" % rule)


def write_set_readme(repo, g, cat, title, up, what, where, expects, regen, n, lic_counts, manifests):
    d = os.path.join(repo, "sets", cat, g)
    os.makedirs(d, exist_ok=True)
    rel = "sets/%s/%s" % (cat, g)
    dc = "deep-cuts/%s/%s" % (cat, g)
    rows = manifests.get(rel, []) + manifests.get(dc, [])
    verd = collections.Counter((r["expected"]["status"], r["expected"]["verdict"]) for r in rows)
    text = HEADMD + "\n# %s\n\n" % title
    if up:
        text += "%d cases, kept whole: this set's reason for being here is that it is %s's.\n\n" % (n["upstream"], "upstream")
    else:
        text += "%d cases selected here, %d more in [`%s`](../../../%s/).%s\n\n" % (n["selected"], n["deep-cut"], dc, dc, "")
    text += "**What it is.** %s\n\n**Where it came from.** %s\n\n**What it expects.** %s Every row's `expected` is what plain stock SANY printed, recorded from the reference build; it says what SANY does and not whether that is right. See the top-level `README.md`.\n\n**How to regenerate.** %s\n\n" % (what, where, expects, regen)
    if rows:
        text += "**Recorded output.** %s.\n\n" % ", ".join("%d %s" % (v, ("crash, counted as reject" if k[0] == "crash" else ("timeout" if k[0] == "timeout" else k[1]))) for k, v in sorted(verd.items(), key=str))
    text += "Files: `manifest.jsonl`, one row per case, the first line the environment it was recorded in; `cases/<id>/`, the root module and its support modules, with library modules as links into `libs/`; `LICENSE.md`.\n"
    with open(os.path.join(d, "README.md"), "w") as f:
        f.write(text)
    # Licence note.
    lic = collections.Counter()
    for (c0, g0, sel), m in lic_counts.items():
        if g0 == g:
            lic.update(m)
    note = HEADMD + "\n# Licences of this set\n\nA reading of the files at the pinned commits. It is not legal advice; `LICENCES.md` at the top lists what needs a human look.\n\n| licence | file | cases, here and in deep-cuts |\n|---|---|---:|\n"
    for (spdx, file), v in sorted(lic.items(), key=str):
        where = "`../../../LICENSE`" if file == "LICENSE" else ("`../../../licences/%s`" % file if file else "none recorded")
        note += "| %s | %s | %d |\n" % (spdx, where, v)
    if g in ("inspection-pit-survivors", "model-pit-survivors"):
        note += "\nThe files say `All rights reserved` in their first lines. Those lines are part of the bytes under test, so they stay; the licence that applies is the one in this table, by the holder's decision for this repository.\n"
    if g == "tlx":
        note += "\nThe 4 `MC_tlx_*` wrappers in `corpus/external/paxosstore` are tlx's own, under tlx's licence. They name modules of `Disalg-ICS-NJU/PaxosStore-TLA`, MIT: `../../../licences/PaxosStore-TLA/LICENSE`. tlx's `corpus/external.json` records that upstream as having no licence; the record is stale.\n"
    if g == "mutants":
        note += "\nEach mutant follows its parent. Mutants of Apache-2.0 files carry a `CHANGES.txt` beside the file.\n"
    if g == "fuzz-jazzer":
        note = HEADMD + "\n# Licence of this set\n\nApache-2.0, at set level: `../../../LICENSE`. No case carries a parent's licence. A reading, not legal advice. `LICENCES.md` at the top says why.\n\n" + seed_note(repo)
    with open(os.path.join(d, "LICENSE.md"), "w") as f:
        f.write(note)
    if not up:
        dd = os.path.join(repo, dc)
        os.makedirs(dd, exist_ok=True)
        with open(os.path.join(dd, "README.md"), "w") as f:
            f.write(HEADMD + "\n# Deep cuts: %s\n\n%d cases not selected, kept as extra data for testing implementations. Each still has its expected output and its `why`. What this set is: [`%s/README.md`](../../../%s/README.md). Licences: [`%s/LICENSE.md`](../../../%s/LICENSE.md).\n" % (title, n["deep-cut"], rel, rel, rel, rel))
        if n["deep-cut"] == 0 and not manifests.get(dc):
            pass


if __name__ == "__main__":
    main()
