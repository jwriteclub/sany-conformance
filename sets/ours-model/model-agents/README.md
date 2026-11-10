<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# Exploring agents

19 cases selected here, 22 more in [`deep-cuts/ours-model/model-agents`](../../../deep-cuts/ours-model/model-agents/).

**What it is.** Specs from sessions of a model with tools on one block of missed lines.

**Where it came from.** Agent sessions; each row names the model and target.

**What it expects.** Nothing beyond SANY's recorded output. Every row's `expected` is what plain stock SANY printed, recorded from the reference build; it says what SANY does and not whether that is right. See the top-level `README.md`.

**How to regenerate.** The bytes are here. Rerunning needs models and gives different cases.

**Recorded output.** 5 accept, 36 reject.

Files: `manifest.jsonl`, one row per case, the first line the environment it was recorded in; `cases/<id>/`, the root module and its support modules, with library modules as links into `libs/`; `LICENSE.md`.
