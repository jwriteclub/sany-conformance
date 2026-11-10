\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: LetInNode.java:115 BooleanTrueReturn: the second check of a bad LET must stay false; outer prime error of E1 must stay hidden
\* Aim: LetInNode.java:197 BooleanTrueReturn: a bad LET under an operator application (E12) must keep the outer prime error hidden
\* Aim: InstanceNode.java:370 BooleanTrueReturn: bad LET INSTANCE substitution (E3)
\* Aim: LabelNode.java:222 ConditionalsBoundary and ORDER_ELSE: second check of a label returns true in the baseline, so E13's outer prime error appears; mutants hide it
\* Aim: OpArgNode.java:83 BooleanTrueReturn: bad operator passed as argument (E14)
\* Aim: APSubstInNode.java:381, AssumeProveNode.java:204, ThmOrAssumpDefNode.java:425/529, NewSymbNode.java:125 BooleanTrueReturn: bad ASSUME or NEW set under an instantiated theorem (E9, E9b, E10)
\* Aim: AssumeProveNode.java:212 NegateConditionals and ORDER_ELSE: message order of the ASSUME clause errors (here the first error comes from Q4's NEW set)
---- MODULE Del04 ----
VARIABLE x
---- MODULE K ----
CONSTANT c
====
---- MODULE Q3 ----
VARIABLE v
CONSTANT a
THEOREM T == ASSUME (v')' PROVE a' = 1
====
---- MODULE Q4 ----
VARIABLE v
CONSTANT a
THEOREM T == ASSUME NEW n \in (v')' PROVE a' = 1
====
IK == INSTANCE K WITH c <- 1
I3(ap) == INSTANCE Q3 WITH v <- x, a <- ap
E9 == I3(x')!T
E9b == I3(x')!T /\ I3(x')!T
I4(ap) == INSTANCE Q4 WITH v <- x, a <- ap
E10 == I4(x')!T
E1 == (LET d == (x')' IN x')'
E3 == (LET J == INSTANCE K WITH c <- (x')' IN x')'
E12 == ((LET d == (x')' IN x') = 1)'
E13 == (l2 :: (x')')'
Bad(q) == (x')' /\ q
F1(G(_), q) == G(q) /\ x'
E14 == (F1(Bad, 1))'
====
