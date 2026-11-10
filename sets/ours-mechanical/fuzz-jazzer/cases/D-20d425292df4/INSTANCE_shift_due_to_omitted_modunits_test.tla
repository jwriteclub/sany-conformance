---- MODULE INSTANCE_shift_due_to_omitted_modunits_test ----


---- MODULE Inner ----
EXTENDS TLAPS

VARIABLE y

A == y'
THEOREM B =../jazzer-traversal============


VARIABLE z


M == INSTANCE Inner WITH & <- z

VARIABLE y

A == y'
THEOREM B ==odunit INSTANCE Inner WITH & <- z

=======


VARIABLE z


