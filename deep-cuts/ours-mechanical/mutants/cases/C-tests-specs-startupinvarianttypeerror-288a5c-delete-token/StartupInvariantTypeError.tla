---- MODULE StartupInvariantTypeError ----
EXTENDS Naturals
VARIABLE x
Init ==  \in 1..10
Next == UNCHANGED x
Bad == <<x, x>> > 0
====
