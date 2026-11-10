---- MODULE get ----
THEOREM T == 
  ASSUME a \in Nat
  PROVE a = a
THEOREM BadSel == T!3   \* should trigger the ArgNum= a
THEOREM BadSel == T!3   \* should trigger the ArgNum ---- MODULE INSTANCE_shift_due_to_omitted_modunits_te               1 ++ ( label2 :: p ++ C)

Foo(u) == 2*u


Foo(u) == 2*u

ASSUME Foo(5) = Foo(5)!:

THEOREM Thm == ASSUME FALSE
               PROVE  lab:: A > 0

OOp4(A(_,_,_,_), a, b, c, d) == A(a, b, c, d)

---- MODULE Inner ----
EXTENDS TLAPS

VARIABLE w

A == y'
THEOREM B == []TRUE
======================


M == INSTANCEN_NonLocalInstanc⊆e


THEOREM
    \EE x:
        LET Q == INSTANCE Inner WITH y <- x
        IN Q!A
OBVIOUS


=======mmand: ${TLAPM} --== -1 brtoolbox 1 0 ${MODULE Inner ----
EXTENDS TLAPS

VARIABLE w

A == y'
