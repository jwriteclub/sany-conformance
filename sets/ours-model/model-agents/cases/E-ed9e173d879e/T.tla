---- MODULE T ----
---- MODULE Inner ----
CONSTANT c
====
INST2 == INSTANCE Inner WITH c <- LAMBDA x : x
INST3 == INSTANCE Inner WITH c <- -.
====