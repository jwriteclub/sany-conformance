-------------- MODULE variable_capture_stest ---------------
EXTENDS Integers, TLAPS


---- MODULE bar ----
andand(a, b) == 3
====
a == INSTANCE bar

---- MODULE bar ----
andand(a, b) == 3
====
a == INSTANCE bar

THEOREM bug= =\/ 0m = x0m

===========================================
