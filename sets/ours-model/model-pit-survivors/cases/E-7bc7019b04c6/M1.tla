---- MODULE M1 ----
EXTENDS Naturals
CONSTANT c
H(g(_)) == g(c)
A == H(LAMBDA y: y + c)
====
