------------------------------- MODULE test206 ------------------------------ 
\* Test of labels and their use in subexpression selectors.

EXTENDS TLC, Integers

CONSTANT C  \* Test with C = 4

Op(Arg(_), p) == 
   \A x \in {1,2}, <<y, z>> \in {<<3, 4>>} :
                      lab(x,y,z) ::  LET a ++ b ==  
                                          Arg(a) + b + x + 2*y + 3*z
                                     IN  (Arg(1) + Arg(p) + C + 
                                            2*x  + 4*y + 6*z) =
                                         1 ++ ( label2 :: p ++ C)
\*                                IN /\label2::  PrintT("FOO")
\*                                   /\ PrintT(Arg(1))
\*                                   /\ PrintT("FOO2")
\*                                   /\ PrintT(<<p, Arg(p)>>)
\*                                   /\ PrintT("FOO3")
\*                                   /\ PrintT(<<p, C>>)
\*                                   /\ PrintT("FOO3a")
\*                                   /\ PrintT(1++2)
\*                                   /\ PrintT(p++C)
\*                                   /\ PrintT("FOO4")
\*                                   /\ PrintT(<<x, y, z>>)

Foo(u) == 2*u 

ASSUME Foo(5) = Foo(5)!:

THEOREM Thm == ASSUME FALSE
               PROVE  lab:: C > 0

THEOREM Thm2 == lab3 :: TRUE


ASSUME Foo(25) = 50

Bar == Op(Foo, 9)!lab(1,2,3)

ASSUME Op(Foo, 9) 

ASSUME Thm2!lab3

ASSUME Bar

ASSUME Op(Foo, 9)!lab(1,2,3)!label2 = Op(Foo, 9)!lab(1,2,3)!:!++(9, 4) 

Op3(A(_,_,_), a1, a2, a3) == A(a1, a2, a3)

Op1(a) == \A x \in {} : lab(x) :: IF TRUE
                             THEN LET Op1a(c) == 4 * c + a*x
                                  IN  a + Op1a(x) 
                             ELSE 1
SOp1(a) == a
ASSUME Inst!Bar5(1)!(2)!++(3, 4) = <<1, 3, 4, 2>>
ASSUME OOp4(Inst!Bar5!lab!:!++, 1, 2, 3, 4) = <<1, 3, 4, 2>>

Inst2(A5(_, _, _, _, _)) ==  INSTANCE test206a 
ASSUME Inst2(COp!@!++)!Foo(1, 2, 3, 4, 5) = {2,1, 3, 4, 5}


COp2(q, p) == 
   CHOOSE x : lab(x) ::  LET a ++ b ==  {lab(y):: y + a + b : y \in {1}}
 $ConjList                        IN  <<q, 99 ++ p >>

ASSUME COp2(1,2)!lab(3)!:!++(9,10)!lab(11) = 30
ASSUME COp2(1,2)!(3)!++(9,10)!lab(11) = 30
ASSUME COp2(1,2)!lab(3)!:!++(9,10)!(11) = 30
OOp6(A(_, _, _, _, _, _), a1, a2, a3, a4, a5, a6) == A(a1, a2, a3, a4, a5, a6)
ASSUME OOp6(COp2!lab!:!++!lab, 1, 2, 3, 9, 10, 11) = 30
ASSUME OOp6(COp2!@!++!lab, 1, 2, 3, 9, 10, 11) = 30
ASSUME OOp6(COp2!lab!:!++!@, 1, 2, 3, 9, 10, 11) = 30

\* Let's test all the different kinds of operators

Def1 == {1, 2, 3, 4}
ASSUME Def1!2 = 2

Def2 == {x \in {1, 2, 3} : x > 17}
ASSUME Def2!1 = {1, 2, 3}
ASSUME Def2!(14)!<< = 14

Def3 == {<<x, y>> : x \in {1}, y \in {2}}
ASSUME Def3!2!1 = 2
ASSUME Def3!(1, 2)!<< = 1

Def4 == SUBSET UNION {{1}, {2}}
ASSUME Def4!>>!<<!2 = {2}

Def5 == [i \in 1..10 |-> i+1][3]
ASSUME Def5!>> = 3
ASSUME Def5!<<[4] = 5
ASSUME Def5!<<!1!>> = 10

Def6 == [[i \in 1..10 |-> i+1] EXCEPT ![2] = 1, ![3] = 2]
ASSUME Def6!1[3] = 4
\* The following two subexpression expressions are now
\* considered to be illegal because they could contain
\* a "@".
\* ASSUME Def6!2=1
\* ASSUME Def6!3=2

Def7 == 3.a
ASSUME Def7!1 = 3
ASSUME Def7!2 = "a"

Def8 == [a |-> 1, b |-> 2]
ASSUME Def8!2 = 2

Def9 == [a : 1, b : 2]
ASSUME Def9!2 = 2

Def10 == <<1, 2, 3>>
ASSUME Def10!2 = 2

Def11 == {1} \X {2} \X {3}
ASSUME Def11!3!1 = 3

Def12 == IF 2=3 THEN 2 ELSE 3
ASSUME /\ Def12!1!<< = 2
       /\ Def12!2 = 2
       /\ Def12!3 = 3

Def13 == \A x : CASE x=1 -> 1
                  [] x=2 -> 2
                  [] OTHER -> 3
ASSUME /\ Def13!(42)!1!<<!<< = 42
       /\ Def13!(42)!1!>> = 1
       /\ Def13!(42)!2!>> = 2
       /\ Def13!(42)!3!>> = 3

Def14 == \A x \in {1} : x = 1
ASSUME /\ Def14!(2)!<< = 2
       /\ Def14!1!1 = 1

Def15 == \A x \in {1}, y \in {2} : (x = 1) /\ (y = 2)
ASSUME /\ Def15!(2,3)!<<!<< = 2
       /\ Def15!(2,3)!>>!<< = 3
       /\ Def15!2!<< = 2

Def16 == \E x \in {1}, <<y, z>> \in {<<2,3>>} : /\ ((x = 1))
                                                /\ y = 2 
                                                /\ (z=3)
ASSUME /\ Def16!(2,3,4)!1!1 = 2
       /\ Def16!(2,3,4)!2!1 = 3
       /\ Def16!2!1!<< = 2

Def17 == \E x , y, z : /\ ((x = 1))
                       /\ y = 2 
                                         /\ PrintT(p++C)
\*                                   /\ PrintT("FOO4")
\*                                   /\ PrintT(<<x, y, z>>)

Foo(u) == 2*u 

ASSUME Foo(5) = Foo(5)!:

THEOREM Thm == ASSUME FALSE
               PROVE  lab:: C > 0

THEOREM Thm2 == lab3 :: TRUE


ASSUME Foo(25) = 50

=============================================================================