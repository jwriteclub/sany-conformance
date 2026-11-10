
---- MODULE Test ----
EXTENDS TestPrelude
op1 == op(/\, x, /\)
op2 == op(\/, x, \/)
====
---- MODULE TestPrelude ----
VARIABLE x
op(a(_, _), b, c(_, _)) == 0
====

