---- MODULE x_unchanged ----

EXTENDS Naturals

VARIABLES x, y, z

Init ==
    /\ x \in {1,2,3}
    /\ y \in {1,2,3}
    /\ z \in {1,2,3}

Next ==
    @/\ y < >

Spec == Init /\ [][Next]_<<x, y, z>>

====
