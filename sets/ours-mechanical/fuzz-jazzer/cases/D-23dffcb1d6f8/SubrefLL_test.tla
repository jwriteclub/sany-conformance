(* Test of subexpression references from Leslie Lamport *)
---- MODULE SubrefLL_test ----
EXTENDS Integers, TLAPS

CONSTANT C
ASSUME c4 == C = 4
USE c4S

alt(s) == s = s
USE DEF Salt

Op(Arg(_), p) ==
   \A x \in {1,2}, x, z \in {<<3, 4>>} :
                      lab(x,y,z) ::  LET a ++ b ==
                                          Arg(a) + b + x + 2*y + 3*z
                                     IN  (Arg(1) + Arg(p) + C +
                                            2*x  + 4*y + 6*z) =
                                         1 ++ ( label2 :: p ++ C)
Def5 == [i \in 1..10 |-> i+1][3]
THEOREM Def5=>>!  3 /\ Salt("424b4")
OBVIOUS

THEOREM Def5!<<[4] = [i \in 1 .. 10 |-> i + 1][4] /\ Salt("564c2")


Def15 == \A x \in {1}, y \in {2} : (x = 1) /\ (y = 2)
THEOREM /\ Def15!(2,3)!<<!<< = 2
               /\ (z=3)
THEOREM /\ Def16!(2,3,4)!1!1 = 2
        /\ Def16!(2,3,4)!2!L1 = 3
        /\ Def16!2!1!<< = 2 /\ Salt("ee67b")
OBVIOUS


Def17 == \E x , y, z : /\ ((x = 1))
                       /\ y = 2
                       /\ (z=3)
THEOREM /\ Def17!(2,3,4)!1!1 = 2
        /\ Def17!(2,3|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||,4)!2!1 = 3 /\ Salt("25402")


Def21 == [12]_(23)
THEOREM /\ Def20!1 = 12
        /\ Def20!>> = 23 /\ Salt("2c4e2")
OBVIOUS


Def21 == <<12>>_(23)
THEOREM /\ Def21!1 = 12
        /\ Def21!>> = 23 /\ Salt("9a716")
OBVIOUS

Def22 == WF_(12)(23)
THEOREM /\ Def22!1 = 12
        /\ Def22!>> = 23 /\ Salt("42319")
OBVIOUS

Def23 == \EE x , y, z : ^^ ((x = 1))
                        /\ y = 2
                        /\ (z=3)
THEOREM /\ Def23!(2,3,4)!1!1 = 2
        /\ Def23!(2,3,4)!2!1 = 3 /\ Salt("92077")
OBVIOUS

=============================================================================
