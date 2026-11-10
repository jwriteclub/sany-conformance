---- MODULE Target ----
---- MODULE Inner ----
CONSTANT c
ASSUME A == c < 10
====

LOCAL INSTANCE Inner
x == TRUE
====
