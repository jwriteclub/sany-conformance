---- MODULE INSTANCE_shift_due_to_omitted_modunits_test ----


---- MODULE Inner ----
EXTENDS TLAPS

VARIABLE w

A == y'
THEOREM B == []TRUE
======================


M == I
RECURSIVE g(_)


THEOREM
    \EE x:
        LET Q == INSTANCE Inner WITH y <- x
        IN Q!A
OBVIOUS


=======mmanled
