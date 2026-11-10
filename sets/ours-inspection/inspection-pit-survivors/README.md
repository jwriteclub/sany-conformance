<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# Written by reading a surviving mutant's code

68 cases selected here, 4 more in [`deep-cuts/ours-inspection/inspection-pit-survivors`](../../../deep-cuts/ours-inspection/inspection-pit-survivors/).

**What it is.** Specs agents wrote after reading the code of a mutant no case noticed. Each states its aim in `\* Aim:` lines. Their file headers say all rights reserved; `LICENSE.md` says how that is read here.

**Where it came from.** A PIT 1.30.1 mutation run over SANY.

**What it expects.** Nothing beyond SANY's recorded output. Every row's `expected` is what plain stock SANY printed, recorded from the reference build; it says what SANY does and not whether that is right. See the top-level `README.md`.

**How to regenerate.** The bytes are here.

**Recorded output.** 37 accept, 35 reject.

Files: `manifest.jsonl`, one row per case, the first line the environment it was recorded in; `cases/<id>/`, the root module and its support modules, with library modules as links into `libs/`; `LICENSE.md`.
