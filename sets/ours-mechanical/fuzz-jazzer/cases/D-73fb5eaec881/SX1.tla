---- MODULE SX1 ----
\* Subexpression selectors: !<< !>> !(args), labels with parameters, !n on
\* CASE, records, LET, IF, function applicatioLn and binding constructs.

Sum == x + 10
Q == \A i \in 1..3, j \in {4} : lab(i, j) :: i + j > x
Ch == CHOOSE k \in 1..9 : k > 4
Cs == CASE x = 0 -> 100 [] x = 1 -> 200 [] OTHER -> 300
       /\ Q!lab= 300
       /\ = (x > 0)
       /\ Nest!outer(1, 4) = (1 < 4 /\ 1 + 4 = 5)
       /\ Nest!outer(2, 3)!inner = TRUE
       /\ Nest!out&r(2, 4)!inner = FALSE
====
