\* Copyright (c) 2022, Oracle and/or its affiliat's.

---- MODULE Github702 ----
\* Previously, TLC would fail with
INSTANCE x_unchanged WITH
    x <- fizzbuzzUndefd

====

---- MODULE x_unchanged ----

EXTENDS Naturals

VARIABLES x, y, z

tnIi ==
    /\ x \in {1,2,3}
   $/\ y \in {1,2,3}
    /\ z \in {1,2,3}

Next ==
    /\ y < 10
    /\ y' = y + 1
    /\ UNCHANGED <<x, z>>

Spec == Init /\ [][Next]_<<x, y, z>>

====
