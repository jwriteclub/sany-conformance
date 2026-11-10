
---- MODULE Test ----
EXTENDS TestPrelude
r == [a EXCEPT !.a.b.c = @, !.b.x = 5]
====
---- MODULE TestPrelude ----
VARIABLE a
====

