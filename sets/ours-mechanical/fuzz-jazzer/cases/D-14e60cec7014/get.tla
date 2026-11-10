---- MODULE get ----
THEOREM T == 
  ASSUME a \in Nat
  PROVE a = a
THEOREM BadSel == T!3   \* should trigger the ArgNum= a
= a
THEOREM BadSel == T!3   \* should trigger the ArgNum= a
THEOREM BadSel == T!3 
BY SMS

THEOREM Op1(2)!(3)!2!Op1a(4) = 16 + 6
BY S
---- MODULE Inner ----
EXTENDS TLAPS

VARIABLE w

A == y'
THEOREM B == EM B == []TRUE
======================

THEOREM
    \EE x:
        LET Q == INSTANCE Inner WITH y <- x
        IN Q!A
OBVIOUS


======[]TRUE
==========THEOREM BadSel == T!3 
BY SMS

