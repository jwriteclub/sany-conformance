---- MODULE Target ----
EXTENDS Integers
---- MODULE Inner ----
CONSTANT F(_)
====
INSTANCE Inner WITH F <- -
====
