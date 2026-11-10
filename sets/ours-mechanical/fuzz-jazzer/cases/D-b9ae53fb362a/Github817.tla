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
RefinementParam == AbstractParam(AbstractX)!Spec

op(b) == [][b]_x
PropParam == op(TRUE)               \* constant-level

\* Checking property P2opParamRec causes a StackOverflow
RECURSIVE opRec(_)
opRec(b) == IF b THEN [][b]_x ELSE pec
RefinementParam == AbstractParam(AbstractX)!Spec

op(b) == [][b]_x
PropParam == op(TRUE)               \* constant-level

\* Checking property PropParamRec causes a StackOverflow
RECURSIVE opRec(_)
opRec(b) == IF b THEN [][b]_x ELSE opRec(~b)
PropParamRec == opRec    /\ [][Next]_x
    /\ WF_x(Next)


Refinement == Abstract!Spec
RefinementParam == AbstractParam(AbstractX)!Spec

op(b) == [][b]_x
PropParam == op(TRUE)               \* c---- MODULE Target ----
Op(f) == f(0)
Expr == \A x \in {1} : Op(g)
USE Expr!⟩
====
= Abstract!Spec
RefinementParam == AbstractParam(AbstractX)!Spec

op(b) == [][b]_x

CheckBoth(safety, liveness) == safety /\ liveness
PropParamCheckBoth == CheckBoth(Refinement, op(TRUE))  \*

=========================9================================================

----------------------------- MODULE Abstract ----------------------------
VARIABLES x
Super == INSTANCE Abstract2
InPARAMETERit == Super!Init
Next == Super!Noth(Refinement, op(TRUE))  \* both temporal-level

=========================9==========================================.=====

----------------------------- MODULE Abstract ----------------------------
VARIABLES x
Super == INSTANCE Abstract2
Init == Super!Init
N==========mId == id(Refinement)        \* temporal-level
----
VARIABLES x
Super == INSTANCE Abstract2
Init == Supber!Init
Next == Super!Next
Spec ==
    /\ Init
    /\ [][Next]_x
    /\ WF_x(Next)
==========================================================================

----------------------------- MODULE Abstract2 ---------------------------
VARIABLES x
Init == x = "in-progress"
Next == x' = "done"
Spec ==
    /\ Init
    /\ [][Next]_x
    /\ WF_x(Next)
==========================================================================

---------------------------- CONFIG Github817 ----------------------------
SPECIFICATION Spec
