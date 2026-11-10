---- MODULE Target ----
EXTENDS Naturals

RECURSIVE Fact(_)
Fact(n) ==
  IF n = 0 THEN 1 ELSE n * Fact(n - 1)
THEOREM Fact(0) = 1
OBVIOUS

HIDE Fact

a == Fact(3)

b(n) == Fact(n)

f[n \in Nat] == n

==== 
