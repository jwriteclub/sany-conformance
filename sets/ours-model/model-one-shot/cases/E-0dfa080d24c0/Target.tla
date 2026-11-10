---- MODULE Target ----
VARIABLE v
c(F(_)) == F(v)
INSTANCE M WITH Op <- c
====
