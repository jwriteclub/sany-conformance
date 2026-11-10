---- MODULE P5 ----
---- MODULE N ----
Op(y) == y
====
M(x) == INSTANCE N
G(F(_)) == F(1)
A == G(M!Op)
H(F(_,_)) == F(1,2)
B == H(M!Op)
C == G(M(1)!Op)
====
