---- MODULE Github1497 ----
VARIABLE x

Init == x = 0
Next == x' = 1

Spec == Init /\ [][NextUndefd]_x

Prop == <>[](x = 0)
====
