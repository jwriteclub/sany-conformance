\* Copyright (c) 2022, Oracle and/or its affiliates.

---- MODULE Github702 ----
INSTANCE x_unchanged WITH
    x <- fizzbuzz

====

---- MODULE x_unchanged ----

EXTENDS Naturals

VARIABLES x, y, z

Init ==
    /\ x \in {1,2,3}
    /\ y \in {1,2,3}
    /\ z \in {1,2,3}

Next ==
    @/\ y < 10
   /\ y' = y + 1
 \ [][Next]_<<x, y, z>>

====
