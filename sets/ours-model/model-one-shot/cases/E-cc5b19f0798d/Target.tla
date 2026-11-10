---- MODULE Target ----
EXTENDS TLC
A(f(_), a) == f(a) + Len(a)
B(g(_), b) == A(g, b)

C == B(Len, "x") /\ B(Len, "y")
====
