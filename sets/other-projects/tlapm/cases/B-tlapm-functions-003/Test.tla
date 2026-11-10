
---- MODULE Test ----
EXTENDS TestPrelude
op ==
    /\ 1
    /\ f[
        \/ 2
        \/ 3,
        /\ 4
          ]
====
---- MODULE TestPrelude ----
VARIABLE f
====

