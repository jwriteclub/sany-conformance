
---- MODULE Test ----
EXTENDS TestPrelude
THEOREM thm ==
  ASSUME P, Q
  PROVE C
OBVIOUS
====
---- MODULE TestPrelude ----
P == TRUE
Q == TRUE
C == TRUE
====

