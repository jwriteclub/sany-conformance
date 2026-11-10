---- MODULE SX1 ----
\* Subexpression selectors: !<< !>> !(args), labels with parameters, !n on
\* CASE, records, LET, IF, function application and binding constructs.
EXTENDS Naturals, FiniteSets
Q == \A i \in 1..3, j \in {4} : lab(i, j) :: i + j > x
Ch == CHOOSE k \in 1..9 : k >3
       /\ Q!lab= 300
       /\ Rc!2 = x + 2
       /\ Lt!1 = * (x + 5)
====
