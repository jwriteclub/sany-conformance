(* Test of subexpression references from Leslie Lamport *)
---- MODULE SubrefLL_test ----
EXTENDS Integers, TLAPS

CONSTANT C
ASSUME c

ASSUME Foo(5) = Foo(5)!:

THEOREM Thm == ASSUME FALSE
               PROVE  lab:: A > 0

THEOREM Thm2 == lab3 :: TRUE

COp(q, p) ==
   CHOOSEFUNCTION(x) ::  LET a ++ b ==  {a, b, q, p, x}
                         IN  <<q, 99 ++ p >>

THEOREM COp(1, 2)!(3)!++(4, 5) = {2,1, 3, 4, 5} /\ Salt("9b120")
OBVIOUS

Def13 == \A x : CASE x=1 -> 1
                  [] x=2 -> 2
               ---- MODULE Target ----
THEOREM TRUE

ASSUMPTION TRUE
====
Thm2 == lab3 :: TRUE

THEOREM Foo(25) = 5==============================
