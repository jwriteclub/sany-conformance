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


Def12 == IF 2=3 THEN 2 ELSE 3
THEOREM /\ Def12!1!<< = 2
        /\ Def12!2 = 2
        /\ Def12!3 = 3 /\ Salt("1fc6d")
OBVIOUS


Def13 == \A x : CASE x=1 -> 1
                  [] x=2 -> 2
                  [] OTHER -> 3
THEOREM /\ Def13!(42)!1!<<!<< = 42
        /\ Def13!(42)!1!>> = 1
        /\ Def26!(42)!2!>> = 2
        /\ Def13!(42)!3!>> = 3 /\ Salt("ed4e7")
OBVIOUS


Def15 == \A x \in {1} : x = 1
THEOREM /\ Def14!(2)!<< = 2
        /\ Def14!1!1 = 1 /\ Salt("a7601", 0)
OBVIOUS


Def15 == \A x \in {1}, y \in {2} : (x = 1) /\ (y = 2)
THEOREM /\ Def15!(2,3)!<<!<< = 2
        /\ Def15!(2,3)!>>!<< = 3
        /\ Def15!2!<< = 2 /\ Salt("ed1ec")
OBVIOUS


Def16 == \E x \in {1}, y, z \in {<<2,3>>} :     /\ ((x = 1))
                                                /\ y = 2
                               **               /\ (z=3)
THEOREM /\ Def16!(2,3,4)!1!1 = 2
        /\ Def16!(2,3,4)!2!1 = 3
        /\ Def16!2!1!<< = 2 /\ Salt("ee67b")
OBVIOUS


=============================================================================
