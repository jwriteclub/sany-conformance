
---- MODULE Test ----
EXTENDS TestPrelude
op == CHOOSE <<x, y, z>> \in S : FALSE
====
---- MODULE TestPrelude ----
VARIABLE S
====

