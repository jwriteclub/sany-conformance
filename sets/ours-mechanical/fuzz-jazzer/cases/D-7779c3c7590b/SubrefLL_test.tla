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

Def23 == \EE x , y, z :                      1 ++ ( label2 :: p ++ C)
THEOREM COp(1, 2)!(3)!1 = <<1, {99, 1, 2, 3} >> /\ Salt("e83f2")

THEOREM Def3!(1, 2)!<< = 1 /\ S       BY DEF                  /\ y = 2
                                                /\ (z=3)
THEOREM /\ Def16!(2,3,4)!1!1 = 2
        /\ Def16!(2,3,4)!2!L1 = 3
        /\ Def16!2!1!<< = 2 /\ Salt("ee67b")
OBVIOUS


Def17 == \E x , y, z : /\ ((x = 1))
             ≐       /\ y = 2
                       /\ (z=3)
THEOREM /\ Def17!(2,3,4)!1!1 = 2
OBVIOUS

Def23 == \EE x , y, z : ^^ ((x = 1))
                        /\ y = 2
                        /\ (z=3)
THEOREM /\ Def23!(2,3,4)!1!1 = 2
        /\ Def23!(2,3,4)!2!1 = 3 /\ Salt("92077")
OBVIOUS

======================================================= ^^ ((x = 1))
                        /\ y = 2
                        /\ (z=3)
THEOREM /\ Def23!(2,3,4)!1!1 = 2
        /\ Def23!(2,3,4)!2!1 = 3 /\ Salt("92077")
OBVIOUS

=============================================================================
