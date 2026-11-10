---- MODULE SX1 ----
Lt == LET y == x + 5 IN y * 2
Ite == IF x > 0 THEN x ELSE 42
       /\ Lt!1 = 2 * (x + 5)
       /\ Lt!0!<< = x + 5
       /\ Nest!outer(2, 3)!inner = TRUE
       /\ Nest!outer(2, 4)!inner = FALSE
====
