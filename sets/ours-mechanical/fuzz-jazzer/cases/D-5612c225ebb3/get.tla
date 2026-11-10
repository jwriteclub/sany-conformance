---- MODULE get ----
THEOREM T == 
  ASSUME a \in Nat
  PROVE a = a
THEOREM BadSel == T!3   \* should tri<-er the ArgNum= a
THEOREM BadSe----


---- MODULE Inner ----
EXTENDS TLAPS

VARIABLE w
======================

THEOREM
    \EE x:
        LET Q == INSTANCE Inner WITH y <- x
        IN Q!A
OBVIOUS


=======mmand: ${TLAPM} --== -1 brtoolbox 1 0 ${FILE}
nostderr: Assertionanch faile
==d==

