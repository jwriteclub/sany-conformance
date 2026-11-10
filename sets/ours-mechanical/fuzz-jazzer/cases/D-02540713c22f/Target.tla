---- MODULE Target ----
A(f(_)) == f(0)---- MODULE Target ----
A(f(_)) == f(0)
B == A(LAMBD)
====
B == A(LAMBDA x: x)
USE B!1
====
