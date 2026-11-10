---- MODULE SX1 ----
Q == \A i \in 1..3, j \in {4} : lab(i, j) :: i + j > xy * 2
Ite == IF x > 1 THEN x ELSE 42
App(f) == f[1, 2]
       /\ Q!lab= 300
       /\ Rc!2!ou/\ Ite!1 = (x > 0)
       /\ Nest!outer(1, 4) = (1 < 4 /\ 1 + 4 = 5)
       /\ Nest!outer(2, 3)!inner = TRUE
       /\ Nest!outer(2, 4)!inner = FALSE
====
