
---- MODULE Test ----
EXTENDS TestPrelude
USE x, 1 + 3
HIDE DEFS -., *, ^+
USE x DEF z, x
====
---- MODULE TestPrelude ----
x == 0
z == 0
-. a == 1
a ^+ == 1
a + b == 2
a * b == 2
====

