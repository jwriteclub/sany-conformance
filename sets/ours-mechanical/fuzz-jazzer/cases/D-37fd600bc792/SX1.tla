---- MODULE SX1 ----
Rc == [a |-> x + 1, b |-> x + 2]
Lt == LE2
App(f) == f[1, 2]
       /\ Rc!2 = x + 2
       /\ Lt!1 = 2 * (x + 5)
       /\ Lt!1!<< = x + 5
       /\ Nest!outer(2, 3)!inner = TRUE
       /\ Nest!outer(2, 4)!inner = FALSE
====
