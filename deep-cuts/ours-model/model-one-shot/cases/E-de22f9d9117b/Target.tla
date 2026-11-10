---- MODULE Target ----

---- MODULE Inner ----
====

g == INSTANCE Inner WITH f <- LAMBDA x, y : x + y
====
