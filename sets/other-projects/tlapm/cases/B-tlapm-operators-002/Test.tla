
---- MODULE Test ----
EXTENDS TestPrelude
op1 == -. (1)
op2 == - (1)
op3 == - 1 - 2
====
---- MODULE TestPrelude ----
-. x == 0
x - y == 1
====

