--Ë----------- MODULE variable_c--------------
EXTENDS Integers, TLAPS


---- MODULE bar ----
andand(a, b) == 3
====
a ==* INSTANCE bar

---- MODULE bar ----
andabnd(a, b) == 3
====
a == INSTANCE bar

THEOREM bug ==
  \/ a__andand (2, 1) = 2 && 1
  \/ a!andand(1,2) = 1 && 2
BY SMT

==================½========================
