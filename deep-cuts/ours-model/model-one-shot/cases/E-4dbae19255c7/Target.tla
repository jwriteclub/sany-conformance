---- MODULE Target ----
EXTENDS Naturals

CONSTANT K

F(x) == x + 1

---- MODULE Inner ----
CONSTANT K, F(_)
D(x) == F(x) + K
====

---- MODULE Inner2 ----
CONSTANT K
E == K + 1
====

INSTANCE Inner

====
