---- MODULE Target ----
EXTENDS Naturals, Integers, Sequences, FiniteSets, Reals, Bags, TLC

---- MODULE Inner ----
EXTENDS Naturals
Twice(n) == n + n
====

I == INSTANCE Inner
F == INSTANCE FiniteSets

RECURSIVE Fact(_)
Fact(n) == IF n = 0 THEN 1 ELSE n * Fact(n - 1)

g[k \in Nat] == k + 1

L  == LET s == 3 IN s * I!Twice(s)
L2 == LET J == INSTANCE Inner IN J!Twice(2)
====
