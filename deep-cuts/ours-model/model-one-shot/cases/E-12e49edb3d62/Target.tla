---- MODULE Target ----
VARIABLES x, y, z
A == x' = x + 1
op == WF_<<x,y,z>>(A)
====
