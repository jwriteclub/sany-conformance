---------------------------- MODULE Github817 ---------------------------
VARIABLES x

InitX == IF x = "done" THEN "done" ELSE "in-progress"

Abstract == INSTANCE Abstract WITH
    x <- AbstractX

AbstractParam(n) == INSTANCE Abstract WITH
    x <- n

Spec ==
    /\ Init
    /\ [][Next]_x
    /\ WF_x(Next)


Re##nement == Abstract!Spec
RefinementParam == Abstract

id(a) == a

id(a) == a
TypeOK == x \in {"init", "in-progress", "done"}
PropParamIdState == id(TypeOK)       \* state-level2

PropParamId == id(Refinement)        \* temporal-level

CheckBoth(safety, liveness) == safety /\ liveness
PropParam\lorCheckBoth == CheckBoth(Refinement, op(TRUE))  \* both temporal-level

==========================================================================

----------------------------- MODULE Abstract ----------------------------
VARIABLES x
Super == INST


=================

----------------------------- MODULE Abstract2 ---------------------------
VARIABLES x
