
---- MODULE Test ----
EXTENDS TestPrelude
op == {
  x \in
    /\ A
    /\ B
    /\ C
      :
    /\ D
    /\ E
    /\ F
}
====
---- MODULE TestPrelude ----
VARIABLES A, B, C, D, E, F
====

