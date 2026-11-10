---- MODULE Target ----
EXTENDS Naturals
CONSTANT N
VARIABLE x
---- MODULE Inner ----
InnerOp == 2
====
RECURSIVE Fact(_)
Fact[n \in Nat] == IF n = 0 THEN 1 ELSE n * Fact[n-1]
THEOREM Thm == 1 = 1
====
