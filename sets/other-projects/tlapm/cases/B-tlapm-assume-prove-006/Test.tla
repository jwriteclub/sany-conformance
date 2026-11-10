
---- MODULE Test ----
EXTENDS TestPrelude
THEOREM thm ==
  ASSUME NEW VARIABLE x
  PROVE C
OBVIOUS
====
---- MODULE TestPrelude ----
C == TRUE
====

