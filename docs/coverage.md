<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# Coverage

## What the corpus covers

The coverage numbers here are for SANY's own Java code, measured with JaCoCo while running the cases. They say how much of SANY the corpus exercises. They are not coverage of prudence or of any other implementation.

The `port` scope is the part of SANY the command line can reach. It leaves out the `-d` Explorer, `XMLExporter`, `SymbolTable.toString` and the `jj_ntk` memo check (see [What coverage means](#what-coverage-means)).

| | reached |
|---|---:|
| handwritten branch outcomes in `port` | 3,063 of 3,573 (85.7%) |
| generated parser and lexer branch outcomes in `port` | 5,093 of 6,251 (81.5%) |
| grammar productions entered | 87 of 88 |
| PIT mutants detected, handwritten `port` code | 80.5% of all 7,160; 87.4% of those some case covers |

Coverage shows that code ran. It does not show that any output was checked. Mutation testing (PIT) asks the second question: whether breaking one instruction changes the output.

What the corpus does not cover:

- **Other SANY builds, JVMs and machines.** Expected output is from Temurin 17.0.21+9 on Linux x86_64. `results/` shows what moves on macOS, with other JVM flags and charsets, and with other releases. Other JVM vendors and major versions were not run.
- **Code the cases do not reach.** A branch no case reaches is not shown to be unreachable; other inputs may reach it.
- **Anything SANY's plain output does not show**, except through the harness column of each manifest row (codes, tree and graph hashes; see [What a manifest row holds](reference.md#what-a-manifest-row-holds)).

[Limits](#limits) below has the full list.

## Why each case is here

An upstream set is here because it is upstream's, so upstream sets are kept whole, duplicates included. Each case of ours has to state why it exists, and that reason must not depend on the order in which cases were added.

`casemap select` computes the reasons. An element is anything a case can be the reason for: a branch outcome, line, method or grammar element from the coverage index, a token kind, an error-code combination (verdict, error codes, warning codes), a SANY-vs-tlx disagreement signature, a PIT mutant kill, or a Java crash signature. Each case's `why` says how many elements only it reaches and how many no upstream set reaches, and names the rarest. For a selected case, `keeps` names what no other selected or upstream case reaches. These are facts about the final set, in no particular order.

**Selection.** Plain greedy set cover picks the case with the most new coverage first, which tends to keep large kitchen-sink cases. This selection works the other way. For each element no upstream set reaches, rarest first, it keeps the smallest case that reaches it, by bytes in its own files. Ties go to the case that reaches least beyond upstream, then to the id. Then it drops cases that turned out redundant, largest first. Everything else goes to `deep-cuts/`; nothing is discarded. The rule's tests are in `cmd/casemap/select_test.go`.

**Evidence cases** keep a stated purpose, because coverage cannot credit them. There are 86:

- the 23 OpenStart modules;
- 13 specs that hit a lookahead outcome and 16 that hit a scan with its predicate false, which JaCoCo cannot record because the probe follows a call that throws;
- 1 that reaches comment joining only under `XMLExporter -u`;
- the 2 case-mismatch imports;
- 31 more whose result moves with the environment.

Mutant kills are the one element whose source is not free of order. PIT stops at the first case that kills a mutant, so most mutants name one killer, and which case that is depended on the order of the run. Selection keeps each recorded killer; it does not show that a deep cut would not also kill the mutant. 20 kills credited to specs that are not kept are left out.

## Coverage kept

`casemap prove` reruns the accounting from the index and exits non-zero if anything is lost. It also recomputes the selection with the cases in reversed and in shuffled order and requires the same answer, and it requires `data/why.jsonl` to match.

| reached | whole suite | upstream + selected |
|---|---:|---:|
| port handwritten | 3063 | 3063 |
| port generated | 5093 | 5093 |
| port rules | 685 | 685 |
| overall handwritten | 3306 | 3306 |
| overall generated | 5270 | 5270 |
| overall rules | 685 | 685 |
| lines | 14401 | 14401 |
| methods | 1315 | 1315 |
| lookahead outcomes | 170 | 170 |
| token kinds | 270 | 270 |
| error-code combinations | 1117 | 1117 |
| tlx disagreements | 55 | 55 |
| mutant kills, primary | 7418 | 7418 |
| mutant kills, wide | 25 | 25 |
| Java crash signatures | 12 | 12 |

Upstream plus selected cases reach everything the whole suite reaches, and all of it is in this tree.

## What coverage means

**Coverage says the code ran. It does not say any output was checked.** Output is checked only by comparing with a manifest, and a manifest records what SANY does.

Coverage was measured with JaCoCo 0.8.15, per case, in fresh JVMs, under a wrapper over SANY's public API (plus `XMLExporter` on accepted specs), with `-XX:TieredStopAtLevel=1`. Those runs, and the code that built the index from them, are not in this repository; `data/` holds the result. `casemap` turns JaCoCo's probes into branch outcomes by replaying JaCoCo's own per-probe flow. It maps a production's alternatives onto class-file arms only where tokens and calls match exactly: 943 elements map, and 0 fail to.

**The `port` scope** is an upper bound on what the command line can run. It is static reachability from `tla2sany.SANY.main` by rapid type analysis, with the `-d` Explorer, `XMLExporter` and `SymbolTable.toString` cut, and the `jj_ntk` ternary left out. A branch in `port` has not been ruled out, but that does not mean some input can take it. `data/staticreach/candidate-map.md` is the full report, and `cmd/staticreach` rebuilds it. Rebuilt from the jar, its branch table is byte-identical to the original. Its method table has one extra row, because the jar holds test classes that the measured class directory did not.

From `casemap grammar` and `casemap sources`, by JaCoCo:

| | reached |
|---|---:|
| handwritten branch outcomes in `port` | 3,063 of 3,573 (85.7%). The static map's own count, with a plain-command pass added, is 3,082 (86.3%), and 3,163 (88.5%) with upstream's unit tests |
| generated branch outcomes in `port` | 5,093 of 6,251 (81.5%) |
| grammar rules in `port` | 685 of 737 |
| productions entered | 87 of 88; the other is `OpenStart`, which only ever runs as lookahead |
| alternatives taken | 391 of 399; the other 8 are `OpenStart`'s |
| optionals and loops, both ways | 91 of 91 |
| syntax-error arms | 25 of 78. A patched parser sees 42 run, since a probe after a throw never fires; 27 are shown unreachable. See `FINDINGS.md` |
| `LOOKAHEAD` decisions and checks both ways | 72 of 98. A patched parser sees 68 of the 74 scans go both ways; the other 6 cannot |
| token kinds | 270 of 295; the rest cannot be returned as tokens |

**Mutation testing** asks whether a wrong answer would be noticed. PIT 1.30.1 changed one instruction at a time in the `port` scope's handwritten code (7,160 mutants) and ran the cases that cover each one against a recorded baseline of outcome, verdict, codes, positions, messages, syntax tree, semantic graph and level data. On macOS arm64, Temurin 17, with a time budget per mutant:

| | mutants | detected of all | detected of those some case covers |
|---|---:|---:|---:|
| handwritten, main run | 7,160 | 80.5% | 87.4% |
| handwritten, with 72 specs written for what survived (`inspection-pit-survivors`) | 7,160 | 83.3% | 90.5% |
| parser action code from `tla+.jj`, main run | 1,570 | 90.0% | 91.3% |
| parser action code, with those specs | 1,570 | 93.2% | 94.5% |

Of the 963 mutants that survived the main run, 677 still survived a wider second look on Linux. Reading each one marked 324 as dead or guarded, 264 as equivalent-looking, and 79 as visible only through something not compared. 10 were read as real gaps that nothing kills yet. Those marks are judgements; 14 marked dead or guarded were later killed by a spec. The PIT run compared more than the plain command prints, so a kill there is not necessarily a kill by this repository's manifests alone. The manifests record the plain command's output plus the harness column, with no level data.

## Limits

- **Expected output is from one build, one JVM, one machine:** Temurin 17.0.21+9, Linux x86_64, ext4, `en_US.UTF-8`. `results/` shows what moves. Other JVM vendors and major versions were not run.
- **Library modules** are linked into each case that parses them, so the plain command needs no `TLA-Library`. The earlier harness passed them with that property, and the results agree (`results/diff-harness.tsv`). A case that imports an unavailable module records that failure, as `B-test-model-bug156-te-9a905e` and tlx's PaxosStore wrappers do, since those are upstream files without their siblings. 173 rows list such imports in `unresolved_imports`.
- **Messages include stack traces and the banner.** A JVM change can move a frame in `stderr`. `verify --soft` ignores progress lines and runs of blank lines when comparing.
- **The stack trace of `E-550917c970f7`,** the one `StackOverflowError`, varies from run to run, so `verify` does not compare that case's `stderr`.
- **The coverage index comes from a different setup** than the expected output: the wrapper with `-XX:TieredStopAtLevel=1`, not the plain command. `results/diff-harness.tsv` shows that their verdicts agree. The index also has no data for the plain command's argument handling.
- **JaCoCo has a blind spot.** A probe does not fire when a call throws. Lookahead routines end that way, and so do syntax-error arms, so an arm that ran can show as not reached.
- **Other inputs can reach what these do not.** A branch no case reaches is not shown to be unreachable.
- **Fuzz traces and licence readings are evidence and judgement**, not proof. See `LICENCES.md`.
- **The maintainer's scripts** in `scripts/maintainer/` built this tree from the measurement outputs, which are not here. They are kept to show how it was done; they will not run without those outputs.
