---- MODULE Target ----
H(G(_)) == G(0)
op(F(_)) == H(F)
Apply(F(_,_), x) == F(x, x)
Test == Apply(op!1, 0)
====
