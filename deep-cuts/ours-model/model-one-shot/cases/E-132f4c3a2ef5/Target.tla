---- MODULE Target ----
Op(x) == x
---- MODULE Inner ----
RECURSIVE Op(_)
Op(x) == x
====

INSTANCE Inner
====
