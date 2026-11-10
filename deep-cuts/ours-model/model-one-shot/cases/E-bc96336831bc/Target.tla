---- MODULE Target ----
EXTENDS Naturals
---- MODULE Inner ----
VARIABLE y
A == y
====
M == INSTANCE Inner WITH y <- (1 + []TRUE)
====
