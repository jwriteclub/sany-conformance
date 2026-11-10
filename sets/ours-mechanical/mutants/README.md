<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# Seeded mutants of the corpus

63 cases selected here, 2937 more in [`deep-cuts/ours-mechanical/mutants`](../../../deep-cuts/ours-mechanical/mutants/).

**What it is.** 3,000 mutants, seed 20261103, of upstream parents: delete a token, swap a delimiter, misalign a junction, use an unknown name, prime in an invariant, bad arity. At most 2 per parent, parents at most 30 KiB. `source` names each parent and rule.

**Where it came from.** `cmd/extract/mutate.go`. Each is a small edit of its parent and follows its parent's licence.

**What it expects.** Nothing beyond SANY's recorded output. `source.intent` (`likely-reject` or `unknown`) is a hint from the rule, not an assertion. Every row's `expected` is what plain stock SANY printed, recorded from the reference build; it says what SANY does and not whether that is right. See the top-level `README.md`.

**How to regenerate.** `scripts/regenerate`, which is seeded and deterministic.

**Recorded output.** 1 crash, counted as reject, 776 accept, 2223 reject.

Files: `manifest.jsonl`, one row per case, the first line the environment it was recorded in; `cases/<id>/`, the root module and its support modules, with library modules as links into `libs/`; `LICENSE.md`.
