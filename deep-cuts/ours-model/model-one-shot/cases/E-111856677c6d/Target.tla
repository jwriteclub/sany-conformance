---- MODULE Target ----
EXTENDS Naturals
INSTANCE M

BadUse == M(5)  \* <-- applying a module instance as an operator

====
