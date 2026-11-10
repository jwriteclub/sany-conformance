---- MODULE Target ----
EXTENDS Naturals

---- MODULE Inner ----
CONSTANT cc
f == cc
====

A == LET B == INSTANCE Inner WITH cc <- 1 IN B!f

C == LET D == 2 E == INSTANCE Inner WITH cc <- D IN E!f

F == LET G == LET H == INSTANCE Inner WITH cc <- 3 IN H!f IN G

I == INSTANCE Inner WITH cc <- 4

J == LET K == 5 IN K

P == LET R(x) == INSTANCE Inner WITH cc <- x IN R(6)!f

Z == A + C + F + I!f + J + P
====
