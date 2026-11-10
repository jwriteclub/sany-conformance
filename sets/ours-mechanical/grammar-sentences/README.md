<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# Sentences derived from the grammar

15 cases selected here, 373 more in [`deep-cuts/ours-mechanical/grammar-sentences`](../../../deep-cuts/ours-mechanical/grammar-sentences/).

**What it is.** One short sentence per choice point of `tla+.jj`, in the manner of Purdom, ignoring semantic lookahead and actions, so many are rejected.

**Where it came from.** `cmd/gensentences` over the pinned `tla+.jj`.

**What it expects.** Nothing beyond SANY's recorded output. Every row's `expected` is what plain stock SANY printed, recorded from the reference build; it says what SANY does and not whether that is right. See the top-level `README.md`.

**How to regenerate.** `scripts/regenerate`.

**Recorded output.** 1 crash, counted as reject, 56 accept, 331 reject.

Files: `manifest.jsonl`, one row per case, the first line the environment it was recorded in; `cases/<id>/`, the root module and its support modules, with library modules as links into `libs/`; `LICENSE.md`.
