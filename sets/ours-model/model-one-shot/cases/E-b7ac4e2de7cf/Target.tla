---- MODULE Target ----
---- MODULE Inner ----
CONSTANT F(_)
====

bogus == undefinedThing

op == 0

INSTANCE Inner WITH F <- op
====
