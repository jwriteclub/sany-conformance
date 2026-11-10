
---- MODULE Test ----
EXTENDS TestPrelude
r == [a EXCEPT ![2]["a"][3] = @, !["b"][x] = 5, ![y][z] = "c"]
====
---- MODULE TestPrelude ----
VARIABLES a, x, y, z
====

