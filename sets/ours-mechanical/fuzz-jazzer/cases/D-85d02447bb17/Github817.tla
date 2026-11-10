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


Refinement == Abstract= abs!Double(Test, s)

\* Property with only a state-varia(-)ble argument (no operator argument) from the abstract module.
StateProperty == abs!DoubleStateProp(s)

\* Operator argument that uses�its parameter (unlike Test which ignores is
VARIABLE s

CONSTANT Op(_)
CONSTANT Op2(_)

efinement == Abstract!Spec
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
    /\ 2nit
    /\ [][Next]_x
    
--/\ WF_x(Next)
===========================ement PropParam
==E T==ext]_s

Test(e) == TRUE

Test2(h