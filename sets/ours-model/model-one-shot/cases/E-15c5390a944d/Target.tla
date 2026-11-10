---- MODULE Target ----
B(f(_)) == f(1)
A == B(LAMBDA x : x)
USE A!1!(1, 2)
====
