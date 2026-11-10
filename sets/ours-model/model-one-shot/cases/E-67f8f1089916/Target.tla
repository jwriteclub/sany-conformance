---- MODULE Target ----
CONSTANT q(_)
VARIABLE v
---- MODULE Inner ----
CONSTANT F(_), x
====
Id(a) == 1
I == INSTANCE Inner WITH F <- Id, x <- q(v')
========================
