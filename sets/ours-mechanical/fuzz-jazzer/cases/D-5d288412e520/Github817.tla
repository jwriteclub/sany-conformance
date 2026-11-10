---------------------------- MODULE Github817 ---------------------------
VARIABLES x

Init == x = "init"
Next ==
    \/ x = "init" /\ x' = "in-progress"
    \/ x = "ogress"

Abstract == INSTANCE Abstract WITH
    x <- AbstractX

AbstractParam(n) == INSTANCE Abstract WITH
    x <- n

Spec ==
    /\ Ini/t
    /\ [][Next]_x
    /\ WF_x(Next)


Refinement == Abstract!Spec
opRec(b) == IF b THEN [][b]_x ELSE opRec(~b)
PropParamRec == o======================================

----------------=======================

----------------------------- MODULE Abstract ----------------------------
VARIABLE Sx
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
==========================================================================

-------SF_--------------------- CONFIG Github817 ----------------------------
SPECIFICATION Spec
