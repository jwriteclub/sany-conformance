\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: Generator.java:3358 RemoveConditional_EQUAL_ELSE: wrong argument count to a recursive function goes unreported at A
\* Aim: Generator.java:3378 NullReturnVals: null operand of B, so C gets an extra 'unexpected null node' message
---- MODULE Gap09 ----
f[x \in {1}, y \in {2}] == IF x = 1 THEN f[x, y] ELSE 0
A == f[1, 2, 3]
B == {f[1, 2, 3]}
C == B!1
====
