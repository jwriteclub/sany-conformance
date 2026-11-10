<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# Results for prudence

At prudence `b801ed2`, over all 9,453 cases (prudence v0.1.0 has the same code; later changes are to comments and documentation), against the recorded SANY output:

| result | cases |
|---|---:|
| match | 9,424 |
| match (order-normalised) | 1 |
| mismatch | 0 |
| prudence incomplete | 0 |
| oracle unavailable (SANY crashed; not scored) | 28 |
| errors, timeouts, false complete | 0 |

Prudence agrees with SANY on all 9,425 cases where SANY reached a verdict. That includes all 4,106 modules SANY accepts, and all 4,128 scored modules from upstream projects; 2 more upstream modules crash SANY and are not scored. The order-normalised case is `fuzz-jazzer/D-54cb9f0ad55d`. TE.tla needs 2,038,985 module-acquisition visits and is scored with the 3,000,000 override; at prudence's default budget of 1,000,000 it returns incomplete.

These numbers have three limits:

- **This corpus was prudence's development set.** Prudence was fixed against these cases until they matched, so this is not an independent test.
- **The cases are not a random sample of TLA+ files**, and many are related: mutants are small edits of upstream specs.
- **Only the plain output is compared.** Warnings, error codes and internal structures are not.

This run's report and per-case results are in [`results/prudence-b801ed2/`](../results/prudence-b801ed2/): `report.md`, and `results.jsonl.gz` (read it with `gzip -dc`).
