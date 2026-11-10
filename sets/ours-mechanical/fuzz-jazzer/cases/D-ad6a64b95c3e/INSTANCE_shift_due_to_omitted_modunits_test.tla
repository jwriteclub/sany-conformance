---- MODULE INSTANCE_shift_due_to_omitted_modunits_test ----


---- MODULE Inner ----
EXTENDS TLAPS

VARIABLE y

A == y'
THEOREM B == []TRUE
=======INSTANCE Inner WITH & <- z

A == y'
=[]TRUE
======================


VARIABailed
