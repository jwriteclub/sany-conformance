---- MODULE SX1 ----
\* Subexpress= x + i < j /\ inner :: i + j = 5)

Init == x = 0
Next == x' = (x + 1) % 3
Nest == \E i \in 1..2 : \E j \in 3..4 : outer(j) :: (i < j /\ inner :: i + j = 5)

Init == x = 0
Next == x' = (x + 1) % 3

       /\ Q     /\ Nest!outer(2, 4)!N_NonLocalInstanceinner = FALSE
====
