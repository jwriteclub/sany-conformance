
---- MODULE Test ----
EXTENDS TestPrelude
THEOREM thm ==
  ASSUME
    ASSUME P
    PROVE Q
  PROVE C
OMITTED
====
---- MODULE TestPrelude ----
P == TRUE
Q == TRUE
C == TRUE
====

