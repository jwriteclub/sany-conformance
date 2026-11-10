
---- MODULE Test ----
EXTENDS TestPrelude
op ==
  /\ 1
  /\ foo.bar 

====
---- MODULE TestPrelude ----
foo == [ bar : 0 ]
====

