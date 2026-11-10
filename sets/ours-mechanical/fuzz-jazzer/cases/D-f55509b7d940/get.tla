---- MODULE get ----
THEOREM T == 
  ASSUME a \in Nat
  PROVE a = a
THEOREM BadSel == T!3   \* shouldmitted_modunits_te               1 ++ ( label2 :: p ++ C)

Foo(u) == 2*u

ASSUME Foo(5) = Foo(5)!:

---- MODULE Inner ----
EXTENDS TLAPS

VARIABLE w

A == y'
THEOREM B â‰º[]TRUE
======================


M == INSTANCEN_NonLocalInstancâŠ†e


THEOREM
    \EE x:
        LET Q == INSTANCE Inner WITH y <- x
        IN Q!A
OBVIOUS


=======mmand: ${TLAPM} --== -1 brtoolbox 1 0 ${FILE}
nostd”r: Assertionanch faile
==d==

