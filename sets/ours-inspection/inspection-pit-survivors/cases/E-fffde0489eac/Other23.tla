\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: OpApplNode.java:645 RemoveConditional_EQUAL_IF: N's parameter loses non-Leibniz status: leibniz flags differ
\* Aim: OpApplNode.java:465 RemoveConditional_ORDER_ELSE: range not level checked: any bounded quantifier throws; hot line
\* Aim: OpApplNode.java:738 RemoveConditional_ORDER_ELSE: K1's q gets no constraint from the operator argument: max0 differs
\* Aim: OpApplNode.java:744 NegateConditionals: same
\* Aim: OpApplNode.java:744 RemoveConditional_EQUAL_ELSE [1109]: same
\* Aim: OpApplNode.java:744 RemoveConditional_EQUAL_ELSE [1112]: same
\* Aim: OpApplNode.java:817 RemoveConditional_ORDER_ELSE: range arg constraints lost: B2 mm differs
\* Aim: OpApplNode.java:818 VoidMethodCall: same
\* Aim: OpApplNode.java:884 RemoveConditional_ORDER_ELSE: range arg params lost: K3's max0 differs (the condition itself is not printed, see report)
\* Aim: OpApplNode.java:920 RemoveConditional_ORDER_ELSE: K4's max0 differs
\* Aim: OpApplNode.java:922 NegateConditionals: same
\* Aim: OpApplNode.java:922 RemoveConditional_EQUAL_ELSE [1906]: same
\* Aim: OpApplNode.java:922 RemoveConditional_EQUAL_ELSE [1909]: same
\* Aim: OpApplNode.java:772 RemoveConditional_EQUAL_IF: theorem Jt!TV reports Leibniz with its first arg flag false: K8's z becomes non-Leibniz
\* Aim: LetInNode.java:165 NegateConditionals: LET def's level constraints not merged: L1 max0 differs
\* Aim: LetInNode.java:165 RemoveConditional_ORDER_ELSE: same
\* Aim: LetInNode.java:166 RemoveConditional_EQUAL_ELSE: same
\* Aim: LetInNode.java:172 NegateConditionals: LET def's arg constraints not merged: L2 mm differs
\* Aim: LetInNode.java:172 RemoveConditional_ORDER_ELSE: same
\* Aim: LetInNode.java:173 RemoveConditional_EQUAL_ELSE: same
\* Aim: LetInNode.java:183 RemoveConditional_EQUAL_ELSE: LET def's arg params lost: K5's max0 differs
\* Aim: LetInNode.java:185 RemoveConditional_EQUAL_ELSE: same
\* Aim: LetInNode.java:191 RemoveConditional_ORDER_ELSE: LET INSTANCE's arg params lost: K6's max0 differs
\* Aim: OpDefNode.java:1049 RemoveConditional_EQUAL_ELSE: LET def d recomputed on the second recursive pass picks up R1's constraint on c: d's lc differs
\* Aim: OpDefNode.java:1049 RemoveConditional_ORDER_ELSE: same
\* Aim: ThmOrAssumpDefNode.java:555 PrimitiveReturns: It!T1's mm0,0 differs
\* Aim: ThmOrAssumpDefNode.java:466 RemoveConditional_ORDER_ELSE: same
\* Aim: ThmOrAssumpDefNode.java:468 RemoveConditional_EQUAL_IF: same
\* Aim: ThmOrAssumpDefNode.java:564 BooleanFalseReturn: K7's max0 differs
\* Aim: ThmOrAssumpDefNode.java:482 RemoveConditional_ORDER_ELSE: same
\* Aim: ThmOrAssumpDefNode.java:523 RemoveConditional_EQUAL_IF: It!T4's arg params lose (H, 0, ap)
---- MODULE Other23 ----
VARIABLE x
CONSTANT c, H(_)
M(a, b) == a /\ ENABLED b
N(p) == M(TRUE, p)
A1(G(_), a) == G(a)
H1(b) == b'
K1(q) == A1(H1, q)
B1(P(_)) == P(x)
B2(Q(_)) == \E i \in B1(Q) : TRUE
B3(R(_), a) == \E i \in R(a) : TRUE
K3(q) == B3(H1, q)
B4(R(_), a) == A1(R, a)
K4(q) == B4(H1, q)
L1(p) == LET d == p' IN 1
L2(P(_)) == LET d == B1(P) IN 1
L3(P(_), a) == LET d == P(a) IN 1
K5(q) == L3(H1, q)
---- MODULE Mm ----
CONSTANT G(_), a
LOCAL V == G(a)
====
IM == INSTANCE Mm WITH G <- H, a <- c
L4(P(_), b) == LET J == INSTANCE Mm WITH G <- P, a <- b IN 1
K6(q) == L4(H1, q)
RECURSIVE R1
R1 == LET d == R1 IN c'
---- MODULE Mt ----
VARIABLE v
CONSTANT G2(_), a2, H2(_)
Bk(P(_)) == P(v)
THEOREM T1 == Bk(G2)
THEOREM T2 == G2(a2)
THEOREM T4 == H2(a2)
====
It(Gp(_), ap) == INSTANCE Mt WITH v <- x, G2 <- Gp, a2 <- ap, H2 <- H
K7(q) == It(H1, q)!T2
---- MODULE Mv ----
CONSTANT e1, e2
THEOREM TV == e1 = e2
====
Jt(u, w) == INSTANCE Mv WITH e1 <- u, e2 <- w
Ar2(G(_,_), y) == G(y, y)
K8(z) == Ar2(Jt!TV, z)
====
