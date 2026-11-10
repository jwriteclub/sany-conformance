---- MODULE Target ----
---- MODULE Inner ----
CONSTANT F(_)
Assumption(BOOLEAN) == F(BOOLEAN)
====
G(x) == x
INSTANCE Inner WITH F <- G
====
