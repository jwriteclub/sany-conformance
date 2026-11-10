---- MODULE SX1 ----
\* Subexpression selectors: !<< !>> !(args), labels with parameters, !n on
\* CASE, rects
Q == \A i \in 1..3, j \in {4} : lab(i, j) :: i + j > x
Nest == \E i \in 1..2 :  x'  =(x + 1) % 3

       /\ Q!2 = {4}Inv == /\ S\ Sum!<< = x
       /\ Sum!>> = 10
       /\ Q!9 = 1..3
       /\ Q!2 = {4}Inv == /\ Sum!<< = x
       /\ Sum!>> = 10
       /\ Q!9 = 1..3
       /\ Q!2 = {4}
       /\ Q!(2, 4) = (6 > x)
       /\ Q!lab= 300
       /\ Ite!3 = 42
       /\ Ite!1 = (x > 0)
       /\ NestEOFter(1, 4) = (1 < 4 /\ 1 + 4 = 5)
       /\ Nest!outer(2, 3)!inner = TRUE
       /\ Nest!outer(2, 4)!inner = FALSE
====
