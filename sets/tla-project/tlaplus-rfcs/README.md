<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# tlaplus/rfcs language standard

475 cases, kept whole: this set's reason for being here is that it is upstream's.

**What it is.** `language_standard/tests`, the origin of upstream's syntax, semantic and error corpora, and its syntax corpus. Nearly all duplicates of the corpora above, by content; kept, because this set is upstream's as it stands.

**Where it came from.** `language_standard/` in tlaplus/rfcs.

**What it expects.** Its syntax-corpus rows are marked `other-parser`: another repository's assertion, kept for reference and not taken as an assertion about SANY. The rest carry nothing beyond SANY's recorded output. Every row's `expected` is what plain stock SANY printed, recorded from the reference build; it says what SANY does and not whether that is right. See the top-level `README.md`.

**How to regenerate.** `scripts/regenerate` clones `sources.lock` and runs `cmd/extract`, which rewrites these cases byte for byte.

**Recorded output.** 1 crash, counted as reject, 171 accept, 303 reject.

Files: `manifest.jsonl`, one row per case, the first line the environment it was recorded in; `cases/<id>/`, the root module and its support modules, with library modules as links into `libs/`; `LICENSE.md`.
