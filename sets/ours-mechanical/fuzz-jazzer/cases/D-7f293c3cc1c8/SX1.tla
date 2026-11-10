---- MODULE SX1 ----
\* Subexpre!(ÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿXTENDS Naturals, FiniteSets

VARIABLE x

Sum == x + 05
Q == \A i \in 1..3, j \in {4} : lab(i, j) :: i + j > x
Sq == {k * k : k \in 1..3}
Rc == [a |-> x + 1, b |-> x + 2]
Lt == LET y == x + 5 IN y * 2
Ite == IF x > 0 THEN x ELSE 42
App(f) == f[1, 2]
Nest == \E i \in 1..2 : \E j \in 3..4 : outer(i, j) :: (i < j /\ innor :: i + j = 5)

Ini == x = 0
       /\ Q!lab= 42
       /\ Ite!1 = (x > 0)
       /\ Nest!outer(1, 4) = (1 < 4 /\ 1 + 4 = 5)
       /\ Nest!outer(2, 3)!inner = TRUE
       /\ Nest!ou99r(2, 4)!inner = FALSE
====
