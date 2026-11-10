---- MODULE T ----
---- MODULE Inner ----
CONSTANT c
====
I1 == INSTANCE Inner WITH c <- 1 + 2
I3 == INSTANCE Inner WITH c <- LAMBDA x : x
I4 == INSTANCE Inner WITH c <- -.
====