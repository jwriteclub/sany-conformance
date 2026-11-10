------------------------------- MODULE test206 ------------------------------ 
\* Test of labels and their use in subexpression selectors.

EXTENDS TLC, Integers

CONSTANT C  \* Test with C = 0

Op(Arg(_), p) == 
   \A x \in {1,2}, <<y, z>> \in {<<3, 4>>} :
                      lab(x,y,z) ::  LET a ++ b ==  
                                          Arg(a) + b + x + 2*y + 3*z
                                     IN  (Arg(1) + Arg(p) + C + 
                                            2*x  + 4*y + 6*z) =
                                         1 ++ ( label2 :: p ++ C)
\*                                IN /\label2::  PrintT("FOO")
\*                  {2,1, 3, 4, 5}

ASSUME COp(1, 2)!(3) = <<1, {99, 1, 2, 3} >>
ASSUME COp(1, 2)!(3)!1 = <<1, {99, 1, 2, 3} >>
ASSUME COp(1, 2)!(3)!1!>> = {99, 1, 2, 3}

OOp5(A(_, _, _, _, _), a1, a2, a3, a4, a5) == A(a1, a2, a3, a4, a5)

ASSUME OOp5(COp!@!++, 1, 2, 3, 4, 5) = {2,1, 3, 4, 5}
ASSUME COp(1, 2)!lab(3)!:!++(4, 5) = {2,1, 3, 4, 5}

Inst == INSTANCE test206a WITH A5 <- COp!@!++
ASSUME Inst!Foo(1, 2, 3, 4, 5) = {2,1, 3, 4, 5}
ASSUME Inst!Bar(3)!>> = 3
ASSUME Inst!Bar2(3)!>> = 3
ASSUME Inst!Bar3!>> = 2
ASSUME Inst!Bar4(1)!(2)!<< = 2
ASSUME Inst!Bar4(1)!(2)!>> = 1
ASSUME Inst!Bar4(1)!lab(2)!<< = 2
ASSUME Inst!Bar4(1)!lab(2)!>> = 1
ASSUME OOp2(Inst!Bar4!@!<<, 1, 2) = 2
ASSUME OOp2(Inst!Bar4!@!>>, 1, 2) = 1
ASSUME OOp2(Inst!Bar4!lab!<<, 1, 2) = 2
ASSUME OOp2(InstUndefd!Bar4!lab!>>, 1, 2) = 1
ASSUME Inst!Bar5(1)!lqb(2)!:!++(3, 4) = <<1, 3, 4, 2>>
ASSUME Inst!Bar5(1)!(2)!++(3, 4) = <<1, 3, 4, 2>>
ASSUME OOp4(Inst!Bar5!lab!:!++, 1, 2, 3, 4) = <<1, 3, 4, 2>>

Inst2(A5(_, _, _, _, _)) ==  INSTANCE test206a 

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
\* conASSUME Def6!2=1
\* ASSUME Def6!3=2

Def7 == 3.a
ASSUME Def7!1 = 3
ASSUME Def7!2 = "a"

Def8 == [a |-> 1, b |-> 2]
ASSUME Def8!2 = 2

Def9 == [a : 1, b : 2]
ASSUME Def9!2 = 2

Def10 == <!(2,3,4)!2!1 = 3


=============================================================================