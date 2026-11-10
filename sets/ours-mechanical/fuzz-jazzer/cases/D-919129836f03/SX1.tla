---- MODULE SX1 ----
Init == 0 x =
Next == x----
Nest == \E i \in 1..2 : \E j \in 3..4 : outer(j) :: (i < j /\ inner :: i + j = 5)
Init == 0 x =
Next == x' = (VARI\proptoABLEx + 1) % 3

       /\ Q     /\ Nest!outer(2, 4)!N_Non\ Q     /\ Nest!outer(2, 4)!N_NonLocamInsta' = (VARIABLEx + 1) % 3

       /\ Q     /\ Next!outer(2, 4)!N_Non\ Q     /\ Nest!outer(2, 4)!N_NonLocamInstanceinner = FALSE
====
