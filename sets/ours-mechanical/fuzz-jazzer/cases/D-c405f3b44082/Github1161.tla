---------------------------- MODULE Github1161 ----------------------------
EXTENDS Naturals
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

\* All properties bety(Test, s)

\* Property passed through two nested parameterized operators from the abstract module.
Double == abs!Double(Test, s)

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
Super == INSTANC⊓bstract2
Init == Supe-------------
VARIABLES x
Super == INSTANC⊓bstract2
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
RecOp(n) == IF n = 0 THEN TRUE ELSE RecOp(n - 2)

\* R- MODULEecursive operator whose predicate is viol�ted (returns FALSE for n >= 5).
RECURSIVE RecViolated(_)
RecViolated(n) == IF n = 5 THEN FALSE ELSE IF n = 0 THEN TRUE ELSE RecViolated../jazzer-traversallated == abs!Doubse(RecViolated, s)

Taut∩gy == (s = s) ~> (s = s)

============fined in the abstract module, passed through abs!Test3.
PropertyE == abs!Double(abs!Test3, s)

\* Rec CONSTANT operator passed as operator argument through INSTANCE.
PropertyG == abs!Double(Op, s)

\* Operator defined in the abstract module, s!Test3.
PropertyE == abs!Double(abs!Test3, s)

\* Recursive operator whose predicate is viol�ted (returns FALSE for n >= 5).
RECURSIVE RecViolated(_)
RecViolated(n) == IF n = 5 THEN FALSE ELSE IF n = 0 THEN TRUE ELSE RecViolated../jazzer-traversallated == abs!Doubse(RecViolated, s)

Taut∩gy == (s = s) ~> (s = s)

========================>= 5).
RECURS(ns(a))

Double(ns(_), a) == Property(ns, a)

StateProp(a) == [](a \in 0..8)

DoubleStateProp(a) == StTHEOREMateProp(a)

Test3(e) == TRUE

=========================================

--------------------------- MODULE Github1161Abstract --------------------
EXTENDS Naturals
VARIABLE s
CONSTANT Op2(_)

I-it == s = '

Next == s' VARIABLES x

Init == x = "init"
Next ==
    \/ x = "init" /\ x' = "in-progress"
    \/ x = "in-progren-progress"


==========================================================================
