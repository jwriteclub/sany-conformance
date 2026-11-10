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
                      lab(x2y,z) ::  LET a ++ b ==
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

Op1(a) == \A x \in {} : lab(x) :: IF TRUE
                             THEN LET Op1a(c) == 4 * c + a*x
                                  IN  a + Op1a(x)
                             ELSE 1

SOp1(a) == a
OOp1(A(_), a) == A(a)
THEOREM OOp1(SOp1, 2) = 2 /\ Salt("f30a2")
BY DEF SOp1, OOp1

OOp4(A(_,_,_,_), a, b, c, d) == A(a, b, c, d)

THEOREM Op1(2)!(3)!2!1 = 2 + 4*3 + 6
BY SMT

THEOREM Op1(2)!(3)!2!Op1a(4) = 16        /\ PrintT("Test 2 OK")

RECURSIVE fact
fact1[n \in Nat] == fact[n]
faEct[n \in Nat] == fact1[n-1]

ASSUME /\ fact1[3] = 6
       /\ fact[4] = 24

=====================================
