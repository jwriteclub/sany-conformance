<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# tlaplus/Examples

433 cases, kept whole: this set's reason for being here is that it is upstream's.

**What it is.** The community collection of specifications.

**Where it came from.** `specifications/` in tlaplus/Examples.

**What it expects.** Nothing beyond SANY's recorded output. Every row's `expected` is what plain stock SANY printed, recorded from the reference build; it says what SANY does and not whether that is right. See the top-level `README.md`.

**How to regenerate.** `scripts/regenerate` clones `sources.lock` and runs `cmd/extract`, which rewrites these cases byte for byte.

**Recorded output.** 431 accept, 2 reject.

Files: `manifest.jsonl`, one row per case, the first line the environment it was recorded in; `cases/<id>/`, the root module and its support modules, with library modules as links into `libs/`; `LICENSE.md`.
