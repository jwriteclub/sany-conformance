<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# Reference

## The sets

`casemap sources` prints this table. Each cell is "reached by this set / reached only by this set", for branch outcomes and grammar rules. `port` is defined under [What coverage means](coverage.md#what-coverage-means).

| set | cases | selected | deep cuts | distinct | `port` handwritten | `port` generated | `port` rules | overall handwritten | overall generated | overall rules |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| **From the TLA+ project** | | | | | | | | | | |
| `tlaplus-syntax-corpus` | 355 | 355 | 0 | 355 | 1495 / 0 | 3566 / 0 | 623 / 0 | 1675 / 0 | 3720 / 0 | 623 / 0 |
| `tlaplus-semantic-corpus` | 27 | 27 | 0 | 27 | 1717 / 0 | 2452 / 0 | 406 / 0 | 1910 / 0 | 2597 / 0 | 406 / 0 |
| `tlaplus-semantic-error-corpus` | 121 | 121 | 0 | 121 | 1879 / 0 | 2035 / 0 | 312 / 0 | 1994 / 0 | 2146 / 0 | 312 / 0 |
| `tlaplus-inline-java` | 52 | 52 | 0 | 52 | 715 / 0 | 1104 / 0 | 94 / 0 | 819 / 0 | 1135 / 0 | 94 / 0 |
| `tlaplus-test-models` | 670 | 670 | 0 | 668 | 2584 / 0 | 3538 / 0 | 547 / 0 | 2814 / 0 | 3702 / 0 | 547 / 0 |
| `tlaplus-other` | 296 | 296 | 0 | 294 | 2452 / 0 | 3417 / 4 | 525 / 0 | 2655 / 0 | 3578 / 4 | 525 / 0 |
| `tlaplus-examples` | 433 | 433 | 0 | 407 | 2347 / 0 | 2972 / 0 | 463 / 0 | 2558 / 0 | 3131 / 0 | 463 / 0 |
| `tlaplus-communitymodules` | 80 | 80 | 0 | 80 | 1901 / 0 | 2582 / 0 | 397 / 0 | 2116 / 0 | 2726 / 0 | 397 / 0 |
| `tlaplus-rfcs` | 475 | 475 | 0 | 473 | 2305 / 0 | 3759 / 0 | 641 / 0 | 2518 / 0 | 3924 / 0 | 641 / 0 |
| **From other projects** | | | | | | | | | | |
| `tlapm` | 677 | 677 | 0 | 676 | 2498 / 0 | 3604 / 0 | 583 / 0 | 2713 / 0 | 3775 / 0 | 583 / 0 |
| `tlx` | 587 | 587 | 0 | 574 | 2219 / 0 | 2781 / 0 | 439 / 0 | 2425 / 0 | 2938 / 0 | 439 / 0 |
| `tree-sitter-tlaplus` | 357 | 357 | 0 | 353 | 1506 / 0 | 3693 / 0 | 623 / 0 | 1686 / 0 | 3847 / 0 | 623 / 0 |
| **Ours, generated mechanically** | | | | | | | | | | |
| `mutants` | 3000 | 63 | 2937 | 3000 | 2638 / 0 | 4111 / 28 | 649 / 8 | 2849 / 0 | 4282 / 28 | 649 / 8 |
| `fuzz-jazzer` | 1235 | 1132 | 103 | 1235 | 2729 / 17 | 4799 / 612 | 632 / 7 | 2830 / 18 | 4975 / 613 | 632 / 7 |
| `grammar-sentences` | 388 | 15 | 373 | 388 | 1156 / 0 | 3087 / 0 | 597 / 0 | 1308 / 0 | 3223 / 0 | 597 / 0 |
| **Ours, model-guided** | | | | | | | | | | |
| `model-one-shot` | 341 | 163 | 178 | 341 | 2730 / 45 | 3011 / 13 | 462 / 2 | 2933 / 52 | 3169 / 13 | 462 / 2 |
| `model-agents` | 41 | 19 | 22 | 41 | 1708 / 0 | 2395 / 0 | 362 / 0 | 1826 / 0 | 2527 / 0 | 362 / 0 |
| `model-scoring` | 8 | 5 | 3 | 8 | 992 / 0 | 1500 / 0 | 195 / 0 | 1000 / 0 | 1577 / 0 | 195 / 0 |
| `model-ternary-hunt` | 0 | 0 | 0 | 0 | 0 / 0 | 0 / 0 | 0 / 0 | 0 / 0 | 0 / 0 | 0 / 0 |
| `model-openstart` | 23 | 23 | 0 | 23 | 796 / 0 | 1144 / 0 | 121 / 0 | 929 / 0 | 1180 / 0 | 121 / 0 |
| `model-pit-survivors` | 27 | 23 | 4 | 27 | 1735 / 0 | 1766 / 0 | 227 / 0 | 1890 / 0 | 1841 / 0 | 227 / 0 |
| **Ours, by inspection** | | | | | | | | | | |
| `inspection` | 188 | 42 | 146 | 188 | 2104 / 3 | 2863 / 17 | 396 / 0 | 2301 / 4 | 3003 / 17 | 396 / 0 |
| `inspection-pit-survivors` | 72 | 68 | 4 | 72 | 2280 / 5 | 2461 / 2 | 350 / 0 | 2451 / 5 | 2582 / 2 | 350 / 0 |
| **whole suite, reached of total** | 9453 | | | 8350 | 3063 of 3573 | 5093 of 6251 | 685 of 737 | 3306 of 5010 | 5270 of 6564 | 685 of 747 |

By kind of source. Each cell is "reached by any set of this kind / reached by no other kind".

| kind of source | cases | `port` handwritten | `port` generated | `port` rules | overall handwritten | overall generated |
|---|---:|---:|---:|---:|---:|---:|
| From the TLA+ project | 2509 | 2801 / 11 | 4056 / 4 | 645 / 0 | 3031 / 11 | 4227 / 4 |
| From other projects | 1621 | 2579 / 0 | 4015 / 0 | 645 / 0 | 2810 / 0 | 4186 / 0 |
| Ours, generated mechanically | 4623 | 2894 / 17 | 5029 / 706 | 681 / 21 | 3124 / 18 | 5206 / 708 |
| Ours, model-guided | 440 | 2802 / 64 | 3195 / 13 | 490 / 2 | 3010 / 72 | 3361 / 13 |
| Ours, by inspection | 260 | 2543 / 11 | 3072 / 27 | 438 / 0 | 2745 / 12 | 3225 / 27 |

Each set directory has a `README.md` (what the set is, where it came from, what it expects and how to regenerate it), a `LICENSE.md`, a `manifest.jsonl`, and `cases/<id>/` with the root module and its support modules. `deep-cuts/` has the same layout. "By inspection" means specs that review agents wrote after reading SANY's code.

`model-ternary-hunt` is empty on purpose: see `FINDINGS.md`.

## What a manifest row holds

The first line of each `manifest.jsonl` records the environment: `reference.jar` (version and SHA-256), `reference.tlaplus_commit`, JVM vendor and exact version, OS and architecture, default charset, whether the file system is case-sensitive, and the exact commands. Each later line is one case:

| field | |
|---|---|
| `id`, `root_file` | which module to give SANY |
| `files` | path, SHA-256 and role: `root`, `support`, `library` (a link into `libs/`), or `upstream-expected-tree` |
| `source` | repository, commit, path and URL; or generator, parent, model, target and aim. Model and agent names record where a case came from. Some paths name the working tree a case was written in, not this repository |
| `upstream_assertion` | what upstream's own test says, in words, if anything |
| `purpose`, `why` | why this case is here. See [Why each case is here](coverage.md#why-each-case-is-here) |
| `licence` | SPDX identifier and file |
| `expected.status` | `ok`; `crash`, with `expected.crash` giving the exception, the frame, and whether it escaped or SANY logged it; or `timeout` at 300 s. None timed out |
| `expected.verdict` | `accept` or `reject`. A crash counts as a rejection at that stage, and `crash_note` says so. Read from SANY's text and confirmed by the `-error-codes` exit code |
| `expected.first_position` | `module:line:col`, read from SANY's text. `module:0:0` means it names a module but no position; `:0:0` means neither; `module:line:0` means a line only |
| `expected.stdout`, `stderr` | the full text, or `*_base64` if it is not valid UTF-8. Paths are mapped to `$CASE`, `$TMP` and `$JAR`; `tlc-<random>` becomes `tlc-N`, and nested modules' `<random>/` becomes `N/`. Nothing else is changed |
| `expected.exit_plain`, `exit_error_codes` | exit status of each command |
| `harness` | codes and hashes the plain command does not print. See below |
| `depends_on` | what moved this case's result, if anything: `results/depends-on.tsv` |

**The harness column.** The plain command does not print error or warning codes, the syntax tree or the semantic graph, and it exits 0 on a semantic rejection. The `harness.*` fields fill that gap. They are labelled separately because they do use a wrapper:

- `cmd/sanyharness` runs `tools/harness/SanyProbe.java` over SANY's public API for `error_codes`, `warning_codes`, `first_error_location` and `syntax_tree_sha256`.
- On accepted specs only, the unmodified `tla2sany.xml.XMLExporter -o -t` gives `semantic_graph_sha256`. Before hashing, `internal/canon` canonicalises the graph. It removes three things and nothing else: the numeric `<UID>` values, which it relabels by first reference from the module list; the order of `<context>` entries, which come from a `HashMap`; and differences in XML serialisation.
- `semantic_graph_state` says `not-exported:rejected`, `export-failed:exit=N`, `crash:...` or the graph's size.
- JVM flags are the defaults, except a private `java.io.tmpdir`.

Earlier results came through the same wrapper with `-XX:TieredStopAtLevel=1`. Against them there are 2 differences, both explained under [The stock check](#the-stock-check). 11 accepted specs (upstream Farsite modules and mutants of them) have no graph hash, because `XMLExporter` runs out of memory on them at the default heap of 15.7 GiB.

## The stock check

Everything below ran on all 9,453 cases. `results/` has every difference, per case.

**Plain stock SANY matches the earlier wrapper on every verdict and code. Only one case's output moved with anything other than charset or platform.** The earlier results came through a wrapper with `-XX:TieredStopAtLevel=1`, so some differences were expected. Over the 9,424 cases with an earlier verdict and codes on record, and every position, tree hash and graph hash of the 7,130 upstream and mutant cases, there are 2 differences, and neither comes from the flags or the wrapper. `B-test-model-bug156-te-9a905e` was recorded as a hang under load during an earlier index pass, and takes 14 s. `D-5ac9249fe910` was recorded on the Mac.

Expected output is from **Linux**, which has a case-sensitive file system and is what most CI runs on. Verdicts: 4,106 accept, 5,319 reject, 28 crash, 0 timeouts at 300 s. The plain text's verdict agrees with the `-error-codes` exit code on all 9,453.

| compared with the reference | cases that differ | what moved |
|---|---:|---|
| a second run, in reverse order | 1 | the stack trace text of `E-550917c970f7`, the one `StackOverflowError`, and nothing else |
| macOS 15.5 arm64, APFS, same jar | 2 | that, and `D-5ac9249fe910`: it imports `FiniteSetS`, APFS finds the library's `FiniteSets.tla`, so SANY reports a name mismatch instead of a missing file |
| `-XX:TieredStopAtLevel=1` | 1 | that stack trace |
| `-Xss512k`; `-Xss16m` | 1; 1 | that stack trace. It overflows at both. Nothing else here is sensitive to stack depth |
| always recomputing `jj_ntk` | 1 | that stack trace |
| `LC_ALL=C`, charset `ANSI_X3.4-1968` | 145 | that, 113 messages quoting a non-ASCII character as `?`, and 31 cases that change their first position, all with the same verdict: a nested module is rewritten to a temporary file in the default charset, so `≃` becomes `?` and a lexical error |
| `-Dfile.encoding=ISO-8859-1` | 144 | that, and 143 messages. Same verdicts |
| v1.8.0 pre-release, 7 commits past `dc6a944`, none touching `tla2sany` | 1 | that stack trace |
| **v1.7.4, GitHub's latest release, from 2024** | all | see below |

In these runs, no case's verdict depends on platform, JVM flags, stack depth or charset. `D-f82fa4953e79`, which earlier runs showed differing with `InTegers`, does not differ with a jar. The earlier difference came from running on a class directory, where APFS also matches the standard modules' names. Of the 196 fuzz outputs that are not valid UTF-8, 176 moved with nothing here, and 20 moved only with the charset, keeping their verdict.

**v1.7.4 against `dc6a944`.** v1.7.4 prints an older banner on every case. Apart from progress lines, messages differ on 3,909 cases. Verdict or status differs on 468:

- 287 accepted now and rejected by v1.7.4, because it cannot find a module nested in the same file, a later feature;
- 140 where v1.7.4 gives a lexical error on Unicode syntax;
- 7 parse errors and 4 others;
- 28 that v1.7.4 crashes on: 16 with `NullPointerException`, 8 with `NumberFormatException`, 4 with `ArrayIndexOutOfBoundsException`;
- 2 that v1.7.4 accepts and that are rejected now, among them `A-fix-sany-github723`.

`results/diff-rel174.tsv.gz` has each one. These are differences between versions, not failures.

## Layout

```
sets/<kind>/<set>/          README.md LICENSE.md manifest.jsonl cases/<id>/
deep-cuts/<kind>/<set>/     README.md manifest.jsonl cases/<id>/
libs/                       library modules, with their licences; cases link into here
licences/                   every licence file a row points to
results/                    the stock check: every difference, per case
data/                       coverage index, why, traces, static map
FINDINGS.md                 oddities in SANY, with reproducers
LICENCES.md                 what licence each kind of case carries, and what needs a human look
scripts/                    compare.sh, verify, setup, regenerate
cmd/, internal/, tools/     compare, casemap, staticreach, sanyharness, extract, gensentences, tracecases, the grammar trace
notes/                      working notes from building the compare script
```

`deep-cuts/` is 29 MB of plain files, mostly mutants of specs up to 30 KiB. It compresses well in git: the whole `.git` is about 20 MB. Nothing was done to shrink it beyond linking library modules, which would otherwise add 44 MB across the tree.
