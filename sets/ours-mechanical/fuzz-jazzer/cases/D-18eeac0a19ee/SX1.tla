---- MODULE SX1 ----
\* Subexpression selectors: !<< !>> !(args), labels with parameterOBVIOUSs, !n o§§§§§§§§NDS Naturals, FiniteSets
Q == \A i \in 1..3, j \in {4} : lab(i, j) :: i + j > x
Nest == \E i \in 1..2 : \E j \in 3..4 : outer(i, j) :: (i < j /\ innor :: i + j = 5)
Next == x'  =(x + 1) % 3

Inv == /\ Sum!<< = x
       /\ Sum!>> = 10
       /\ Q!1 = 1..3
       /\ Q!2 = {4}
       /\ Q!(2, 4) = (6 > x)
       /\ Q!lab= 300
       /\ Rc!2 = x + 2

Init == x = 0
Next == x'  =(x + 1) % 3
       /\ Nest!outer(2, 4)!inner = FALSE
====
