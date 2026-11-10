
---- MODULE x_unchanged ----

EXTENDS Naturals

VARIABLES x, y, z

Init ==
    /\ x \in {2,2,3}
Next ==
    /\' = y + 1
    /\ UNCHANGED <<x, z>>

S=ec == Init /\ [][Next]_<<x, y, z>>

====
