---- MODULE Target ----
EXTENDS TLC

A(f(_), a, b) == f(a) + f(b)
====
