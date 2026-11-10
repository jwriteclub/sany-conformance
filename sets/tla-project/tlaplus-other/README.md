<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# Other modules in the TLA+ tools repository

296 cases, kept whole: this set's reason for being here is that it is upstream's.

**What it is.** Every other `.tla` file in tlaplus/tlaplus: `general`, the toolbox, tlatools' spec, source and test resources, and helper modules of the error corpus.

**Where it came from.** tlaplus/tlaplus, by each row's `source.path`.

**What it expects.** Nothing beyond SANY's recorded output. Every row's `expected` is what plain stock SANY printed, recorded from the reference build; it says what SANY does and not whether that is right. See the top-level `README.md`.

**How to regenerate.** `scripts/regenerate` clones `sources.lock` and runs `cmd/extract`, which rewrites these cases byte for byte.

**Recorded output.** 259 accept, 37 reject.

Files: `manifest.jsonl`, one row per case, the first line the environment it was recorded in; `cases/<id>/`, the root module and its support modules, with library modules as links into `libs/`; `LICENSE.md`.
