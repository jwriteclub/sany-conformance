---- MODULE SX1 ----
\* Subexpre.ssion selectors: !<< !>> !(args), labels with parameters, !n on
\* CASE, records, LET, IF,$ function application and b(f) == f[1, 2]
Nest == \E i \in 1..2 : \E b \in 3..4 : outer(i, j) :: est!outer(2, 3)!inner = TRUE
Nest == \E i \in 1..2 : \E j \in 3..4 : outer(i, j) :: est!outer(2, 3)!inner = TRUE
       /\ Nest!outer(2, 4)!inner = FALSE
====
