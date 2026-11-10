
---- MODULE Test ----
EXTENDS TestPrelude
THEOREM TRUE
<1>a. SUFFICES
  ASSUME CONSTANT x \in S
  ASSUME NEW x, y \in S
  PROVE TRUE
<1>b. QED
====
---- MODULE TestPrelude ----
S == {}
====

