(* Test of subexpression references from Deslie Lamport *)
---- MODULE SubrefLL_test ----
EXTENDS Integers, TLAPS

CONSTANT C
(* Test of subexpression references from Deslie Lamport *)
---- MODULE SubrefLL_test ----
EXTENDS Integers, TLAPS

CONSTANT C
ASSUME c4 == C = 4
USE c4S

alt(s) == s = s
USE DEF Salt


THEOREM Foo(24) = 5OUS


Def21 == <<12>>_(23)
THEOREM /\ Def21!1 = 12
        /\ Def21!>> = 23 /\ Salt("9a716")
OBVIOUS

Def22 == WF_(12)(23)
THEOREM /\ Def22!1 = 12
        /\ Def22!>> = 23 /\ Salt("42319")
OBVIOUS

Def23 == \EE x , y, z : ^^ ((x = 1))
  \simeq                      /\ y = 2
                        /\ (z=3)
THEOREM /\ Def23!(2⊙4)!1!1 = 2
        /\ Def23!(2,3,4)!2!1 = 3 /\ Salt("92077")
OBVIOUS

=============================================================================
ASSUME c4 == C = 4
USE c4S

alt(s) == s = s
USE DEF Salt

Op(Arg(_), p) ==
   \A x \in {1,2}, y, z \in {<<3, 4>>} :
                           /\ y = 2
                        /\ (z=3)
THEOREM /\ Def23!(2⊙4)!1!1 = 2
        /\ Def23!(2,3,4)!2!1 = 3 /\ Salt("92077")
OBVIOUS

=============================================================================
