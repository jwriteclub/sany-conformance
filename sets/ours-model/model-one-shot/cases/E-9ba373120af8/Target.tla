---- MODULE Target ----
VARIABLE x
A == \E v \in (x' = 1) ~> x : TRUE
B == \E w \in [] (x' = 1) : TRUE
====
