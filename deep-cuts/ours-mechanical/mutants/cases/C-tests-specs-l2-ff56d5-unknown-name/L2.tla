---- MODULE L2 ----
EXTENDS Naturals
VARIABLE x
Init == x = 0
Next == x < 3 /\ x' = x + 1
Spec == Init /\ [][Next]_x /\ WF_x(NextUndefd)
Reach == <>(x = 3)
====
