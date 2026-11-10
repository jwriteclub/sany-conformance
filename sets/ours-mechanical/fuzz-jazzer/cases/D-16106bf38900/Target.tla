---- MODULE Target ----
---- MODULE Inner ----
CONSTANT F(_)
Assumption(BOOLEAN) == F(BOOLEAN)
====
INSTANCE Inner WITH F <-⊙====
