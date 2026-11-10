<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# tlx

587 cases, kept whole: this set's reason for being here is that it is upstream's.

**What it is.** Specs from the tlx project: its tests, its fuzz directory, and `corpus/external`, tlx's wrapper modules around third-party specs that it does not vendor.

**Where it came from.** tlaplus-x/tlx, by each row's `source.path`.

**What it expects.** Nothing beyond SANY's recorded output. tlx's own verdict is not an assertion about SANY. Every row's `expected` is what plain stock SANY printed, recorded from the reference build; it says what SANY does and not whether that is right. See the top-level `README.md`.

**How to regenerate.** `scripts/regenerate` clones `sources.lock` and runs `cmd/extract`, which rewrites these cases byte for byte.

**Recorded output.** 548 accept, 39 reject.

Files: `manifest.jsonl`, one row per case, the first line the environment it was recorded in; `cases/<id>/`, the root module and its support modules, with library modules as links into `libs/`; `LICENSE.md`.
