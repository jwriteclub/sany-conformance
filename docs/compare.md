<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# Comparing and verifying

## Compare an implementation

`scripts/compare.sh` runs each case through prudence (`prudence.Check`), compares the answer with SANY's recorded output, and writes a report. It needs Go and a clean prudence checkout at the commit in `cmd/compare/prudence.pin`. Set `PRUDENCE_DIR` if the checkout is not at `../prudence`. The script refuses any other prudence commit, and a checkout with local changes. Nothing is downloaded.

```sh
scripts/compare.sh --set tlaplus-syntax-corpus --limit 100                    # 100 cases from one set
scripts/compare.sh --jobs 4                                                   # every case in sets/ (5,683)
scripts/compare.sh --deep-cuts --jobs 4                                       # sets/ and deep-cuts/ (9,453)
scripts/compare.sh --java --jar .work/jars/dc6a944/tla2tools.jar --limit 20   # run stock SANY again too
```

By default the SANY side is the output recorded in each `manifest.jsonl`. With `--java --jar J`, the script runs stock SANY again (`java -cp J tla2sany.SANY <root>`) in a private copy of each case; `scripts/setup reference` builds that jar, and `--javabin` picks the `java` binary. Other options:

- `--set NAME`: one set.
- `--limit N`: the first N cases of each set, in manifest order. This is not a random sample.
- `--ids FILE`: only the case ids in FILE, one per line.
- `--stage S`: `syntax` or `semantic`. The default is prudence's full check, semantic and level.
- `--jobs N`, `--out DIR` (default `compare-out`), `--pull` (git pull first).
- `--timeout DUR`: per-case timeout for each side (default `300s`).
- `--info`: keep prudence's non-error diagnostics, such as the reason it stopped, in `results.jsonl`.
- `--no-3d`: turn normalisation rule 3d off, so a nested-module error stays where rule 3c puts it. The [published numbers](results-prudence.md) use it.

The script exits nonzero only when the harness itself fails. Cases that differ do not change the exit status.

`cmd/compare` is written for prudence. The calls into prudence, and the conversion of its answers into the form the scorer compares, are in `cmd/compare/prudence.go`; another implementation would need its own version of that file.

The [published prudence numbers](results-prudence.md) were run with these options. That run called the compare binary directly, with the same flags, which is what this command does:

```sh
scripts/compare.sh --deep-cuts --jobs 32 --timeout 120s --info --no-3d
```

## Check a SANY build

`scripts/verify` runs a `tla2tools.jar` over every case and compares its output with the recorded output. You need Python 3.8 or later, a JVM and a jar; `verify` itself needs no network and no other tool. [How to verify](#how-to-verify) below has the details.

```sh
scripts/setup reference    # once, needs network, git and Ant: builds tlaplus at the pinned commit
scripts/verify --jar .work/jars/dc6a944/tla2tools.jar --java /path/to/java --jobs 4
```

## Reading the report

`compare.sh` writes two files to `--out`:

- `report.md`: what was compared, a summary table per set, then each mismatch with both sides.
- `results.jsonl`: one row per case, with both sides as the scorer saw them, prudence's completion and level, and the result.

Each case gets one result:

| result | meaning |
|---|---|
| match | same class (accepted, parse error, semantic error, missing module, module cycle), same number of errors, and each error has the same line, column, module and message text |
| match (order-normalised) | a match except that a 4332 "declares extra parameter(s)" list names the same parameters in a different order (rule 8). Counted on its own, never as a match |
| mismatch | anything else. The kind says how: `class`, `position` (same class, different places) or `text` (same places, different words) |
| prudence incomplete | prudence stopped before deciding, and SANY did not reject the case at parsing |
| oracle unavailable | SANY crashed on the case, so there is no verdict to compare with. Never scored as a match |
| errors, timeouts | prudence returned an error, or ran out of time |
| false complete | a separate count: prudence says it finished and found no error, but SANY reported one |

Only what SANY's plain output shows is compared: accept or reject, the stage, and each error's module, line, column and text. The report's header lists every normalisation, and that list is complete. A summary:

1. SANY's parse output is split into one error per record its parser logged.
2. SANY reprints its accumulated semantic errors after each module; repeats are dropped. Warnings are not compared.
3. Prudence's positions are converted to SANY's: line numbering per module, tab stops of eight, columns in UTF-16 units, and the module SANY would name (rules 3a to 3e).
4. Runs of whitespace in a message become one space.
5. A SANY error with no position (a missing module, for example) is matched by its text.
6. A SANY crash is `oracle-unavailable`.
7. Nothing else is changed, apart from rule 8.
8. The one exception: SANY prints the names in a 4332 extra-parameter list by iterating a `HashSet`, so their order is arbitrary. Two such lists with the same distinct names in a different order count as match (order-normalised).

**Not compared:** the syntax tree, the semantic graph and UIDs, error and warning codes, end positions, and anything on stderr. A match means the two agree on accept or reject, and on where the errors are and what they say. It does not mean their internal structures agree.

**Overrides.** `cmd/compare/overrides.go` holds a recorded table that can raise prudence's module-acquisition visit budget for a named case. It has one entry: TE.tla (`tlaplus-test-models/B-test-model-bug156-te-9a905e`) at 3,000,000. That case is scored normally, and its row and a report footnote show the value.

The scoring rule itself is `Compare` in `cmd/compare/compare.go`. Its doc comment states the rule in full.

## How to verify

You need Python 3.8 or later, a JVM and a `tla2tools.jar`. `scripts/verify` needs no network and no other tool; `scripts/setup` needs network once, to fetch or build the jars.

```sh
scripts/setup release      # once, needs network: v1.7.4 and v1.8.0, checked against pinned hashes
scripts/setup reference    # once, needs network, git and Ant: builds tlaplus at the pinned commit
scripts/verify --jar .work/jars/dc6a944/tla2tools.jar --java /path/to/java --jobs 4
scripts/verify --jar J --sets sets/tla-project            # one part of the tree
scripts/verify --jar J --only '^A-err-'                   # by id
scripts/verify --jar J --show A-err-e4200                 # one case, both sides
scripts/verify --jar J --report mismatches.tsv            # every difference
```

It prints matches and mismatches per set and exits 1 on any mismatch. A mismatch is a difference from the reference, not necessarily a bug. The report says when a case is already known to depend on platform, flags, stack size or charset.

Results so far:

- From a fresh clone on the reference machine: 9,453 match, 0 mismatch.
- On macOS arm64 (APFS): 9,452 match. The one mismatch is `D-5ac9249fe910`, which differs on any case-insensitive file system (see [The stock check](reference.md#the-stock-check)).
- A full run took about 12 minutes at 24 jobs on a loaded 40-thread Xeon, and about 14 minutes at one job without `-error-codes` on an M-series laptop.

Notes:

- A build of `dc6a944` from source gives the same class files as the recorded jar, but not the same SHA-256: only `META-INF/MANIFEST.MF` differs, by build date (checked).
- For a jar without `-error-codes`, `verify` reads the verdict from the text alone. On all 9,453 cases here the text agreed with the exit code.
- For v1.7.4, add `--isolate-tmp` if `--jobs` is above 1: that jar unpacks standard modules into one shared directory.
- SANY leaves an empty `tlc-*` directory in `java.io.tmpdir` for each JVM. `verify` removes the new ones unless you pass `--keep-tmp`.

`scripts/regenerate` needs git and Go. It clones the pinned repositories, rebuilds the upstream sets, the mutants and the grammar sentences, and compares them with the tree: 7,518 same, 0 differ.

`cmd/casemap` answers questions about coverage:

```sh
go run ./cmd/casemap case A-err-e4200                # what one case reaches, and why it is here
go run ./cmd/casemap who OpApplNode.levelCheck       # which cases reach a method
go run ./cmd/casemap who TLAplusParser.java:5551     # a line, with its arms
go run ./cmd/casemap who AssumeProve                 # a grammar production
go run ./cmd/casemap who token:BOXASSUME             # a token kind
go run ./cmd/casemap grammar -prod AssumeProve
go run ./cmd/casemap sources
go run ./cmd/casemap select                          # recompute every why
go run ./cmd/casemap prove                           # the coverage proof: coverage.md, "Coverage kept"
go run ./cmd/staticreach -jdk "$JAVA_HOME"           # the static map; first scripts/regenerate, scripts/setup reference, scripts/maintainer/unpack-classes
```
