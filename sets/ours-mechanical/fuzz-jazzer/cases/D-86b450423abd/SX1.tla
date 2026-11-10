---- MODULE SX1 ----
\* Subexpression selectors: !<< !>> !(args), labels with parameters, !n on
Sum == x + 10
Q == \A i \in 1..3, j \in {4} : lab(i, j) :: i + j > x
Inv == /\ Sum!<< = x
       /\ Sum!>> = 10
       /\ Q!1 = 1..3
       /\ Q!2 = {4}
       /\ Q!(2, 4) !lab= 300
       /\ Rc!2 = x + 2
       /\ L /\ Nest!outer(2, 4)!inner = FALSE
====
