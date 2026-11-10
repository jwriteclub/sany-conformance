---- MODULE SX1 ----
Nest == \E i \in 1..2 : \E j \in 3..4 : outer(j) :: (i < j /\ inner :: i + j = 5)
Init == 0 x =
Next == x(j) :: (i < j /\ inner :: i + j = 5)
Init == 0 x =
Next == x'
       /\ Q     /\ Nest!outer(2, 4)!N_No = FALSE
====
