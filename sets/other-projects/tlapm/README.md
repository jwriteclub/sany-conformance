<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# tlapm, the TLA+ proof manager

677 cases, kept whole: this set's reason for being here is that it is upstream's.

**What it is.** The proof manager's library, examples, tests and its copy of the syntax corpus. `test/sany/equivalence` compares tlapm's parser with SANY, tolerating SANY failure.

**Where it came from.** tlaplus/tlapm, by each row's `source.path`.

**What it expects.** Its syntax-corpus rows are `other-parser`: tlapm's own test, nothing about SANY. The rest carry nothing beyond SANY's recorded output. Every row's `expected` is what plain stock SANY printed, recorded from the reference build; it says what SANY does and not whether that is right. See the top-level `README.md`.

**How to regenerate.** `scripts/regenerate` clones `sources.lock` and runs `cmd/extract`, which rewrites these cases byte for byte.

**Recorded output.** 589 accept, 88 reject.

Files: `manifest.jsonl`, one row per case, the first line the environment it was recorded in; `cases/<id>/`, the root module and its support modules, with library modules as links into `libs/`; `LICENSE.md`.
