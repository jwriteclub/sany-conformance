\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: LetInNode.java:127 RemoveConditional_EQUAL_ELSE: LET stays level-correct after a bad definition: E1 gets a second error from the outer prime
\* Aim: LetInNode.java:131 RemoveConditional_EQUAL_ELSE: same for a bad body: E2
\* Aim: LetInNode.java:135 RemoveConditional_EQUAL_ELSE: same for a bad INSTANCE: E3
\* Aim: InstanceNode.java:148 RemoveConditional_ORDER_ELSE: bad substitution not noted: E3
\* Aim: InstanceNode.java:149 RemoveConditional_EQUAL_ELSE: same: E3
\* Aim: InstanceNode.java:145 RemoveConditional_EQUAL_ELSE: bad instantiated module not noted: E4
\* Aim: OpApplNode.java:472 RemoveConditional_EQUAL_ELSE: bad range not noted: E5
\* Aim: OpDefNode.java:1177 BooleanTrueReturn: bad definition reported as fine: E6 gets an argument level error
\* Aim: SubstInNode.java:328 RemoveConditional_EQUAL_ELSE: same through an instantiated definition: E7
\* Aim: SubstInNode.java:457 BooleanTrueReturn: same: E7
\* Aim: SubstInNode.java:332 RemoveConditional_EQUAL_ELSE: bad substitution: E8
\* Aim: APSubstInNode.java:319 RemoveConditional_EQUAL_ELSE: bad ASSUME under an instantiated theorem not noted: E9 gets an argument level error
\* Aim: APSubstInNode.java:323 RemoveConditional_EQUAL_ELSE: bad substitution: E10 gets an argument level error
---- MODULE Other22 ----
VARIABLE x
---- MODULE K ----
CONSTANT c
====
---- MODULE B ----
VARIABLE v
bad == (v')'
====
---- MODULE Q ----
VARIABLE v
b(q) == q' /\ (v')'
====
---- MODULE Q2 ----
VARIABLE v
g(q) == q' /\ v = v
====
IK == INSTANCE K WITH c <- 1
IB == INSTANCE B WITH v <- x
---- MODULE Q3 ----
VARIABLE v
CONSTANT a
THEOREM T == ASSUME (v')' PROVE a' = 1
====
---- MODULE Q3b ----
VARIABLE v
CONSTANT a
THEOREM T == ASSUME v = v PROVE a' = 1
====
I3(ap) == INSTANCE Q3 WITH v <- x, a <- ap
E9 == I3(x')!T
I4(ap) == INSTANCE Q3b WITH v <- (x')', a <- ap
E10 == I4(x')!T
E1 == (LET d == (x')' IN x')'
E2 == (LET d == 1 IN (x')')'
E3 == (LET J == INSTANCE K WITH c <- (x')' IN x')'
E4 == (LET J == INSTANCE B WITH v <- x IN x')'
E5 == (\E i \in (x')' : x')'
F(q) == q' /\ (x')'
E6 == F(x')
I1 == INSTANCE Q WITH v <- x
E7 == I1!b(x')
I2 == INSTANCE Q2 WITH v <- (x')'
E8 == I2!g(x')
====
