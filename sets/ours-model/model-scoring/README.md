<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# Scoring specs

5 cases selected here, 3 more in [`deep-cuts/ours-model/model-scoring`](../../../deep-cuts/ours-model/model-scoring/).

**What it is.** Specs a model wrote to back a `yes` verdict while scoring the remaining lines. None took its line.

**Where it came from.** Scoring sessions; each row names the model and target.

**What it expects.** Nothing beyond SANY's recorded output. Every row's `expected` is what plain stock SANY printed, recorded from the reference build; it says what SANY does and not whether that is right. See the top-level `README.md`.

**How to regenerate.** The bytes are here. Rerunning needs models and gives different cases.

**Recorded output.** 8 reject.

Files: `manifest.jsonl`, one row per case, the first line the environment it was recorded in; `cases/<id>/`, the root module and its support modules, with library modules as links into `libs/`; `LICENSE.md`.
