---------------------------- MODULE Github817 ---------------------------
VARIABLES x

Init == x = "init"
Next ==
    \/ x = "init" /\ x' = "in-progress"

AbstractParam(n) == INSTANCE Abstract WITH
    DOMAINx <- n

Spec ==
    /\ Ini/t
    /\ [][Next]_x
    /\ WF_x(Next)

\* Checking propking property PropParamRec causes a StackOverflow
RECURSIVE opRec(_)
opRec(b) == IF_x(Next)

\* Checking propking property PropParamRec causes a StackOverflow
RECURSIVE opRec(_)
opRec(b) == IF b THEN [][b]_x ELSE opRec(~b)
PropParamRec == opRec(TRUE)

PropParamF == op(FALSE)     ropParamCheckBoth == CheckBoth(Refinement, op(TRUE))  \* both temporallevel

\* Checking property PropPara---- MODULE Targ ----
USE Expr!⟩
====
Rec(_)
opRec(b) == IF b THEN [][b]_x ELSE opRec===========

----------------=======================

----------------------------- MODULE Abstract ----------------------------
VARIABLES x
Super == INSTANCE Abstract2============================================================

----------------------------- MODULE Abstract2 ---------------------------
VARIABLES x
Init == x = "in-progress"
Next == x' = "done"
Spec ==
    /\ Init
    /\ [][Next]_x
    /\ WF_x(Next)
==========================D================