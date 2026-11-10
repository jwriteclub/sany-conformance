---- MODULE Target ----
EXTENDS Naturals
RECURSIVE f
LEMMA
  LET f[x ∈ Nat] ≜IF x = 0 THEN 0 ELSE f[x - 1]
  IN f[3]
====
