---- MODULE Target ----
A(x, y) == x + y
f(g(_)) == g(1)
res == f(A)
====
