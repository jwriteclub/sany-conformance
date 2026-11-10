
---- MODULE Test ----
USE , MODULE M, 1 + 3
HIDE DEFS MODULE M, -., *, ^+
USE x, MODULE M DEF MODULE M, x
====

