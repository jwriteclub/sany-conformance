<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# OpenStart scan modules

23 cases selected here, 0 more in [`deep-cuts/ours-model/model-openstart`](../../../deep-cuts/ours-model/model-openstart/).

**What it is.** Modules a model (GPT-6 Astra) wrote with no tools and without running them, aimed at the eight alternatives of the `OpenStart` lookahead scan. Evidence: they reach nothing new by coverage, and each has a stated purpose.

**Where it came from.** One prompt, recorded in each row's provenance.

**What it expects.** Nothing beyond SANY's recorded output. The author's expectation is not an assertion. Every row's `expected` is what plain stock SANY printed, recorded from the reference build; it says what SANY does and not whether that is right. See the top-level `README.md`.

**How to regenerate.** The bytes are here.

**Recorded output.** 19 accept, 4 reject.

Files: `manifest.jsonl`, one row per case, the first line the environment it was recorded in; `cases/<id>/`, the root module and its support modules, with library modules as links into `libs/`; `LICENSE.md`.
