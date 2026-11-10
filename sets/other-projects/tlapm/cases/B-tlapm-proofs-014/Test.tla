
---- MODULE Test ----
EXTENDS TestPrelude
PROPOSITION P => Q
PROOF
  <1> P
  <1> Q
    PROOF BY P
  <1> QED
=====================
---- MODULE TestPrelude ----
VARIABLES P, Q
====

