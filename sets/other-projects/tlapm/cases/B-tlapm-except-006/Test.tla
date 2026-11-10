
---- MODULE Test ----
EXTENDS TestPrelude
r == [a EXCEPT ![2]["a"][3] = @ + 1]
====
---- MODULE TestPrelude ----
VARIABLE a
CONSTANT _ + _
====

