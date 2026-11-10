<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# Results of the stock check

Every difference from the reference, one row per case and field. Reference: unmodified build of tlaplus `dc6a944`, jar SHA-256 `171c46c1ec562ff1893a709916ab2616865bb3a635529ec7da22c8699ae9c35a`, Temurin 17.0.21+9, Linux x86_64, ext4, `LC_ALL=en_US.UTF-8`. Every run covers all 9,453 cases.

| file | what was changed |
|---|---|
| `diff-r2.tsv.gz` | nothing: a second run, in reversed order |
| `diff-mac1.tsv.gz` | macOS 15.5, arm64, APFS, same jar, same JVM build |
| `diff-rel174.tsv.gz` | the jar: v1.7.4, GitHub's latest release, SHA-256 `936a262061c914694dfd669a543be24573c45d5aa0ff20a8b96b23d01e050e88`. Run with `-Djava.io.tmpdir` per JVM, because it unpacks standard modules into one shared directory and the run was parallel |
| `diff-rel180.tsv.gz` | the jar: v1.8.0, pre-release of 2026-11-03, SHA-256 `e20447644bfabd29d6bd0fc2c0eb1473fb301c1a0e4dfe1961938ce5a191edbf`, 7 commits past `dc6a944`, none touching `tla2sany` |
| `diff-c1.tsv.gz` | `-XX:TieredStopAtLevel=1`, which earlier measurements used |
| `diff-xss512k.tsv.gz`, `diff-xss16m.tsv.gz` | thread stack: `-Xss512k`, `-Xss16m`. The default is 1 MiB on this Linux and 2 MiB on the Mac |
| `diff-posix.tsv.gz` | `LC_ALL=C`: default charset `ANSI_X3.4-1968` |
| `diff-latin1.tsv.gz` | `-Dfile.encoding=ISO-8859-1` |
| `diff-ntk.tsv.gz` | a jar with every `(jj_ntk==-1)?jj_ntk():jj_ntk` in the parser replaced by `jj_ntk()`: see FINDINGS.md |
| `diff-harness.tsv` | the harness column against the stock verdict, and against earlier results that came through the wrapper with `-XX:TieredStopAtLevel=1` |
| `depends-on.tsv` | every case whose result moved with platform, flags, stack or charset, with what moved. Also in each manifest row as `depends_on` |
| `stock-summary.json` | counts of all of the above, and each run's environment |

Fields: `status`, `verdict`, `first_position`, `crash`; `stdout-progress-lines-only` when only SANY's progress lines changed, `stdout-messages` when anything else did; `stderr`.

`prudence-b801ed2/` is a different kind of record: the published comparison of prudence with the recorded SANY output (`report.md`, and `results.jsonl.gz`, one row per case), made with `scripts/compare.sh --deep-cuts --jobs 32 --timeout 120s --info --no-3d`.
