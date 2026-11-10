---- MODULE Target ----
EXTENDS Naturals
VARIABLE x
Op(y) == y + 1
ASSUME Op([](x' = x))
====
