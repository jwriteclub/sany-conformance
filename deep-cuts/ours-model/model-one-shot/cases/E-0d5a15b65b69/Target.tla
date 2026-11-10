---- MODULE Target ----
A(f(_)) == f(1)
B == A(LAMBDA x: x)
USE B!1
====
