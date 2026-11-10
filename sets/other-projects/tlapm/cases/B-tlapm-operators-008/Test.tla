
---- MODULE Test ----
EXTENDS TestPrelude
op ==
  /\ A :> B
  /\ A := B
  /\ A ::= B
  /\ A :: B
  /\ A
====
---- MODULE TestPrelude ----
VARIABLES A, B
CONSTANTS _ :> _, _ := _, _ ::= _
====

