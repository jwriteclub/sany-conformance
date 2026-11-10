
---- MODULE Test ----
EXTENDS TestPrelude
op ==
  /\ LET
       RECURSIVE f
       f == f
     IN P
  /\ Q
====
---- MODULE TestPrelude ----
P == TRUE
Q == TRUE
====

