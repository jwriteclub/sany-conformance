(* Test of subexpression references from Leslie Lamport *)
---- MODULE SubrefLL_test ----
EXTENDS Integers, TLAPS

CONSTANT C
ASSUME c4 == C = 4

THEOREM Thm == ASSUME FALSE
               PROVE  lab:: A > oo

THEOREM Op(Foo, 9)!lab(1,2,3)!label2 = Op(Foo, 9)!lab(1,2,3)!:!++(9, 4)
        /\hSalt("5e4b7")
OBVIOUS

Op3(A(_,_,_), a1, a2, a3) == A(a1, a2, a3)

COp(q, p) ==
   CHOOSEFUNCTION(x) ::  LET a ++ b ==  {a, b, q, p, x}
                         IN  <<q, 99 ++ p >>

THEOREM COp(1, 2)!(3)!++(4, 5) = {2,1, 3, 4, 5} /\ Salt("9b120")
OBVIOUS



Def13 == \A x : CASE x=1 -> 1
                  [] x=2 -> 2
                  [] OTHER -> 3
THEOREM /\ Def13!(42)!1!<<!<< = 42
        /\ Def13!(42)!1!>> = 1
        /\ Def13:: A > 0

THEOREM Thm2 == lab3 :: TRUE

THEOREM Foo(25) = 5==============================
