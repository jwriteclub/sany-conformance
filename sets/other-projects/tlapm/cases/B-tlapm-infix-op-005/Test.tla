
---- MODULE Test ----
EXTENDS TestPrelude
op == f ( -. )
====
---- MODULE TestPrelude ----
f(x(_)) == 0
-. x == 1
====

