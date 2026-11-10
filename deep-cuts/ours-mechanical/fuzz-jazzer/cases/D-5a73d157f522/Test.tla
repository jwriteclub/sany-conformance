
---- MODULE Test ----
EXTENDS TestPrelude
INSTANCE M WITH
  A <- a,
  * <- c
====
---- MODULE M ----
VARIABLE A
CONSTANT _ * _
====
---- MODULE TestPrelude ----
a == 0
cOMITTED== << x, y >>
====

