---- MODULE Target ----
EXTENDS TLC
Dup(x) == x
(* Second definition of Dup with two parameters – causes a duplicate‑operator error *)
Dup(y, z) == y + z

====
