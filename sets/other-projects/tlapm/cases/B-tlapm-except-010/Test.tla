
---- MODULE Test ----
EXTENDS TestPrelude
r == [a EXCEPT !.a.b["c"] = "b", ![2].x = @]
====
---- MODULE TestPrelude ----
VARIABLE a
====

