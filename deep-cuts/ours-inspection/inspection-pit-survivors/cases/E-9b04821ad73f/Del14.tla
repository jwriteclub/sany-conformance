\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: LetInNode.java:167 VoidMethodCall: L2 keeps the argument level constraint (P, 0) from the LET definition d; mutant drops it (visible in getArgLevelConstraints, not in the XML)
---- MODULE Del14 ----
VARIABLE x
B1(P(_)) == P(x)
L2(P(_)) == LET d == B1(P) IN 1
====
