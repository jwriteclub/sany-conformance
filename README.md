<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# sany-conformance

A test corpus for SANY, the parser and semantic checker in the TLA+ tools. It holds 9,453 TLA+ modules. Each one comes with the exact output stock SANY gave on it, so another implementation or a later build can be compared with it case by case. The repository also has the scripts that run those comparisons.

**This is not a correctness suite.** The cases aren't chosen because they're right or wrong TLA+. Each case is here because it exercises branches in the Java implementation. The core suite is assembled from publicly available TLA+ files. In addition, a number of cases were generated through fuzzing and agent-guided mutation. Together they're the library of inputs you need to use the SANY jar as an oracle: an implementation that matches SANY on all of these should be substantially accurate.

## What's in the tin

The cases come from three places:

- **4,130 modules from the TLA+ project and other open-source projects** (tlapm, tlx, tree-sitter-tlaplus). All of them are in `sets/`, unchanged, duplicates included.
- **1,553 modules we generated**: mutants of upstream specs, fuzzer output, short inputs generated from SANY's grammar, and specs that language models wrote, aimed at specific parts of SANY's code. These all cause the Java code to execute branches that it does not with the other examples.
- **3,770 deep cuts**: These are in `deep-cuts/` as extra test data. They were all generated in the same way as the other generated cases. These are kept in deep cuts, since some of them are useful for finely exercising very specific code paths, but one of the other sources also executes these paths, so these are not strictly necessary for comparison.

The expected output for each case comes from an unmodified build of tlaplus commit `dc6a944`, run as `java -cp tla2tools.jar tla2sany.SANY <file>` with default JVM flags on Linux. That commit is newer than v1.7.4, the latest release, and v1.7.4 gives a different verdict or status on 468 cases (see [The stock check](docs/reference.md#the-stock-check)). A second run gave the same output on every case except one stack trace.

## What's not in the tin

Despite our best efforts (and billions of tokens spent on guided fuzzing) some paths in SANY simply could not be reliably triggered by inputs. LLM agents have plausibly inferred that some of them are not reachable, _however_ there remain about 1,000 branch outcomes in SANY (510 in the handwritten code and 523 in the generated parser and lexer, within the `port` scope; see [docs/coverage.md](docs/coverage.md)) that are not covered in this set _and_ for which it wasn't possible to plausibly mark them unreachable. As a practical matter, however, none of these branches are triggered on any public TLA+ file that I could find, so for editors and similar tools, the cases here represent a very good starting point for behavior.

## Documentation

- [docs/compare.md](docs/compare.md): comparing an implementation with SANY, checking a SANY build, reading the report, and verifying.
- [docs/results-prudence.md](docs/results-prudence.md): the published results for prudence.
- [docs/results-manyan.md](docs/results-manyan.md): results for [manyan](https://github.com/soasme/manyan), front end against SANY and checking against TLC.
- [docs/coverage.md](docs/coverage.md): what the corpus covers, why each case is here, coverage kept, what coverage means, and limits.
- [docs/reference.md](docs/reference.md): the sets, the manifest format, the stock check, and the repository layout.
- [FINDINGS.md](FINDINGS.md): oddities in SANY that this work turned up, with reproducers.
- [LICENCES.md](LICENCES.md): how each kind of case is licensed, and what needs a human look.

## Quick start

Needs Go and a clean prudence checkout at the commit in `cmd/compare/prudence.pin` (set `PRUDENCE_DIR` if it is not at `../prudence`):

```sh
scripts/compare.sh --set tlaplus-syntax-corpus --limit 100                    # 100 cases from one set
scripts/compare.sh --jobs 4                                                   # every case in sets/ (5,683)
```

Options, the report format and checking a SANY build are in [docs/compare.md](docs/compare.md).

## Licence and provenance

Our code and our cases are Apache-2.0 (`LICENSE`). Each upstream case keeps its own licence, and every manifest row names it. `LICENCES.md` explains how each kind of case is licensed and lists what still needs a human look. This is not legal advice.

Upstream repositories and commits are pinned in `sources.lock`. Each manifest row's `source` field says where its case came from: repository, commit and path for upstream cases; generator, parent, model and target for ours.

`FINDINGS.md` lists oddities in SANY that this work turned up, each with a way to reproduce it.
