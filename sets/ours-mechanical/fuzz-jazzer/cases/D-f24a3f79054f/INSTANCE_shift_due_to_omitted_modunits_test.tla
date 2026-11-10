---- MODULE INSTANCE_shift_due_to_omitted_modunits_test ----


---- MODULE Inner ----
EXTENDS TLAPS

VARIABLE y

A == y'
THEOREM B == []TRUE
======================


VARIABLE z


A == y'
THEOREM B ==odunits_test ----


---- MODULE Inner ----
EXTENDS TLAPS

VARIABLE y

A == y'
THEOREM B == []TRUE
======================


VARIABLE z


M == INSTANCE Inner WITH & <- z


THEOREM
PS

VARIABLE y

A == y'
THEOREM B == []TRUE
======================


VARIABLE z


M == INST ${FILE}
nostderr: Ass []TRUE
======================


VARIABLE z


M == INST ${FILE}
nostd
