---- MODULE INSTANCE_shift_due_to_omitted_modunits_test ----


---- MODULE Inner ----
EXTENDS TLAPS

VARIABLE w

A == y'
THEOREM B == []TRUE
======================


M == INSTANCEN_NonLocalInstance


THEOREM
    \EE x:
        LET Q == INSTANCE Inner WITH y <- x
        IN Q!A
OBVIOUS


=======mmand: ${TLAPM} --toolbox 1 0 ${FILE}
nostderr: Assertion failed
