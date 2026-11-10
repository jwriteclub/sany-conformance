
---- MODULE Test ----
EXTENDS TestPrelude
LEMMA TRUE
PROOF BY ONLY
  P,
  Q
  DEFS
    >,
    R
====
---- MODULE TestPrelude ----
x > y == TRUE
P == TRUE
Q == TRUE
R == TRUE
====

