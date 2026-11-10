---- MODULE Target ----
EXTENDS Naturals
A ≜
rals
RECURSIVE f
RECURSIVE g(_)
A ≜
  LET f[x ∈ Nat] ≜ IF x = 0 THEN 0 ELSE f[x - 1]
  IN f[3]/\  LET f[x ∈ Nat] ≜ IF x = 0 THEN 0 ELSE f[x - 1]
  IN f[3]/\ g[1]
====
