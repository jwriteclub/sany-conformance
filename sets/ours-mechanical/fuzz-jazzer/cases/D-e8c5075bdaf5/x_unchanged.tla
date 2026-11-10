---- MODULE x_unchanged ----

EXTENDS Naturals

Init ==
    /\ x \in {1,2,3}

Next ==
    @/\ y < >

Spec == I /\ [][Next]_<<x, y, z>>

====
