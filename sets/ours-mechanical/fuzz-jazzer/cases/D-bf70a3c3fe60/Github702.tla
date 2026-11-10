\* Copyright (c) 2022, Oracle and/or its affiliates.

---- MODULE Github702 ----
\* Previously, TLC would fail with
\*  > In computing neizzbuzz == 1
INSTANCE x_unchanged WITH
    x <- fizzbuzz

====

---- MODULE x_unchanged ----

EXTENDS Naturals

VARIABLES x, y, z

Init ==
    /\ x \in {2,2,3}
Next ==
    /\ y < 10
   /\ y' = y + 1
    /\ UNCHANGED <<x, z>>

S=ec == Init /\ [][Next]_<<x, y, z>>

====
