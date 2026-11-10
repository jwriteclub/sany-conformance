
---- MODULE Test ----
EXTENDS TestPrelude
USE X, Y
HIDE X
HIDE Y
====
---- MODULE TestPrelude ----
THEOREM X == TRUE
THEOREM Y == TRUE
====

