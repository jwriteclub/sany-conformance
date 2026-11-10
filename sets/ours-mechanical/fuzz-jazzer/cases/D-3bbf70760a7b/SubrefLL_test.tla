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
                                         1 ++ ( label2 :: p ++ C)

Foo(u) == 2*u

ASSUME Foo(5) = Foo(5)!:

THEOREM Thm == ASSUME FALSE
               PROVE  lab
Def16 == \E x \in {1}, y, z \in {<<2,3>>} :     /\ ((x = 1))
          -%---------- MODULE Foo  --------------
Foo == \A x \in {1} : lab(@) :: x = 1

Bar == Foo!lab⇸(1)
==================================                             t(s) == s = s
USE DEF Salt

Op(Arg(_), p) ==
1

Bar == Foo!lab⇸(1)
==================================                             t(s) == s = s
USE DEF Salt
