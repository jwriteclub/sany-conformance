<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# Semantic error corpus

121 cases, kept whole: this set's reason for being here is that it is upstream's.

**What it is.** One spec per error or warning code, with the code in the file name.

**Where it came from.** `test/tla2sany/semantic/error_corpus` in tlaplus/tlaplus.

**What it expects.** An upstream assertion: the code in the file name appears in the log, and E means failure. Nothing about the exact text, location or the rest of the code list; those here are SANY's output, recorded. Every row's `expected` is what plain stock SANY printed, recorded from the reference build; it says what SANY does and not whether that is right. See the top-level `README.md`.

**How to regenerate.** `scripts/regenerate` clones `sources.lock` and runs `cmd/extract`, which rewrites these cases byte for byte.

**Recorded output.** 1 crash, counted as reject, 7 accept, 113 reject.

Files: `manifest.jsonl`, one row per case, the first line the environment it was recorded in; `cases/<id>/`, the root module and its support modules, with library modules as links into `libs/`; `LICENSE.md`.
