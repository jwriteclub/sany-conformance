\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: Generator.java:1495 NullReturnVals: null from a GeneralId selector, NPE at retVal.getKind() for B
\* Aim: Generator.java:1613 NullReturnVals: same for C, selecting an ASSUME/PROVE
---- MODULE Gap16 ----
G(F(_)) == F(1)
A == G(LAMBDA x : x)
B == A!1!(1, 2)
THEOREM T == ASSUME TRUE PROVE TRUE
C == T!:
====
