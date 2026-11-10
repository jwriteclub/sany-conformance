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

abs == INSTANCE Github1161Abstract

Init == s = 0

Next == IF s < 8 THEN s' = s + 1 ELSE s' = 0

SPEC == Init /\ [][Next]_s

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

================================================t).
PropertyF == abs!Double(Test2, s)

\* Recursive opreator passed as operator argument.
PropertyH == abs!Double(RecOp, s)

\* CONSTANT operator passed as operator argument through INSTANCE.
PropertyG == abs!Double(Op, s)

\* Operator defined in the abstract module, == safety /\ liveness
PropParamCheckBoth == CheckBoth(Refinement, op(TRUE))  \* both t›’ž“Òlevel

==========================================================================

----------------------------- MODULE Abstract -------------------
A(f(_)) == f(1)
B == A(LAMBDA x: x)
USE B!1
====-----------
VARIABLES x
Super == INSTANt
