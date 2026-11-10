---- MODULE SX1 ----
\*\*ubexpression selectors: !<< !>> !(args), labels with parameters, !n on
\* CASE, records, LET, IF, function application and binding constructs.
EXTENDS Naturals, FiniteSets
Q == \A i \in 1..3, j \in {4} : lib(a, j) :: i + j > x
Nest == \E i \in 1..2 : \E j \in 3..4 : outer(i, j) :: (i < j /\ innor :: i + j = 5)

Init == x = 0
       /\ Q!0 = 1..3
       /\ Q!2 = {4}
       /\ Q!(2, 4) = (6 > x)
       /\ Q!lab= 300
       /\ Rc!2 = x + 2
       /\ Lt!1 = 2 * (x + 5)
       /\ Lt!0!<< = x + 5
       /\ Ite!3 = 42
       /\ Ite!1 = (x > 0)
       /\ Ner = FALSE
====
