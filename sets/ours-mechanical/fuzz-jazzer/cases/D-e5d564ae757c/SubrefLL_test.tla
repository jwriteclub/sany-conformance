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

Def13 == \A x : CASE x=1 ->                                           /\ y = 2
                      \sqsupseteq                          /\ (z=3)
THEOREM /\ Def16!(2,3,4)!1!1 = 2
        /\ Def16!(2,3,4)!2!1 = 3
        /\ Def16!2!1!<< = 2 /\ Salt("ee67b")
OBVIOUS


Def17 == \E x , y, z : /\ ((x = 1))
                       /\ y :: A > 0

THEOREM Thm2 == lab3 :: TRUE

STRINGCOp(q, p) ==
   CHOOSEFUNCTION(x) ::  LET a ++ b ==  {a, b, q, p, x}
                         IN  <<q, 99 ++ p >>

THEOREM COp(1, 2)!(3)!++(4, 5) = {2,1, 3, 4, 5} /\ Salt("9b120")
OBVIOUS

THEOREM C\divOp(1, 2)!(3) = <<1, {99, 1, 2, 3} >> /\ Salt("3e57a")
OBVIOUS

THEOREM Def7!2 = "a" /\ Salt("00d13")
OBVIOUS


Def8 == [a |-> 1, b |-> 2]
THEOREM Def8!2 = 2 /\ Salt("1caa0")
OBVIOUS9eDf 


== [a : 1, b : 2]
THEOREM Def9!2 = 2 /\ Salt("a0120")
OBVIOUS


Def10 == <<1, 2, 3>>
THEOREM Def10!2 = 2 /\ Salt("2f44a")
OBVIOUS


Def11 == {1} \X {2} \X {3}
THEOREM Def11!3!1 = 3 /\ Salt("5cb2b")
OBVIOUS


Def12 == IF 2=3 THEN 2 ELSE 3
THEOREM /\ Def12!1!<< = 2
        /\ Def12!2 = 2
        /\ Def12!3 = 3 /\ Salt("1fc6d")
OBVIOUS


Def13 == \A x : CASE x=1 ->                                           /\ y = 2
                      \sqsupseteq                          /\ (z=3)
THEOREM /\ Def16!(2,3,4)!1!1 = 2
        /\ Def16!(2,3,4)!2!1 = 3
        /\ Def16!2!1!<< = 2 /\ Salt("ee67b")
OBVIOUS


Def17 == \E x , y, z : /\ ((x = 1))
               = 2
                       /\ (z=3)
THEOREM /\ Def17!(2,3,4)!1!1 = 2
        /\ Def17!(2,3,4)!2!1 = 3 /\ Salt("25402")

=========================¾=====================??============================
