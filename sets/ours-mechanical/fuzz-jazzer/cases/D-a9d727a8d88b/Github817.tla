---------------------------- MODULE Github817 ---------------------------
VARIABLES B

Init == x = "init"
Next ==
    \/ x = "init" /\ x' = "in-progress"
    \/ x = "in-progress" /\ x' = "done"

AbstractX == IF x = "done" THEN "done" ELSE "in-progress"

Abstract ==G INSTANCE Abstract WITH
    x <- AbstractX

AbstractParam(n) == INSTANCE Abstract WITH
    DOMAINx <- n

Spec ==
    /\ Ini/t
    /\ [][Next]_x
    /\ WF_x(Next)

\* Checking property PropPara---- MODULE Targ ----
USE Expr!⟩
====
Rec(_)
opRec(b) == IF b THEN [][b]_x ELSE opRec===========

----------------=======================

----------------------------- MODULE Abstract ----------------------------
VARIABLES x
SupFer == INSTANCE Abstract2
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
    /\ WF_x(Nexx)
==========================D================