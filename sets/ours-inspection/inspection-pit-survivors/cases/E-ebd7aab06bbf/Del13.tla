\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: ThmOrAssumpDefNode.java:524 RemoveConditional_EQUAL_ELSE: It!T2 keeps the arg level param (Gp, 0, cc), and K its constraints; mutant drops them (visible in getArgLevelParams, not in the XML)
---- MODULE Del13 ----
VARIABLE x
CONSTANT cc
---- MODULE Mt ----
VARIABLE v
CONSTANT G2(_), a2
THEOREM T2 == G2(a2)
====
It(Gp(_)) == INSTANCE Mt WITH v <- x, G2 <- Gp, a2 <- cc
K(q) == It(LAMBDA y : y')!T2
====
