<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# TLC test models

670 cases, kept whole: this set's reason for being here is that it is upstream's.

**What it is.** Modules under `test-model`. 12 are fixtures of SANY tests; the rest exist for TLC's tests and SANY only happens to read them.

**Where it came from.** `tlatools/org.lamport.tlatools/test-model` in tlaplus/tlaplus.

**What it expects.** The 12 fixtures carry an upstream assertion: 2 reject with message text, 10 accept weakly (XML export exit 0, or no exception), one of them `@Ignore`d. The rest assert nothing about SANY. Every row's `expected` is what plain stock SANY printed, recorded from the reference build; it says what SANY does and not whether that is right. See the top-level `README.md`.

**How to regenerate.** `scripts/regenerate` clones `sources.lock` and runs `cmd/extract`, which rewrites these cases byte for byte.

**Recorded output.** 637 accept, 33 reject.

Files: `manifest.jsonl`, one row per case, the first line the environment it was recorded in; `cases/<id>/`, the root module and its support modules, with library modules as links into `libs/`; `LICENSE.md`.
