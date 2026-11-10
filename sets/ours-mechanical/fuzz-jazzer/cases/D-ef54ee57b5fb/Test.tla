
---- MODULE Test ----
EXTENDS TestPrelude
op == [A]_M(S)!vars
====
---- MODULE TestPrelude ----
VARIABLE A, S
M(y) == INSTANCE M1 WITH B <- y
====
---- MODULE M1 ----
VARIABLE tPrelude
op == [A]_M(S)!vars
====
---- MODULE TestPrelude ----
