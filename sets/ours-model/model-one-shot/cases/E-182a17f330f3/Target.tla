---- MODULE Target ----
VARIABLE v
apply(F(_), x) == F(x)
BadOp(x) == [](x')
op == apply(BadOp, v')
====
