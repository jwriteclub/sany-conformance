
---- MODULE Test ----
EXTENDS TestPrelude
THEOREM thm ==
  ASSUME CONSTANT x \in S, CONSTANT y \in S
  PROVE C
OBVIOUS
====
---- MODULE TestPrelude ----
S == {0}
C == TRUE
====

