<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# Jazzer fuzz outputs

1132 cases selected here, 103 more in [`deep-cuts/ours-mechanical/fuzz-jazzer`](../../../deep-cuts/ours-mechanical/fuzz-jazzer/).

**What it is.** What the Jazzer 0.30.0 fuzzer found on SANY, minimised by lines and tokens. Some contain bytes that are not valid UTF-8. 8 held a local file path, which was deleted, with verdict, exit status, codes and coverage measured unchanged; `source.derived_from` names the original.

**Where it came from.** A harness over SANY seeded as `LICENSE.md` lists. libFuzzer records no parent, so `source.strategy` names only the strategy and unit. `source.trace`, where present, is the file closest by content (`cmd/tracecases`): provenance, not a licence claim.

**What it expects.** Nothing beyond SANY's recorded output. Every row's `expected` is what plain stock SANY printed, recorded from the reference build; it says what SANY does and not whether that is right. See the top-level `README.md`.

**How to regenerate.** The bytes are here. Rerunning a fuzzer gives different cases.

**Recorded output.** 16 crash, counted as reject, 3 accept, 1216 reject.

Files: `manifest.jsonl`, one row per case, the first line the environment it was recorded in; `cases/<id>/`, the root module and its support modules, with library modules as links into `libs/`; `LICENSE.md`.
