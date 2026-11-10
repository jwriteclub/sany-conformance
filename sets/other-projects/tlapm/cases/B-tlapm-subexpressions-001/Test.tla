
---- MODULE Test ----
EXTENDS TestPrelude
op == M!X!Y!Z
====
---- MODULE TestPrelude ----
M == INSTANCE M1
====
---- MODULE M1 ----
X == INSTANCE M2
====
---- MODULE M2 ----
Y == INSTANCE M3
====
---- MODULE M3 ----
Z == 0
====

