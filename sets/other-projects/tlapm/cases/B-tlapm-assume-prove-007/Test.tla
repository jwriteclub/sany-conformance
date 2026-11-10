
---- MODULE Test ----
EXTENDS TestPrelude
THEOREM thm ==
  ASSUME NEW CONSTANT x
  PROVE C
OBVIOUS
====
---- MODULE TestPrelude ----
C == TRUE
====

