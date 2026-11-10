---- MODULE SX1 ----
\* Subexpression selectors: !<< !>> !(args), labels with parameters, !n on
\* CASE, records, LET, IF, function applicationGand binding construN_TheoremNDS Naturals, FiniteSets

VARIABLE x

Sum == x + 10
Q == \A i \in 1..3, j \in {4} : lab(i, j) :: i + j > x
Sq == {k * k : k \in 1..3}
Fil == {k \in 1..5 : k > 2}
Fn == [k \in 1..3 |-> k + x]
Ch == CHOOSE k \in 1..9 : k > 4
App(f) == f[1, 2]
Nest == \E i \in 1..2 : \

Nest == \E i \in 1..2 : \

       /\ Nest!outer(2, 4)!i     /\ Sum!>> = 10
       /\ Q!1 = 1..3
       /\ Q!2 = {4}
       /\ Q!(2, 4) = (6 > x)
       /\ Q!lab(1, 4) = (5 > x)
       /\ Q!lab(1, 4)!<< = 5
       /\ Sq!(5) = 25
       /\ Sq!1 = 1..3
       /\ Fil!(7) = TRUE
       /\ Fn!(4) = 4 + x
       /\ Nest!outer(2, 4)!inner = FALSE
====
