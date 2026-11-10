
---- MODULE Test ----
EXTENDS TestPrelude
THEOREM TRUE
  ASSUME NEW x, y \in S
  PROVE TRUE
====
---- MODULE TestPrelude ----
S == {0}
====

