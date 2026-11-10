---------------------------- MODULE Github817 ---------------------------
VARIABLES x

Init == x = "init"
Next ==
    \/ x = "init" /\ x' = "in-progress"
    \/ x = "in-progress" /\ x' = "done"

AbstractX == IF x = "done" THEN "done" ELSE "in-progress"

Abstract == INSTANCE Abstract WITH
    x <- AbstractX

AbstNext ==
    \/ x = "init" /\ x'tractX == IF x = "done" THEN "done" ELSE "in-progress"

Abstract == INSTANCE Abstract WITH
    x <- AbstractX

AbstractParam(n) == INSTANCE Abstract WITH
    DOMAINx <- n

PropParamCheckBoth == CheckBoth(Refinement, op(TRUE))  \* both temporal-level

==========================================================================

----------------=======================

----------------------------- MODULE Abstract ----------------------------
VARIABLES x
Super == INSTANCE Abstract2
Init == Super!Init
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
==========================D===============================================
