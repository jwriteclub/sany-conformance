---- MODULE RichExport ----
EXTENDS Naturals
CONSTANT C, Op(_)
VARIABLE x
---- MODULE In ----
CONSTANT K
VARIABLE v
THEOREM TI == ASSUME NEW y \in K, NEW STATE s, NEW ACTION a, NEW TEMPORAL t, NEW VARIABLE w, NEW CONSTANT d PROVE lab :: y = y
ASSUME AI == K = K
F(p) == lb :: p + 1
====
I == INSTANCE In WITH K <- C, v <- x
J(q) == INSTANCE In WITH K <- q, v <- x
LOCAL L == 1.5
Dec == \b101 + \O7 + \HaF
G(P(_)) == P(1)
H == G(LAMBDA z : z) + G(I!F) + G(I!F!lb) + G(Op) + G(~)
R == [ i \in 1..2 |-> [x EXCEPT ![1] = @ + 1] ]
A1 == \A a \in C, <<b,c>> \in C : a = b
A2 == I!AI
THEOREM Thm == ASSUME NEW CONSTANT e PROVE e = e
<1> DEFINE Df == 1
<1> USE DEF Df
<1> INSTANCE In WITH K <- C, v <- x
<1>a. SUFFICES ASSUME TRUE PROVE TRUE
  OBVIOUS
<1>b. TRUE BY <1>a, Thm!1
<1> HIDE <1>b
<1>c. CASE TRUE
<1>d. PICK z \in C : z = z
  OBVIOUS
<1>e. HAVE TRUE
<1>f. TAKE u \in C
<1>g. WITNESS 1 \in C
<1> QED
  OBVIOUS
====
