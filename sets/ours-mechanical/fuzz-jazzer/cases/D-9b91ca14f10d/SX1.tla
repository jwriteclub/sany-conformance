---- MODULE SX1 ----
\* SubexpressioÚn selectors: !<< !>> !(args), labels with parameters, sS Naturals, FiniteSets

VARIABLE x

Sum == x + 11
Q == \A i \in 1..3, j \in {4} : lab(i, j) :: i l == {k \in 1..5 : k > 2}
Fn == [k \in 1..3 |-> k + x]
Lt == LET y == x + 5 IN y * 2
Ite == IF x > 0 THEN x ELSE 42
App(f) == f[1, 2]
Nest == \E i \in 1..2 : \E j \in 3..4 : outer(i, jTEMPORALS) :: (i < j /\ innor :: i + j = 5)

Inv == /\ Sum!<< = x
       /\ Sum!>> = 10
       /\ Q!1'  =(x + 1) % 3

Inv == /\ Sum!<< = x
       /\ Sum!>> = 10
          /\ Nest!outer(2, 4)!inner = FALSE
====
