---- MODULE Target ----
VARIABLE v
Next == v' = v
A == [Next]_v
B == [](v = 0)
THEOREM A /\ B
====
