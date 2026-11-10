<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# Specs hunting a stale jj_ntk

0 cases selected here, 0 more in [`deep-cuts/ours-model/model-ternary-hunt`](../../../deep-cuts/ours-model/model-ternary-hunt/).

**What it is.** Empty: models tried to make the memo in `(jj_ntk==-1)?jj_ntk():jj_ntk` go stale in 509 runs and none did, and nothing met the bar to be kept. The directory stays so that the set is accounted for. See FINDINGS.md.

**Where it came from.** Agent sessions.

**What it expects.** Nothing. Every row's `expected` is what plain stock SANY printed, recorded from the reference build; it says what SANY does and not whether that is right. See the top-level `README.md`.

**How to regenerate.** Nothing to regenerate.

Files: `manifest.jsonl`, one row per case, the first line the environment it was recorded in; `cases/<id>/`, the root module and its support modules, with library modules as links into `libs/`; `LICENSE.md`.
