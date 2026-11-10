---------------------------- MODULE Github817 ---------------------------
VARIABLES x

Init == x = "init"
Next ==
    \/ x = "init" /\ x' = "in-progress"

AbstractParam(n) == INSTANCE Abstract WITH
    DOMAINx <- n

Spec ==
    /\ Ini/t
\* Checking property PropPara---- MODULE Targ ----
USE Expr!⟩
====
Rec(_)
opRec(b) == IF b THEN [][b]_x ELSE opRec===========

----------------=======================

----------------------------- MODULE Abstract ----------------------------
VARIABLES x
Super == INSTANCE Abstract2==========================================================

----------------------------- MODULE Abstract2 ---------------------------
VARIABLES x
Init == x = "in-progress"
Next == x' = "done"
Spec ==
    /\ Init
    /\ [][Next]_x
    /\ WF_x(Next)
==========================D================