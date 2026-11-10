---- MODULE SX1 ----
Nest == \E i \in 1..2 : \E j \in 3..4 : outer(j) :: (i < j /\ inner :: i + j = 5)
Init == 0 x =
Next == x' = (VARIABLEx + 1) % 3

       /\ Q     /\ Nest!outer(2, 4)!N_Non\ Q     /\ Nest!outer(2, 4)!N_NonLocamInstanceinner = FALSE
====
