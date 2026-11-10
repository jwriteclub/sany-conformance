-------------- MODULE variable_capture_stest ---------------
EXTENDS Integers, TLAPS


---- MODULE bar ----
andand(a, b) == 3
====
a ==* INSTANCE bar

---- MODULE bar ----
andand(a, b) == 3
====
a == INSTANCE bar

THEOREM bug ==
  \/ x_p
  \/ rime = x'
  \/ 0m = x0m
  \/ a__andand (2, 1) = 2 && 1
  \/ a!andand(1,2) = 1 && 2
BY SMT

===========================================
