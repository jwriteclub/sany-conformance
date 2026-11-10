<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# Licence of this set

Apache-2.0, at set level: `../../../LICENSE`. No case carries a parent's licence. A reading, not legal advice. `LICENCES.md` at the top says why.

## What the fuzzer was fed

Measured from the directories the runs used, not assumed. `data/fuzz-seeds.tsv` lists every seed; `data/fuzz-seed-sources.json` holds these totals.

**Seeds: 907 distinct cases, chosen per strategy.** The licences below are the seeds' own, kept in `licences/` so that their notices stay with this repository.

| from | licence | files | mutants of files there | notice |
|---|---|---:|---:|---|
| `CommunityModules` | MIT | 2 | 6 | `../../../licences/CommunityModules/LICENSE` |
| `Examples` | MIT | 24 | 28 | `../../../licences/Examples/LICENSE.md` |
| `Examples/specifications/KeyValueStore` | BSD-2-Clause | 0 | 1 | `../../../licences/Examples/specifications/KeyValueStore/LICENSE` |
| our own model specs | Apache-2.0 | 146 | | `../../../LICENSE` |
| `rfcs` | MIT | 97 | 1 | `../../../licences/rfcs/LICENSE` |
| `tlaplus` | MIT | 174 | 155 | `../../../licences/tlaplus/LICENSE` |
| `tlaplus/general/performance/PaxosMadeSimple` | MIT | 1 | 0 | `../../../licences/tlaplus/general/performance/PaxosMadeSimple/LICENSE.md` |
| `tlapm` | BSD-2-Clause | 98 | 50 | `../../../licences/tlapm/LICENSE` |
| `tlapm/zenon` | BSD-3-Clause | 5 | 7 | `../../../licences/tlapm/zenon/LICENSE` |
| `tlx` | MIT | 24 | 44 | `../../../licences/tlx/LICENSE` |
| `tlx/crates/tlx-sema/modules/community` | MIT | 1 | 0 | `../../../licences/tlx/crates/tlx-sema/modules/community/LICENSE` |
| `tree-sitter-tlaplus` | MIT | 43 | 0 | `../../../licences/tree-sitter-tlaplus/LICENSE` |

**Also fed to it:** the dictionary `../../../data/fuzz-dictionary.txt`, generated from `tla+.jj` in tlaplus/tlaplus (MIT); 1576 earlier kept outputs (1235 of this set, 340 model specs, 1 inspection); 1322 model replies that got near a target; and up to 3000 of its own earlier units. `CommunityModules` (MIT) and tlapm's `library` (BSD-2-Clause) were on the library path, and `../../../libs/` holds them with their licences.

**Nothing else.** No seed comes from tlx's `corpus/external`, and none of those 7822 files mentions PaxosStore. Model specs are ours, written from prompts that carried corpus text.
