---- MODULE Target ----
EXTENDS Naturals
RECURSIVE f
RECURSIVE g(_)
LEMMA
  LET f[x ∈ Nat] ≜ IF x = 0 THEN 0 ELSE f[x - 1]
  IN f[3]

g[y ∈ Nat] ≜ y

Init ≜ A /\ g[1]
====
