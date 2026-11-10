\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: ModuleNode.java:770 BooleanTrueReturn: B1's cached level result (false) gates the outer prime of G1; mutant adds an error at G1
\* Aim: ModuleNode.java:935 RemoveConditional_EQUAL_ELSE: B1 is incorrect only through a definition; same gate G1
\* Aim: ModuleNode.java:954 RemoveConditional_EQUAL_ELSE: B3..B6 are incorrect only through top-level theorem nodes; same gates G3..G6
\* Aim: ModuleNode.java:921 RemoveConditional_EQUAL_ELSE and :1033 BooleanTrueReturn: O7 is incorrect only through its inner module B7; gate G7
\* Aim: AssumeNode.java:107 BooleanTrueReturn: unnamed ASSUME (v')' in B2; gate G2
\* Aim: LeafProofNode.java:85, NonLeafProofNode.java:159, DefStepNode.java:60, UseOrHideNode.java:117, LevelNode.java:178, TheoremNode.java:217 BooleanTrueReturn: bad proof leaf, step, DEFINE step and USE in B3..B6; gates G3..G6
\* Aim: InstanceNode.java:132 and UseOrHideNode.java:116 BooleanTrueReturn: B8 extends X8 whose top-level USE, and B9 extends X9 whose top-level INSTANCE, are incorrect; the second level check of the copied node must stay false (gates G8, G9; depends on module check order)
\* Aim: ModuleNode.java:920 ORDER_ELSE, Context.java:342/344 EQUAL_ELSE: inner modules must still be level checked (their errors appear)
\* Aim: SymbolTable.java:273 PrimitiveReturns (ModuleName.hashCode): inner modules are checked in context-table order, so the order of the error messages changes
\* Aim: AssumeNode.java:91 RemoveConditional_EQUAL_IF: unnamed ASSUME has no definition; mutant throws
---- MODULE Del03 ----
VARIABLE x
---- MODULE K ----
CONSTANT c
====
---- MODULE B1 ----
VARIABLE v
bad == (v')'
====
---- MODULE B2 ----
VARIABLE v
ASSUME (v')'
====
---- MODULE B3 ----
VARIABLE v
THEOREM T3 == TRUE
  BY (v')'
====
---- MODULE B4 ----
VARIABLE v
THEOREM T4 == TRUE
  <1>1. TRUE
    BY (v')'
  <1>2. QED
    BY <1>1
====
---- MODULE B5 ----
VARIABLE v
THEOREM T5 == TRUE
  <1> DEFINE d == (v')'
  <1> QED
    OBVIOUS
====
---- MODULE B6 ----
VARIABLE v
THEOREM T6 == TRUE
  <1> USE (v')'
  <1> QED
    OBVIOUS
====
---- MODULE O7 ----
---- MODULE B7 ----
VARIABLE v
bad == (v')'
====
====
---- MODULE X8 ----
VARIABLE v
USE (v')'
====
---- MODULE B8 ----
EXTENDS X8
====
---- MODULE X9 ----
VARIABLE v
I9 == INSTANCE K WITH c <- (v')'
====
---- MODULE B9 ----
EXTENDS X9
====
T0 == INSTANCE K WITH c <- 1
T1 == INSTANCE B1 WITH v <- x
T2 == INSTANCE B2 WITH v <- x
T3 == INSTANCE B3 WITH v <- x
T4 == INSTANCE B4 WITH v <- x
T5 == INSTANCE B5 WITH v <- x
T6 == INSTANCE B6 WITH v <- x
T7 == INSTANCE O7
T8 == INSTANCE B8 WITH v <- x
T9 == INSTANCE B9 WITH v <- x
G1 == (LET J == INSTANCE B1 WITH v <- x IN x')'
G2 == (LET J == INSTANCE B2 WITH v <- x IN x')'
G3 == (LET J == INSTANCE B3 WITH v <- x IN x')'
G4 == (LET J == INSTANCE B4 WITH v <- x IN x')'
G5 == (LET J == INSTANCE B5 WITH v <- x IN x')'
G6 == (LET J == INSTANCE B6 WITH v <- x IN x')'
G7 == (LET J == INSTANCE O7 IN x')'
G8 == (LET J == INSTANCE B8 WITH v <- x IN x')'
G9 == (LET J == INSTANCE B9 WITH v <- x IN x')'
====
