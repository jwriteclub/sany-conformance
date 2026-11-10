
---- MODULE Test ----
EXTENDS TestPrelude
r == [f EXCEPT ![<<1, 1>>, 1] = 0]
====
---- MODULE TestPrelude ----
VARIABLE f
====

