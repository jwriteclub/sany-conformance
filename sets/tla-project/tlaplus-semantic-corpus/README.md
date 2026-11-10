<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# Semantic corpus

27 cases, kept whole: this set's reason for being here is that it is upstream's.

**What it is.** Specs upstream's semantic tests expect to pass with no errors or warnings.

**Where it came from.** `test/tla2sany/semantic/corpus` in tlaplus/tlaplus.

**What it expects.** An upstream assertion: no errors, no warnings, embedded RefersTo and IsLevel assertions hold. `NegativeOpTest` is skipped upstream (tlaplus/tlaplus#1130). Every row's `expected` is what plain stock SANY printed, recorded from the reference build; it says what SANY does and not whether that is right. See the top-level `README.md`.

**How to regenerate.** `scripts/regenerate` clones `sources.lock` and runs `cmd/extract`, which rewrites these cases byte for byte.

**Recorded output.** 26 accept, 1 reject.

Files: `manifest.jsonl`, one row per case, the first line the environment it was recorded in; `cases/<id>/`, the root module and its support modules, with library modules as links into `libs/`; `LICENSE.md`.
