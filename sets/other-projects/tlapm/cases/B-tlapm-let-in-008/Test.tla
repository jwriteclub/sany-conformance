
---- MODULE Test ----
EXTENDS TestPrelude
op ==
  \/ LET M1 == INSTANCE M2 IN P
  \/ Q
====
---- MODULE M2 ----
====
---- MODULE TestPrelude ----
P == TRUE
Q == TRUE
====

