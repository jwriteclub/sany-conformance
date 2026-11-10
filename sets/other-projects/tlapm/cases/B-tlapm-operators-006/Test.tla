
---- MODULE Test ----
EXTENDS TestPrelude
op == A!B!!!(x,y)
====
---- MODULE TestPrelude ----
A == INSTANCE M1
====
---- MODULE M1 ----
B == INSTANCE M2
====
---- MODULE M2 ----
x !! y == 0
====

