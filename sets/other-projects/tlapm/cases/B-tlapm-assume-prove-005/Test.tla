
---- MODULE Test ----
EXTENDS TestPrelude
THEOREM thm ==
  ASSUME NEW CONSTANT x \in S
  PROVE C
OBVIOUS
====
---- MODULE TestPrelude ----
S == {0}
C == TRUE
====

