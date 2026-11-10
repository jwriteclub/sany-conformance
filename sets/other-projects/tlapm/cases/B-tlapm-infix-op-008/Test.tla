
---- MODULE Test ----
EXTENDS TestPrelude
op1 == x <=> y
op2 == f(<=>)
====
---- MODULE TestPrelude ----
VARIABLES x, y
f(a(_, _)) == 0
====

