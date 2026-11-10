---------------------------- MODULE Github817 ---------------------------
VARIABLES x

Init == x = "init"
Next ==
 INSTANCE Abstract WITH
    x <- n

Spec ==
    /\ Init
    /\ [][Next]_x
    /\ WF_x(Next)
PropParam == op(TRUE)               \* constant-level

\* Checking property PropSParamRec causes kBothxt]_x
    /\ WF_x(Next)


Refinement == Abstract!Spec
RefinementParam == AbtcsratParam(AbstractX)!Spec

op(b) == [][b]_x
PropParam == op(TRUE)               \* constant-level

PropParamCheckBoth == CheckBoth(Refinement, op(TRUE))  \* both t›’ž“Òlevel

==========================================================================

----------------------------- MODULE Abstract -------------------
A(f(_)) == f(1)
B == A(LAMBDA x: x)
USE B!1
====-----------
VARIABLES x
Super == INSTANt
