<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# Library modules

Modules that some cases import and that SANY does not carry: `CommunityModules` from tlaplus/CommunityModules (MIT) and `tlapm-library` from tlaplus/tlapm's `library` (BSD-2-Clause), at the commits in `sources.lock`, each with its licence. A case that parses one has a link to it in its own directory, so the plain command finds it there. `files` in each manifest row lists them with `role: library`. Note that `tlapm-library/FiniteSets.tla` is not SANY's own: a case linking it reads that copy.
