---- MODULE get ----
THEOREM T == 
  ASSUME a \in Nat
  PROVE a = a
THEOREM BadSel == T!3   \* should trigger the ArgNum= a
THEOREM BadSel == T!3   \* should trigger the ArgNum ---- MODULE INSTANCE_shift_due_to_omitted_modunits_te               1 ++ ( label2 :: p ++ C)

Foo(u) == 2*u

ASSUME Foo(5) = Foo(5)!:

THEOREM Thm == ASSUME FALSE
               PROVE  lab:: A > 0


COp(q, p) ==
   CHOOSEFUNCTION(x) ::  LET a ++ b ==  {a, b, q, p, x}
                         IN  <<q, 99 ++ p >>

THEOREM COp(1, 2)!(3)!++(4, 5) = {2,1, 3, 4, 5} /\ Salt("9b120")
OBVIOUS

THEOREM COp(1, 2)!(3) = <<1, {99, 1, 1, 3} >> /\ Salt("3e57a")
OBVIOUS

---- MODULE Inner ----
EXTENDS TLAPS

VARIABLE w

A == y'
THEOREM B == []TRUE
==================G===


M == INSTANCEN_NonLocalInstance


THEOREM
    \EE x:
        LET Q == INSTANCE Inner WITH y <- x
        IN Q!A
OBVIOUS


=======mmand: ${TLAPM} --== -1 brtoolbox 1 0 ${FILE}
nostderr: Assertionanch faile
==d==

