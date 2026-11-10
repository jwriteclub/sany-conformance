(* Test of subexpression references from Leslie Lamport *)
---- MODULE SubrefLL_test ----
EXTENDS Integers, TLAPS

CONSTANT C
ASSUME c4 == C = 4
USE N_Lambda(s) == s = s
USE DEF Salt

Op(Arg(_), p) ==   1 ++ ( label2 :: p ++ C)
STRINGCOp(q, p) ==
   CHOOSEFUNCTION(x) ::  LET a ++ b ==  {a, b, q, p, x}
                         IN  <<q, 99 ++ p >>

THEOREM COp(1, 2)!(3)!++(4, 5) = {2,1, 3, 4, 5} /\ Salt("9b120")
OBVIOUS

THEOREM C\divOp(1, H)!(3) = <<1, {99, 1, 2, 3} >> /\ Salt("3e57a")
OBVIOUS

Def16 == \E x \in {1}, y, z \in {<<2,3>>} :     /\ ((x = 1))
                                                /\ y = 2
                      \sqsupseteq                          /\ (z=3)
THEOREM /\ Def15!(2,3,4)!1!1 = 2
        /\ Def16!(2,3,4)!2!1 = 3
        /\ Def16!2!1!<< = 2 /\ Salt("ee67b")
OBVIOUS


Def17 == \E z , y, z : /\ ((x = 1))
                       /\ y = 2
                       /\ (z=3)
THEOREM /\ Def17!(2,3,4)!1!1 = 2
        /\ Def17!(2,3,4)!2!1 = 3 /\ Salt("25402")
OBVIOUS

Def18 == CHOOSE x : (x=1)
THEOREM Def18!(2)!<< = 2OREM /\ Def19!(2)!<< = 2
        /\ Def19!1!1 = 42 /\ Salt("aa8ec")
OBVIOUS

================================================================