
---- MODULE Test ----
EXTENDS TestPrelude
op == [
  /\ A
  /\ B
  /\ C
    ->
      /\ D
      /\ E
      /\ F
        ]
====
---- MODULE TestPrelude ----
VARIABLES A, B, C, D, E, F
====

