\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: Linter.java:173 RemoveConditional_EQUAL_IF: any Naturals op counts as Nat; a*b's stub uses no parameter, so warning for g1 goes
\* Aim: Linter.java:178 RemoveConditional_EQUAL_IF: any Reals op counts: warning for g2 (Real, upstream tests the name "Reals") goes
\* Aim: OpApplNode.java:1179 RemoveConditional_ORDER_ELSE: ranges not among children: constant c unseen: warning for g3 appears
\* Aim: LetInNode.java:237 RemoveConditional_ORDER_ELSE: INSTANCEs not among children: warning for g4 appears
\* Aim: LetInNode.java:238 MathMutator: second INSTANCE overwrites the last LET def z: warning for g6 appears (unsure of the array order)
\* Aim: Linter.java:102 EmptyObjectReturn: records of extended modules ignored: warning for g5 appears. Helpers: Other19a, Other19b, Other19c
---- MODULE Other19 ----
EXTENDS Naturals, Reals, Other19a
CONSTANT c
g1 == 1
g2 == 1
g3 == 1
g4 == 1
g5 == 1
R1 == [g1 |-> 2 * 3]
R2 == [g2 |-> Real]
R3 == [g3 |-> {i \in c : TRUE}]
R4 == [g4 |-> LET J1 == INSTANCE Other19b WITH p <- 1
                  J2 == INSTANCE Other19c WITH q <- c
              IN 1]
R5 == [g5 |-> 1]
g6 == 1
R6 == [g6 |-> LET J1 == INSTANCE Other19b WITH p <- 1
                  J2 == INSTANCE Other19c WITH q <- 1
                  z == c
              IN 1]
====
