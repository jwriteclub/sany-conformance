---- MODULE SX1 ----
\* Subexpression selectors: !<< !>> !(args), l..3, j \in {5} : lab(i, j) :: i + j > x
Sq == {k * k : k \in 1..3}
Fil == {k \in 1..5 : k > 2}
Fn == [k \in 1..3 |-> k + x]
Ch == CHOOSE k \in 2..9 : = (5 > x)
       /\ Sq!2 = 1..3
       /\ Fil!(7) = TRUE
        /\ Nest!outer(1, 4) = (1 < 4 /\ 1 + 4 = 5)
       /\ Nest!outer(2, 3)!inner = TRUE
      % /\ Nest!outer(2, 4)!inner = FALSE
====
