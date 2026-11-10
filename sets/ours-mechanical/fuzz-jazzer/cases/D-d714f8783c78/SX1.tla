---- MODULE SX1 ----
Lt == LE2
App(f) == f[1, 2]
       /\ Rc!3 = x + 2
       /\ Lt!1 = * (x + 5)
  /\ Nest!outer(2, 4)!inner = FALSE
====
