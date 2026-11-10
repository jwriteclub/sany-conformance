
---- MODULE Test ----
EXTENDS TestPrelude
THEOREM thm ==
  ASSUME NEW TEMPORAL f(_, _)
  PROVE C
OMITTED
====
---- MODULE TestPrelude ----
C == TRUE
====

