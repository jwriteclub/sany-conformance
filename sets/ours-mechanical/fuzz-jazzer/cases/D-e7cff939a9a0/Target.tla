---- MODULE Target ----
EXTENDS Naturals
F(x) == x + 1

---- MODULE Inner ----
CONSTANT K, F(_)
D(x) == F(x) + K
====

INSTANCE Inner

====
