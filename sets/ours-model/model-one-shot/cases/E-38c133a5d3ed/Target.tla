---- MODULE Target ----
THEOREM T == TRUE
PROOF OMITTED
USE Naturals

VARIABLES x
CONSTANTS c
RECURSIVE Fact(_)
f[x \in {1}] == x
I == INSTANCE Integers
INSTANCE Naturals
ASSUME TRUE
====
====
