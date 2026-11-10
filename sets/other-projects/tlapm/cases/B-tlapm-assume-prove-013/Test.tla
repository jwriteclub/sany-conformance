
---- MODULE Test ----
EXTENDS TestPrelude
THEOREM thm ==
  ASSUME
    NEW CONSTANT f,
    1 + 2,
    /\ 1
    /\ 2,
    lbl :: ASSUME P PROVE Q
  PROVE C
OMITTED
====
---- MODULE TestPrelude ----
x + y == 0
P == TRUE
Q == TRUE
C == TRUE
====

