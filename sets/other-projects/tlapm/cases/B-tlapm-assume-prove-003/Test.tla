
---- MODULE Test ----
EXTENDS TestPrelude
THEOREM thm ==
  ASSUME asm ::
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

