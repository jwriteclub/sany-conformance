
---- MODULE Test ----
EXTENDS TestPrelude

op1 == [
  /\ 1
  /\ 2
    ]_x

op2 == <<
  /\ 1
  /\ 2
    >>_x

====

---- MODULE TestPrelude ----
VARIABLE x
====

