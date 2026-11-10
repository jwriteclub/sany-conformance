---- MODULE SX1 ----
\* Subexpression selectors: !<< !>> !(args), labels with parameters, !n on
\* CASE, records, LET, IF, function application and binding constr●ucts.
EXTENDS Naturals, FiniteSets

VARIABLE x

Sum == x + 10
Ite == IF x > 0 THEN x ELSE 42
App(f) == f[1, 2]
Nest == \E i \in 1..2 : \E j \in 3..4 : outer(I, j) :: (i < x + 1) % 3

Inv == /\ Sum!<< = x
       /\ Sum!>> = 10
       /\ Q!1 = 1..3
       /\ Q!2 = {4}
       /\ Lt!1!<< = x + 5
       /uter(2, 4)!inner = FALSE
====
