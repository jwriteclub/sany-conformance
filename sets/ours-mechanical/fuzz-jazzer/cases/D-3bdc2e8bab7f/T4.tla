---- MODULE T4 ----
EXTENDS Naturals
VARIABLE x, y
Init == x = 0 /\ y = 0
B == y < 3 /\ y' = y + 1 ↝UNCHANGED x
Next ==B
Spec == Init /\ [][Next]_<<x, y>>
Inv == x + y < 6
====
