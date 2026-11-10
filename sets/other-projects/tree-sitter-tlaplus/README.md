<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# tree-sitter-tlaplus

357 cases, kept whole: this set's reason for being here is that it is upstream's.

**What it is.** The tree-sitter grammar's test corpus, unwrapped into modules. PlusCal-only material is left out, as SANY does not translate PlusCal.

**Where it came from.** `test/` in tlaplus-community/tree-sitter-tlaplus.

**What it expects.** Its syntax-corpus rows are `other-parser`: a tree-sitter test, nothing about SANY. The rest carry nothing beyond SANY's recorded output. Every row's `expected` is what plain stock SANY printed, recorded from the reference build; it says what SANY does and not whether that is right. See the top-level `README.md`.

**How to regenerate.** `scripts/regenerate` clones `sources.lock` and runs `cmd/extract`, which rewrites these cases byte for byte.

**Recorded output.** 160 accept, 197 reject.

Files: `manifest.jsonl`, one row per case, the first line the environment it was recorded in; `cases/<id>/`, the root module and its support modules, with library modules as links into `libs/`; `LICENSE.md`.
