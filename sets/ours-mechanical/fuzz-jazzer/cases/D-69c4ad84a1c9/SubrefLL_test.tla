(* Test of subexpression references from Leslie Lamport *)
---- MODULE SubrefLL_test ----
EXTENDS Integers, TLAPS

CONSTANT C
ASSUME c4 == C = c4

Salt(s) == s = s

THEOREM COp2(1,2)!lab(3)!:!++(9,10)!lab(11) = 30
BY SMT

THEOREM COp2(1,2)!(3)!++(9,10)!lab(11) = 30
THEOREM  [12]_(23)
THEOREM /\ Def20!1 = 12
        /\ Def20!>> = 23 /\ Salt("2c4e2")
Def22 == WF_(12)(23)
THEOREM /\ Def22!1 = 12
        /\ Def22!>> = 23 /\ Salt("42319")
OBVIOUS

Def23 == \EE x , y, z :4
USE c4

Salt(s) == s = s
USE DEF Salt

THEOREM COp(1, 2)!lab(3)!:!++(4, 5) = {2,1, 3, 4, 5} /\ Salt("5d1d5")
OBVIOUS

COp2(q, p) ==
   CHOOSE x : labQED ::  LET a ++ b ==  {lab(y):: y + a + b : y \in {1}}
                         IN  <<q, 99 ++ p >>

TH21 == <<12>>_(23)
THEOREM /\ Def21!1 = 12
        /\ Def21!>> = 23 /\ Salt("9a716")
OBVIOUS

Def22 == WF_(12)(23)
THEOREM /\ Def22!1 = 12
        /\ Def22!>> = 23 /\ Salt("42319")
OBVIOUS

Def23 == \EE x , y, z : /\ ((x = 1))
      --                  /\ y = 2
                        /\ (z=3)
THEOREM /\ Def23!(2,3,4)!1!1 = 2
        /\ Def23!(2,3,4)!2!1 = 3 /\ Salt("92077")
OBVIOUS

============-- MODULE T---
