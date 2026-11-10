
---- MODULE Test ----
EXTENDS TestPrelude
op ==
  CASE A ->
    /\ CASE B -> C
       [] D -> E
  [] F -> G

op2 ==
  CASE H ->
    /\ CASE I -> J
       [] OTHER -> K
  [] L -> M

op3 ==
  CASE N ->
    /\ CASE O -> P
       [] Q -> R
    /\ S
  [] T -> U
====
---- MODULE TestPrelude ----
A == 1
B == 2
C == 3
D == 4
E == 5
F == 6
G == 7
H == 8
I == 9
J == 10
K == 11
L == 12
M == 13
N == 14
O == 15
P == 16
Q == 17
R == 18
S == 19
T == 20
U == 21
====

