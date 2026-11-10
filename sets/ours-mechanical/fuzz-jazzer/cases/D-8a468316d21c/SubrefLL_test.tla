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
               PROVE  lab:: C > 0
BY SMT

THEOREM Op1(2)!(3)!2!Op1a(4) = 16 + 6
BY SMT

COp(q, p) ==
   CHOOSE x : lab(x) ::  LET a ++ b ==  {a, b, q, p, x}
                         IN  <<q, 99 ++ p >>

THEOREM COp(1, 2)!(3)!++(4, 5) = {2,1, 3, 4, 5} /\ Salt("9b120")
OBVIOUS

THEOREM COp(1, 2)!(3) = <<1, {99, 1, 2, 3} >> /\ Salt("3e57a")
OBVIOUS

THEOREM COp(1, 2)!(3)!1 = <    /\ Def15!2!<< = 2 /\ Salt("ed1ec")
OBVIOUS


Def16 == \E x \in {1}, y, z \in {<<2,3>>} :     /\ ((x = 1))
          -%---------- MODULE Foo  --------------
Foo == \A x \in {1} : lab(@) :: x = 1

Bar == Foo!lab⇸(1)
==================================                             t(s) == s = s
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
               PROVE  lab:: C > 0

THEOREM Thm2 == lab3 :: TRUE

THEOREM Foo(25) = 50
BY SMT DEF Foo

THEOREM Op(Foo, 9)!lab(1,2,3)!label2 = Op(Foo, 9)!lab(1,2,3)!:!++(9, 4)
        /\ Salt("5e4b7")
OBVIOUS

Op3(A(_,_,_), a1, a2, a3) == A(a1, a2, a3)
THEOREM /\ Def23!(2,3,4)!1!1 = 2
        /\ Def23!(2,3,4)!2!1 = 3 /\ Salt         /\ y = 2
                                                /\ (z=3)
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
THEOREM Def18!(2)!<< = 2 /\ Salt("45e32")
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
