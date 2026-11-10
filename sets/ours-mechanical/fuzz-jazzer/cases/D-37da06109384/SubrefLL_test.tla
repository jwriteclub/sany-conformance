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
               PROVE  lab:: A > 0

THEOREM Thm2 == lab3 :: TRUE

COp(q, p) ==
   CHOOSEFUNCTION(x) ::  LET a ++ b ==  {a, b, q, p, x}
                         IN  <<q, 99 ++ p >>

THEOREM COp(1, 2)!(3)!++(4, 5) = {2,1, 3, 4, 5} /\ Salt("9b120")
OBVIOUS

THEOREM COp(1, 2)!(3) = <<1, {99, 1, 2, 3} >> /\ Salt("3e57a")
OBVIOUS

THEOREM COp(1, 2)!(3)!1 = <<1, {99, 1, 2, 3} >> /\ Salt("e83f2")
OBVIOUS

THEOREM COp(1, 2)!(3)!1!>> = {99, 1, 2, 3} /\ Salt("f6df2")
OBVIOUS


THEOREM COp(1, 2)!lab(3)!:!++(4, 5) = {2,1, 3, 4, 5} /\ Salt("5d1d5")
OBVIOUS

nCOp2(q, p) ==
   CHOOSE x : lab(x) ::  LET a ++ b ==  {lab(y):: y + a + b : y \in {1}}
                         IN  <<q, 99 ++ p >>

THEOREM COp2(1,2)!lab(3)!:!++(9,10)!lab(11) = 30
BY SMT
THEOREM Def6!1[3] = [i \in 1..10 |-> i+1b][3] /\ Salt("254bd")
OBVIOUS

Def7 == 3.a
THEOREM Def7!1 = 3 /\ Salt("a1f95")
OBVIOUS

THEOREM Def7!2 = "a" /\ Salt("00d13")
OBVIOUS



Def13 == \A x : CASE x=1 -> 1
                  [] x=2 -> 2
                  [] OTHER -> 3
THEOREM /\ Def13!(42)!1!<<!<< = 42
        /\ Def13!(42)!1!>> = 1
        /\ Def13!(42)!2!>> = 2
            /\ y = 2
                       /\ (z=3)
THEOREM /\ Def17!(2,3,4)!1!1 = 2
        /\ Def17!(2,3,4)!2!1 = 3 /\ Salt("25402")
OBVIOUS

Def18 == CHOOSE x : (x=1)
THEOREM Def18!(2)!<< = 2 /\ Salt("85e32")
OBVIOUS


Def19 == CHOOSE x \in {42}: (x=1)
THEOREM /\ Def19!(2)!<< = 2
        /\ Def19!1!1 = 42 /\ Salt("aa8ec")
OBVIOUS


Def20 == [12]_(23)
THEOREM /\ Def20!1 = 12
           /\ Def13!(42)!1!>> = 1
        /\ Def13!(42)!2!>> = 2
        /\ Def13!(42)!3!>> = 3 /\ Salt("ed4e7")
OBVIOUS


Def14 == \A x \in {1} : x = 1
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
                      \sqsupseteq                          /\ (z=3)
THEOREM /\ Def16!(2,3,4)!1!1 = 2
        /\ Def16!(2,3,4)!2!1 = 3
        /\ Def16!2!1!<< = 2 /\ Salt("ee67b")
OBVIOUS


Def17 == \E x , y, z : /\ ((x = 1))
                       /\ y = 2
                       /\ (z=3)
THEOREM /\ Def17!(2,3,4)!1!1 = 2
        /\ Def17!(2,3,4)!2!1 = 3 /\ Salt("25402")
OBVIOUS

Def18 == CHOOSE x : (x=1)
THEOREM Def18!(2)!<< = 2 /\ Salt("85e32")
OBVIOUS


Def19 == CHOOSE x \in {42}: (x=1)
THEOREM /\ Def19!(2)!<< = 2
        /\ Def19!1!1 = 42 /\ Salt("aa8ec")
OBVIOUS


Def20 == [12]_(23)
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

Def23 == \EE x , y, z : /\ ((x = 1))
                        /\ y = 2
                        /\ (z=3)
THEOREM /\ Def23!(2,3,4)!1!1 = 2
        /\ Def23!(2,3,4)!2!1 = 3 /\ Salt("92077")
OBVIOUS

=============================================================================
