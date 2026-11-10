\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: Linter.java:162 RemoveConditional_EQUAL_ELSE: BOOLEAN no longer counts as built from declarations: warning for f4 appears
\* Aim: Linter.java:167 RemoveConditional_EQUAL_IF: a user Int counts as Integers' Int: warning for f1 goes
\* Aim: Linter.java:172 RemoveConditional_EQUAL_IF: a user Nat counts as Naturals' Nat: warning for f2 goes
\* Aim: Linter.java:177 RemoveConditional_EQUAL_IF: a user Reals counts as the Reals stub: warning for f3 goes
---- MODULE Other18 ----
Int == 1
Nat == 1
Reals == 1
f1 == 1
f2 == 1
f3 == 1
f4 == 1
R1 == [f1 |-> Int]
R2 == [f2 |-> Nat]
R3 == [f3 |-> Reals]
R4 == [f4 |-> BOOLEAN]
====
