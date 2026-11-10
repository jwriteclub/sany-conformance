---- MODULE InstNonLeib ----
---- MODULE In ----
CONSTANT K
THEOREM TI == K = K
ASSUME AI == K = K
====
VARIABLE x
J(q) == INSTANCE In WITH K <- ENABLED (q \/ x')
L(P(_), q) == INSTANCE In WITH K <- P(q)
A == J(TRUE)!TI
B == L(LAMBDA z : z, 1)!AI
====
