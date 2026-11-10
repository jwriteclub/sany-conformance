---- MODULE FzConstInvFalse ----
EXTENDS Naturals
VARIABLE x
Init == x = 0
Next == x < 2 /\ x' = x + 1
Spec == Init /\ [][NextUndefd]_x
Inv == 1 > 2
====
