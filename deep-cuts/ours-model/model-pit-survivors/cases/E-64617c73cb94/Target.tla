\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: tla2sany/semantic/APSubstInNode.java:311:BooleanTrueReturnValsMutator (model @cf/zai-org/glm-5.3-flash)
---- MODULE Target ----
---- MODULE Inner ----
VARIABLE y
THEOREM B == ENABLED []y
======================
VARIABLE z
M == INSTANCE Inner WITH y <- ENABLED []ENABLED []z
============================================================
