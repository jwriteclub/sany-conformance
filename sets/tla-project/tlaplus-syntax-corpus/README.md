<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# Unified syntax corpus

355 cases, kept whole: this set's reason for being here is that it is upstream's.

**What it is.** The syntax corpus SANY's own tests run, including the Unicode files. One snippet per case.

**Where it came from.** `test/tla2sany/corpus` in tlaplus/tlaplus at the pinned commit. Each snippet is wrapped into a module by `cmd/extract`.

**What it expects.** An upstream assertion, syntax stage only: accept or reject, and for accepts an expected S-expression tree, kept as `expected.sexp`. Many snippets use undefined names, so an accept there does not mean full SANY accepts, and the expected output here is full SANY's. 12 are in upstream's `expectFailures`: SANY deviates from them, recorded in `upstream_assertion`. Every row's `expected` is what plain stock SANY printed, recorded from the reference build; it says what SANY does and not whether that is right. See the top-level `README.md`.

**How to regenerate.** `scripts/regenerate` clones `sources.lock` and runs `cmd/extract`, which rewrites these cases byte for byte.

**Recorded output.** 139 accept, 216 reject.

Files: `manifest.jsonl`, one row per case, the first line the environment it was recorded in; `cases/<id>/`, the root module and its support modules, with library modules as links into `libs/`; `LICENSE.md`.
