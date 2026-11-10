
---- MODULE Test ----
EXTENDS TestPrelude
r == [a EXCEPT !.x.y.z = "b"]
====
---- MODULE TestPrelude ----
VARIABLE a
====

