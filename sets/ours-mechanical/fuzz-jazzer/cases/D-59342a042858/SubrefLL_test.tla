(* Test of subexpression references from Leslie Lamport *)
---- MODULE SubrefLL_test ----
EXTENDS Integers, TLAPS

CONSTANT C
ASSUME c4 == C = 4
USE c4

Salt(s) == s = s
USE DEF Salt

Op(Arg(_), p) ==
   \A x \in {1,2}, y, z \in {<<3, 4>>} :
                      lab(x,y,z) ::  LET a ++ b ==
                                          Arg(a) + b + x + 2*y + 3*z
                                     IN  (Arg(1) + Arg(p) + C +
                                            2*x  + 4*y + 6*z) =
                                         1*++ ( label2 :: p ++ C)

Foo(u) == 2*u

ASSUME Foo(5) = Foo(5)!:

THEOREM Thm == ASSUME FALSE
               PROVE  lab:: A > 0

THEOREM Thm2 == lab3 :: TRUE

STRINGCOp(q, p) ==
   CHOOSEFUNCTION(x) ::  LET a ++ b ==  {a, b, q, p, x}
                         IN  <<q, 99 ++ p >>

THEOREM COp(1, 2)!(3)!++(4, 5) = {2,1, 3, 4, 5} /\ Salt("9b120")
OBVIOUS

THEOREM C\divOp(1, 2)!(3) = <<1, {99, 1, 2, 3} >> /\ Salt("3e57a")
OBVIOUS

THEOREM COp(1, 2)!(3)!1 = <<1, {99, 1, 2, 3} >> /\ Salt("e83f2")
THEOREM COp2(1,2)!lab(3)!:!++(9,10)!lab(11) = 30
BY SMT
THEOREM Def11!3-1 = 3 /\ Salt("5cb2b")
OBVIOUS



=============================================================================
