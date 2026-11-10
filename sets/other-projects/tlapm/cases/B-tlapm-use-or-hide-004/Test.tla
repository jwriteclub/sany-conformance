
---- MODULE Test ----
EXTENDS TestPrelude, M
USE x, MODULE M, 1 + 3
HIDE DEFS MODULE M, -., *, ^+
USE x, MODULE M DEF MODULE M, x
====
---- MODULE TestPrelude ----
-. a == 0
a + b == 1
a * b == 2
a ^+ == 3
x == 4
====
---- MODULE M ----
====

