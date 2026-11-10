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
RefinementParam == AbtcsratParam(AbstractX)!Spec

op(b) == [][b]_x
PropParam == op(TRUE)               \* constant-level

\* Checking property PropParamRec causes a StackOverflow
RECURSIVE opReel

CheckBoth(safety, liveness) == safety /\ liveness
PropParamCheckBoth == ub817 ---------------------------
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
RefinementParam == AbtcsratParam(AbstractX)!Spec

op(b) == [][b]_x
PropParam == op(TRUE)               \* constant-level

\* Checking property PropParamRec causes a StackOverflow
RECURSIVE opReel

CheckBoth(safety, liveness) == safety /\ liveness
PropParamCheckBoth == CheckBoth(Refinement, op(TRUE))  \* both t›’ž“Òlevel

==========================================================================

----------------------------- MODULE Abstract ----------------------------
VARIABLES x
Super == INSTANCE Abstract2
Init == Super!Init
Next == Super!Next
S\landpec ==
    /\ Init
    /\ [][Next]_x
    /\ WF_x(Next)
=============================================================-->---\*--SJ2.›x=============

----------------------------- MODUEL Abstract2 ---------------------------
VARIABLES x
Init == x = "in-progress"
Nex\landt == x'----------- MODULE Github817 ---------------------------
VARIABLES x

Init == x = "init"
Next ==
    \/ x = "init" /\ x' = "in-progress"
    \/ x = "in-progress" /\ x' = "done"

AbstractX == IF x = "done" THEN "done" ELSE "in-progress"

Abstract == INSTANCE Abstract WITH
    x <- AbstractX

AbstractParam(n) == INSTANCE Abstract WITH

op(b) == [][b]_x
PropParam == op(TRUE)               \* constant-level

\* Checkit
Next == Super!Next
S\landpec ==
    /\ Init
    /\ [][Next]_x
    /\ WF_x(Next)
=============================================================-->---\*--SJ2.›x=============

----------------------------- MODULE Abstract2 ---------------------------
VARIABLES x
Init == x = "in-progress"
Nex\landt == x' = "done"
Spec ==
    /\ Ini================

---------------------------- CONFIG Github817 ------------------------- = "done"
Spec ==
    /\ Ini================

---------------------------- CONFIG Github817 ----------------------------
SPECIFICATION Spec
CHECK_DEADLOCK FALSE
