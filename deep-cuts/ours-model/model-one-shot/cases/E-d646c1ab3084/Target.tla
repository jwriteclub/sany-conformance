---- MODULE Target ----
Apply(F(_), x) == F(x)
SomeOp(G(_)) == Apply(G, 0)
Test == Apply(SomeOp!1, 1)
====
