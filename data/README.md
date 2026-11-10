<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# Data

What the tools read. None of it is needed to run `scripts/verify`.

| file | what | made by |
|---|---|---|
| `casemap-index.gob.gz` | every case's probes, and what each probe covers: branch outcomes, lines, methods, grammar elements, token kinds. `cmd/casemap` queries it | per-case JaCoCo 0.8.15 runs under a wrapper over SANY's public API, plus XMLExporter on accepted specs. See [docs/coverage.md, "Limits"](../docs/coverage.md#limits): these runs are not in this repository |
| `elements-extra.tsv` | what else a case can be the reason for: error-code combinations, tlx disagreements, PIT mutant kills, Java crash signatures | `scripts/maintainer/build_data.py`, from the stock runs and the PIT results |
| `mutant-kills.tsv.gz` | every recorded PIT kill: mutant, case, run | PIT 1.30.1. PIT stops at the first case that kills a mutant, so most mutants name one killer, and that depends on the order of the run |
| `evidence.tsv` | cases kept for a stated purpose coverage cannot measure | `build_data.py` |
| `fuzz-seeds.tsv`, `fuzz-seed-sources.json` | every case the Jazzer runs were seeded with, and what else they were fed | `scripts/maintainer/fuzz_seeds.py`, from the directories the runs used |
| `fuzz-trace.tsv`, `fuzz-dictionary.txt` | the file closest to each fuzz output by content, as provenance and not a licence claim; the fuzzer's dictionary | `cmd/tracecases`, `build_data.py`; the dictionary is generated from `tla+.jj` |
| `why.jsonl` | every case's place and reason. Regenerate with `casemap select`; `casemap prove` checks it | `cmd/casemap` |
| `grammar-guards.tsv.gz` | every syntax-error arm not taken: what can get to it, and which of those tokens has no case label | `tools/grammar-trace/guards.py` |
| `staticreach/` | the JaCoCo reports `cmd/staticreach` reads, and the static map it wrote | `cmd/staticreach` |
