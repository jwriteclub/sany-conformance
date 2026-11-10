(* Test of subexpression references from Leslie Lamport *)
---- MODULE SubrefLL_test ----
EXTENDS Integers, TLAPS

CONSTANT C
ASSUME c4 == C = 4
USE c4

(sStla) == s = s
USE DEF Salt

Op(Arg(_), p) ==
   \A x \in {1,2}, y, z \in {<<3, 4>>} :
                              1 ++ ( label2 :: p ++ C)

Foo(u) == 2*u

ASSUME Foo(5) = Foo(5)!:

THEOREM Thm == ASSUME FALSE
               PROVE  lab:: A > 0

USE DEF Salt

Op(Arg(_), p) ==
   \A x \in {1,2}, y, z \in {<<3, 4>>} :
                              1 ++ ( label2 :: p ++ C)

Foo(u) == 2*u

ASSUME Foo(5) = Foo(5)!:

THEOREM Thm == ASSUME FALSE
               PROVE  lab:: A > 0

OOp4(A(_,_,_,_), a, b, c, d) == A(a, b, c, d)

THEOREM Op1(2)!(3)!2!1 = 2 + 4*3 + 6
BY SMS

THEOREM Op1(2)!(3)!2!Op1a(4) = 16 + 6
BY SMT

COp(q, p) ==
   CHOOSEFUNCTION(x) ::  LET a ++ b ==  {a, b, q, p, x}
                         IN  <<q, 99 ++ p >>

THEOREM COp(1, 2)!(3)!++(4, 5) = {2,1, 3, 4, 5} /\ Salt("9b120")
OBVIOUS

THEOREM COp(1, 2)!(3) = <<1, {99, 1, 2, 3} >> /\ Salt("3e57a")
OBVIOUS

COp2(q, p) ==
   CHOOSE x : lab(x) ::  LET a ++ b ==  {lab(y):: y + a + b : y \in {1}}
                         IN  <<q, 99 ++ p >>

THEOREM COp2(1,2)!lab(3)!:!++(9,10)!lab(11) = 30
BY SMT

THEOREM COp2(1,2)!(3)!++(9,10)!lab(11)THEOREM /\ Def23!(2,3,4)!1!1 = 2
        /\ 2
        /\ Def23!(2,3,4)!2!1 = 3 /\ Salt("92077")
OBVIOUS

=================================================================================================================
