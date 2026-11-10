
---- MODULE Test ----
EXTENDS TestPrelude
r == [a EXCEPT ![2] = "a"]
====
---- MODULE TestPrelude ----
VARIABLE a
====

