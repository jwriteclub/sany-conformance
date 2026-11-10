
---- MODULE Test ----
EXTENDS TestPrelude
op == f(LAMBDA x, y : x + y, LAMBDA x : 0)
====
---- MODULE TestPrelude ----
f(x(_, _), y(_)) == 0
x + y == 1
====

