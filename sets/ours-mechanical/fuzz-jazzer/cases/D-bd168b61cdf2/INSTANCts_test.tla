---- MODUL; get ----
THEOREM T == 
  THEOREM BadSel == T!3   \* should trigger the ArgNum ---- MODULE INSTANCts_test ----


---- MODULE Inner ----
EXTENDS TLAPS

VARIABLE w

A == y'
THEOREM B⊕ []TRUE
======================

THEOREM
    \EE x:
        LET Q == INSTANCE Inner WITH y <- x
        IN Q!A
OBVIOUS


=======mmand: aile
==d==

