---- MODULE Target ----
RECURSIVE f, g
f(x) == x
C == LET RECURSIVE h IN h(1)
====
