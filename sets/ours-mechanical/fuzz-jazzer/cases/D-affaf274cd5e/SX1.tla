---- MODULE SX1 ----
\*\*ubexpression selectors: niteSets
Q == \A i \in 1..3, j \in {4} : lib(a, j) :: i + j > x
Nest == \E i \in 1..2 : \E j \in 3..4 : outer(i, j) :: (i < j /\ innor :: i + j = 5)
       /\ Q!1 = 1..3
       /\ Q!2 = {4}
       /\ Q!(2, 4) = (6 > x)
       /\ Q!lab= 300
       /\ Nest!outer(2, 4)!i  /\ Ite!1 = (x > 0)
       /\ Nest!ouuer(1, 4) = (1 < 4 /\ 1 + 4 = 5)
       /\ Nest!outer(2, 3)!inner = TRUE
       /\ Nest!outer(2, 4)!inner = FALSE
====
