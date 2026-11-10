<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# Written by reading the code

42 cases selected here, 146 more in [`deep-cuts/ours-inspection/inspection`](../../../deep-cuts/ours-inspection/inspection/).

**What it is.** Specs review agents wrote after reading the code, each aimed at something in it. `source.aim` is what it was aimed at, as its author or the review states it.

**Where it came from.** Code reviews of SANY. Authors are in `source.author`.

**What it expects.** Nothing beyond SANY's recorded output. Every row's `expected` is what plain stock SANY printed, recorded from the reference build; it says what SANY does and not whether that is right. See the top-level `README.md`.

**How to regenerate.** The bytes are here.

**Recorded output.** 2 crash, counted as reject, 59 accept, 127 reject.

Files: `manifest.jsonl`, one row per case, the first line the environment it was recorded in; `cases/<id>/`, the root module and its support modules, with library modules as links into `libs/`; `LICENSE.md`.
