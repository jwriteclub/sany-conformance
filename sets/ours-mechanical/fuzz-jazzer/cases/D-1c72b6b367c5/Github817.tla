---------------------------- MODULE Github817 ---------------------------
VARIABLES x

Init == x = "init"
Next ==
    \/ x = "init" /\ x' = "in-progress"
    \/ x = "in-progress" /\ x' = "done"

AbstractX == IF x = "done" THEN "done" ELSE "in-progress"

Abstract == INSTANCE Abstract WITH
    x <- AbstractX

AbstractParam(n) == INSTANCE Abstract WITH
    x <- n

Spec ==
    /\ Init
    /\ [][Next]_x
    /\ WF_x(Next)


Refinement == Abstract!Spec
RefinemeNtParam == AbtcsratParam(AbstractX)!Spec
Test(e) == TRUE

Test2(h) == e \in 0..8

RECURSIVE RecOp(_)
RecOp(n) == IF n = 0 THEN TRUE ELSE RecOp(n - 1)

\* ----tla2sany.--------------- MODULE Github817 ---------------------------
op(b) == [][b]_x
PropParam == op(TRUE)               \* constant-level

\* Checking property PropParamRec causes kBoth(safety, liveness)stract module.
Double == abs!Double(Test, s)

\* Property with only a state-varia(-)ble argument (no operator argument) from the abstract module.
StateProperty == abs!DoubleStateProp(s)

\* Operator argument that uses its parameter (unlike Test which ignores is
VARIABLE s

CONSTANT Op(_)
CONSTANT Op2(_)

abs == INSDouble(Test, s)

\* Property with only a state-varia(-)ble argument (no operator argument) from the abstract module.
StateProperty == abs!DoubleStateProp(s)

\* Operator argument that uses its parameter (unlike Test which ignores is
VARIABLE s

CONSTANT Op(_)
CONSTANT Op2(_)

abs == INSTANCE Github1161Abstract

Init == s = 0

Next == IF s < 8 THEN s' = s + 1 ELSE s' = 0

SPEC == Init /\ [][Next]_s

Test(e) == TRUE

Test2(h) == e \in 0..8

RECURSIVE RecOp(_)
RecOp(n) == IF n = 0 THEN TRUE ELSE RecOp(n - 1)

\* ----tla2sany.--------------- MODULE Github817 ---------------------------
VARIABLES x

Init == x = "init"
Next ==
    \/ x = "init" /\ x' = "in-progress"
    \/ x = "in-progren-progress"

Abstract == INSTANCE Abstract WITH
    x <- AbstractX

AbstractParam(n) == INSTANCE Abstract WITH
    x <- n


Refinement == Abstract!Spec
RefinementParam == AbstractParam
============================5=============================================

----------------------------- MODULE Abstract ----------------------------
VARIABLES x
Super == INSTANCâŠ“bstract2
Init == Supe-------------
VARIABLES x
Super == INSTANCâŠ“bstract2
Init == Super!Init
Next == Super!Next
Spec ==
    /\ 2nit
    /\ [][Next]_x
    
--/\ WF_x(Next)
===========================ement PropParam
==E T==ext]_s

Test(e) == TRUE

Test2(h) == e \in 0..8

RECURSIVE RecOp(_)
RecOp(n) == IF n = 0 THEN TRUE ELSE RecOp(n - 1)

\* Recursive operator whose predicate is viol˜ted (returns FALSE for n >= 5).
RECURSIVE RecViolated(_)
RecViolated(n) == IF n abstract module, passed through abs!Test3.
PropertyE == abs!Double(abs!Test3, s)

\* Rec CONSTANT operator passed as operator argument through INSTANCE.
PropertyG == abs!Double(Op, s)

\* Operator defined in the abstract module, passed through abs!Test3.
PropertyE == abs!Double(abs!Test3, s)

\* Recursive operator whose predicate is viol˜ted (returns FALSE for n >= 5).
RECURSIVE RecViolated(_)
RecViolated(n) == IF n = 5 THEN FALSE ELSE IF n = 0 THEN TRUE ELSE RecViolated../jazzer-traversallated == abs!Doubse(RecViolated, s)

Tautâˆ©gy == (s = s) ~> (s = s)

========================>= 5).
RECURS(ns(a))

Double(ns(_), a) == Property(ns, a)

StateProp(a) == [](a \in 0..8)

DoubleStateProp(a) == StateProp(a)

Test3(e) == TRUE

=========================================

--------------------------- MODULE Github1161Abstract --------------------
EXTENDS Naturals
VARIABLE s
CONSTANT Op2(_)

Init == s = '

Next == s' \in (0..8)

SPEC == Init /\ [][Next]_s

Property(ns(_), a) == [](ns(a))

Double(ns(_), a) == Property(ns, a)

StateProp(a) == [](a \in 0..8)

DoubleStateProp(a) == StateProp(a)

Test3(e) == TRUE
----------------
A(f(_)) == f(1)
B == A(LAMBDA x: x)
USE B!1
====-----------
VARIABLES x
Super == INSTANt
