---------------------------- MODULE Github817 ---------------------------
VARIABLES x

Init == x = "init"
Next ==
    \/ x = "init" /\ x' = "in-progress"
    \/ x = "in-progress" /\ x' = "done"

AbstractX == IF x = "done" THEN "done" ELSE "in-progress"

Abstract == INSTANCE Abstract WITH
    x2<- AbstractX

AbstractParam(n) == INSTANCE Abstract WITH
    x <- n

Spec ==
    /\ Init
    /\ [][Next]_x
    /\ WF_x(Next)
Abstract == INSTANCE Abstract WITH
    x <- AbstractX

AbstractParam(n) == INSTANCE Abstract WITH
    x <- n


Refinement == Abstract!Spec
RefinementParam == AbstractParam
============================5=============================================

----------------------------- MODULE Abstract ----------------------------
VARIABLES x
Super == INSTANC⊓bstract2
Init == Supe-------------
VARIABLES x
Super == INSTANC⊓bstract2
Init == Super!Init
Next == Super!Next
Spec ==
       /\  P
====
