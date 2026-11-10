<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# One-shot model loop

163 cases selected here, 178 more in [`deep-cuts/ours-model/model-one-shot`](../../../deep-cuts/ours-model/model-one-shot/).

**What it is.** Specs from a loop giving a model the target code, taking one reply, and replaying it. `source.model` and `source.target` name them. Prompts carried nearby specs from the corpus, so a reply can echo them; `source.trace` measures that by content.

**Where it came from.** Model replies, kept when a replay in a fresh JVM showed a new branch, error-code combination, tlx disagreement or Java crash.

**What it expects.** Nothing beyond SANY's recorded output. Every row's `expected` is what plain stock SANY printed, recorded from the reference build; it says what SANY does and not whether that is right. See the top-level `README.md`.

**How to regenerate.** The bytes are here. Rerunning needs models and gives different cases.

**Recorded output.** 6 crash, counted as reject, 77 accept, 258 reject.

Files: `manifest.jsonl`, one row per case, the first line the environment it was recorded in; `cases/<id>/`, the root module and its support modules, with library modules as links into `libs/`; `LICENSE.md`.
