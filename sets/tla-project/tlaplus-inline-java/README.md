<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# Specs written inline in Java tests

52 cases, kept whole: this set's reason for being here is that it is upstream's.

**What it is.** The `TestLevelChecking` table and one `ParseErrorTests` entry, written as Java strings upstream and extracted by regular expression.

**Where it came from.** `test/tla2sany/semantic/TestLevelChecking.java` and `test/tla2sany/parser/ParseErrorTests.java` in tlaplus/tlaplus.

**What it expects.** An upstream assertion: parse and semantics succeed with the level check result listed, or a syntax rejection with message text. Every row's `expected` is what plain stock SANY printed, recorded from the reference build; it says what SANY does and not whether that is right. See the top-level `README.md`.

**How to regenerate.** `scripts/regenerate` clones `sources.lock` and runs `cmd/extract`, which rewrites these cases byte for byte.

**Recorded output.** 30 accept, 22 reject.

Files: `manifest.jsonl`, one row per case, the first line the environment it was recorded in; `cases/<id>/`, the root module and its support modules, with library modules as links into `libs/`; `LICENSE.md`.
