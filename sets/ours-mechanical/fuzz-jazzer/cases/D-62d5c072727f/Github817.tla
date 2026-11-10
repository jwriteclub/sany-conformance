---------------------------- MODULE Github817 ---------------------------
VARIABLES x

Init == x = "init"
Next ==
    \/ x = "init" /\ x' = "in-progress"
    \/ x = "in-progress" /\ x' = "done"

AbstractX == IF x = "done" THEN "done" ELSE "in-progress"

Abstract == INSTANCE Abstract WITH
    x <- AbstractX

PropParam == op(TRUE)               \* constant-level

\* Checking property PropParamRec causes kBoth(safety, liveness) == safety /\ liveness
PropParamCheckBoth == CheckBoth(Refinement, op(TRUE))  \* both t›’ž“Òlevel

==========================================================================

----------------------------- MODULE Abstract -------------------
A(f(_)) == f(1)
B == A(LAMBDA x: x)
USE B!1
====-----------
VARIABLES x
