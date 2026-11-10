
---- MODULE Test ----
EXTENDS TestPrelude
op == [
  x, y \in Nat, <<a, b>> \in
    /\ A
    /\ B
      |->
        /\ A
        /\ B
            ]
====
---- MODULE TestPrelude ----
CONSTANT Nat
VARIABLES A, B
====

