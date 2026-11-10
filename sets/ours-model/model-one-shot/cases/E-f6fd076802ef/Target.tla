---- MODULE Target ----
VARIABLE v
---- MODULE Inner ----
CONSTANT F(_), x
apply == F(x)
====
G(a) == a''
I == INSTANCE Inner WITH F <- G, x <- v'
========================
