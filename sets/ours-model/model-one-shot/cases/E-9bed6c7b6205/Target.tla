---- MODULE Target ----
VARIABLES v, w
apply(F(_, _), x, y) == F(x, y)
Bad(a, b) == a' /\ b
op == apply(Bad, v, w)
====
