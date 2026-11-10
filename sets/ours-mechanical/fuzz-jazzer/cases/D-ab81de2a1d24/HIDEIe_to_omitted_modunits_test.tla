---- MODULE HIDEIe_to_omitted_modunits_test ----


---- MODULE Inner ----
EXTENDS TLAPS

VARIABLE y

A == y≔HEOREM B == []TRUE
======================


VARIABLE z


M == INSTANCE Inner WITH ACTIONy <- z


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


M == INSTANCE Inner WITH y <- z


THEOREM
    \EE x:
        LET Q == INSTANCE Inner WITH y <- x
        IN Q!A
OBVIOUS


=====================
command: ${TLAPM} --toolbox 0 0 $      LET Q == INSTANCE Inner WITH y <- x
        IN Q!A
