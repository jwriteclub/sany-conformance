<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# Model specs aimed at surviving mutants

23 cases selected here, 4 more in [`deep-cuts/ours-model/model-pit-survivors`](../../../deep-cuts/ours-model/model-pit-survivors/).

**What it is.** Specs models wrote, aimed at mutants of SANY that no earlier case noticed, that then killed mutants nothing earlier in the order did. The model is in `source.model`, the target in `source.aim`.

**Where it came from.** A PIT 1.30.1 mutation run over SANY: the `\* Aim:` line names the mutant.

**What it expects.** Nothing beyond SANY's recorded output. Every row's `expected` is what plain stock SANY printed, recorded from the reference build; it says what SANY does and not whether that is right. See the top-level `README.md`.

**How to regenerate.** The bytes are here. Rerunning needs models and gives different cases.

**Recorded output.** 15 accept, 12 reject.

Files: `manifest.jsonl`, one row per case, the first line the environment it was recorded in; `cases/<id>/`, the root module and its support modules, with library modules as links into `libs/`; `LICENSE.md`.
