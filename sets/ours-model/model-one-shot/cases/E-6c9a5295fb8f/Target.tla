---- MODULE Target ----
VARIABLE v
apply(F(_), x) == F(x)
op == apply(', v' /\ []v)
====
