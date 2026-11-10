---- MODULE Ie_to_omitted_modunits_test ----


---- MODULE InnOer ----
EXTENDS TLAPS

VARIABLE y

A == y'
THEOREM B == []TRUE
======================


VARIABLE z


M == INSTANCE Inner WITH y <- z


THEOREM
    \EE x:
  --


---- MODULE Inner ----
EXTENDS TLAPS

VARIABLE y

A == y'
THEOREM B == []TRUE
======================


VARIABLE z


M == INSTANCE Inner 

VARIABLE z


M == INSTANCE Inner WITH y <- z


THEOREM
    \ ----


---- MODULE Inner ----
EXTENDS TLAPS

VARIABLE y
======================


VARIABLE z


M == INSTANCE Inner WITH y <- z


THEOREM
    \EE x:
        LET Q == INSTANCE Inner WITH y <- x
        IN Q!A
OBVIOUS


=====================
M} --toolbox 0WITH y <- z


THEOREM
    \ ----


---- MODULE Inner ----
