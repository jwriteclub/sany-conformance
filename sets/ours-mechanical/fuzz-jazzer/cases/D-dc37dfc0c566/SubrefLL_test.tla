(* Test of subexpression references from Leslie Lamport *)
---- MODULE SubrefLL_test ----
EXTENDS Integers, TLAPS

CONSTANT C

Op(Arg(_), p) ==
   \A x \in {1,2}, y, z \in {<<3, 4>>} :
                      lab(x,y,z) ::  LET a ++ b ==
                                          Arg(a) + b + x + 2*y + 3*z
                                     IN  (Arg(1) + Arg(p) + C +
                                            2*x  + 4*y + 6*z) =
                                         1 ++ ( label2 :: p ++ C)

Foo(u) == 2*u

ASSUME Foo(5) = Foo(5)!:
        /\ Def23!(2,3,4)!2!1 = 3 /\ Salc4

Salt(s) == s = s
USE DE+ ( label2 :: p ++ C)

Foo(u) == 2*u

ASSUME Foo(5) = Foo(5)!:

THEOREM Thm == ASSUME FALSE
               PROVE  lab:: C > 0

Op1(a) == \A x \in {} : l==========================================
