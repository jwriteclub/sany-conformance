<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# Licences

- Every source repository's top-level licence allows redistribution as long as its notice is kept. All are MIT except tlapm, which is BSD-2-Clause.
- Our own cases and code are Apache-2.0.
- A mutant takes its parent's licence.
- The fuzz set is Apache-2.0 as a whole.
- All 9,453 cases are in this tree.

This is our reading of the files at the pinned commits. It is not legal advice. "Needs a human look" below lists what it cannot settle.

Every row of every `manifest.jsonl` has `licence.spdx` and `licence.file`. `licences/` holds a copy of each licence file a case points to, under `<repository>/<path>`. Each set's `LICENSE.md` counts the licences of its own cases.

## Repositories

| repository | commit | licence | holders |
|---|---|---|---|
| tlaplus/tlaplus | `dc6a944429fe` | MIT | 199? HP Corporation; 2003 Microsoft Corporation; 2023 Linux Foundation |
| tlaplus/Examples | `0e32a039a428` | MIT | 2016: The TLA+ project and other contributors |
| tlaplus/CommunityModules | `1805e9d5e1ce` | MIT | 2019 TLA+ |
| tlaplus/rfcs | `f917f9c83769` | MIT | 2025 Linux Foundation |
| tlaplus/tlapm | `568722253719` | BSD-2-Clause | 2008 INRIA and Microsoft Corporation; 2023 Linux Foundation |
| tlaplus-x/tlx | `bec5985379ad` | MIT | 2026 The tlx authors |
| tlaplus-community/tree-sitter-tlaplus | `add40814fda3` | MIT | 2020 Andrew Helwer |
| Disalg-ICS-NJU/PaxosStore-TLA | `0188955e7023` | MIT | 2019 hengxin (Hengfeng Wei) |

Some repositories have licence files further down, which govern the files beneath them. Each row these govern names the file:

| file | licence | holders |
|---|---|---|
| `Examples/specifications/KeyValueStore/LICENSE` | BSD-2-Clause | 2020, ING Bank / CWI |
| `Examples/specifications/SDP_Verification/LICENSE` | MIT | 2022 Dong Luming |
| `Examples/specifications/byihive/LICENSE` | Apache-2.0 | named in each file's header |
| `tlaplus/general/performance/Bookkeeper/LICENSE` | MIT | 2021 Jack Vanlightly |
| `tlaplus/general/performance/PaxosMadeSimple/LICENSE.md` | MIT | 2015 nano-o |
| `tlapm/zenon/LICENSE` | BSD-3-Clause | 1997-2006, Institut National de Recherche en Informatique et en Automatique |
| `tlx/crates/tlx-sema/modules/community/LICENSE` | MIT | 2019 TLA+ |

## How each kind of case is read

- **Upstream cases** carry the nearest licence file above their path in the source repository.
- **Mutants** are one small edit of a parent, so each takes its parent's licence. MIT stays MIT, and `tlapm` stays BSD-2-Clause.
- **Apache-2.0 parents.** 12 mutants derive from `Examples/specifications/byihive`. Section 4(b) of Apache-2.0 asks that a modified file carry a notice that it was changed. We cannot edit the bytes, because they are under test, so each mutant has a `CHANGES.txt` beside it. The set's 6 unmodified files keep their notices as they are.
- **Fuzz outputs** are Apache-2.0 for the set as a whole, with no per-case parent licence. libFuzzer records no parent, and everything the runs were fed is open source or ours, so nothing is compiled against them. `sets/ours-mechanical/fuzz-jazzer/LICENSE.md` lists what the runs were fed, measured from the directories they used, and points to each source's licence in `licences/`, so that the MIT and BSD notices stay with this repository. `source.trace` and `data/fuzz-trace.tsv` name the closest file by content. That is provenance, not a licence claim.
- **Grammar sentences** are generated from `tla+.jj`, and we read them as taking its licence: MIT.
- **Model and agent specs** are ours: Apache-2.0. The prompts included nearby corpus specs, so a reply could copy from them. We measured all 700: the longest run of bytes any of them shares with any corpus file has a median of 13 bytes and a maximum of 127. Each row's `source.trace` names the closest corpus file.
- **Library modules** in `libs/` keep their own licences beside them.

## PaxosStore-TLA

tlx's `corpus/external.json` records `Disalg-ICS-NJU/PaxosStore-TLA` as having no licence. That record is out of date: at the commit tlx pins, `0188955e70`, the repository's root `LICENSE` is MIT, and GitHub reports the same. The licence is in `licences/PaxosStore-TLA/LICENSE`, and the commit is in `sources.lock`.

The tree holds tlx's 4 wrapper modules `corpus/external/paxosstore/MC_tlx_*.tla`, which are tlx's own work under its MIT licence, and 5 mutants of them. tlx does not vendor the modules they extend, so the recorded SANY output reports those modules missing. None of upstream's modules is in this tree, because adding them would change the recorded output for 6 of the 9.

## Needs a human look

1. **`All rights reserved` in 99 files.** The specs from the mutation run (`inspection-pit-survivors`, `model-pit-survivors`) start with `Copyright 2026 Christopher O'Connell` and `All rights reserved`. Those lines are bytes under test, so they stay: rewriting them would change every position in those cases. The holder has licensed these sets under Apache-2.0 for this repository, but a reader will see both statements.
2. **Whether `CHANGES.txt` satisfies Apache-2.0 section 4(b).** The section says modified files must "carry" a notice. Ours sits beside the file.
3. **Files with their own headers: 77 cases**, in `tlaplus-examples`, `tlaplus-test-models`, `tlaplus-other` and copies of them. Some name other holders (Oracle, Backyard Innovations). A group of `aba_asyn_byz`-style files says "subject to the license that is bundled together with this package ... in the file LICENSE", and no such file is beside them at the pinned commit. tlaplus/Examples's `LICENSE.md` says all its examples are MIT, so we read them as MIT, but the header points elsewhere and we did not ask its holder. 6 fuzz outputs keep a third party's copyright line in their text.
4. **tlx's other specs were not traced.** tlx is MIT, by "The tlx authors", and its `tests/specs` state no other origin. By content, 24 of its rows duplicate CommunityModules files. Its other `corpus/external` wrappers name upstream projects that tlx records as CC-BY-3.0-US, Apache-2.0 or MIT. We did not read those upstream repositories, and tlx's record for PaxosStore-TLA turned out to be out of date.
5. **Model providers' terms** were not audited. `source.model` names the model behind each spec.

## Not checked

- Whether each file's stated licence is valid for that file. A repository saying MIT does not show that its author could grant it.
- Upstream repositories named only in `tlx/corpus/external.json`.
- Files with no header in repositories that have a licence file. We read them as covered by that file.
