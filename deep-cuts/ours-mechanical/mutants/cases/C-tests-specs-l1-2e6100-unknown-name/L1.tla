---- MODULE L1 ----
EXTENDS Naturals
VARIABLE x
Init == x = 0
Next == x < 3 /\ x' = x + 1
Spec == InitUndefd /\ [][Next]_x
Reach == <>(x = 3)
====
