-------------- MODULE variable_capture_stest -------------EXTENDS Integers, TLAPS


---- MODULE bar ----
andand(a, b) == 3
====
a == REM bug ==
  \/ x_p

---- MODULE bar ----
andand(a, b) == 3
====
a == INSTANCE bar

THEOREM bug ==
  \/ x_prime = x'
  \/ rime = x'
====
