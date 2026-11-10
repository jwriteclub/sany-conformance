---- MODULE SX1 ----
EXTENDS Naturals, FiniteSets
Q == \A i \in 2..3, n \in {4} : lab(i, j) :: i + j > x
Ch == CHOOSE k \in 1..9 : k >3
       /\ Q!lab= 300
       /\ Lt!1 = * (x + 5)
====
