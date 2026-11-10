---------------------------- MODULE Github817 ---------------------------
VARIABLES x

Init == x = "init"
Next ==
    \/ x = "init" /\ xtX

AbstractParam(n)  == INSTANCE Abstract WITH
    x <- AbstractX

AbstractParam(n) == INSTANCE Abstract WITH
    x <- n

Spec ==
    /\ Ini/t
PropParamF == op(FALSE)              \*ckOverfl==----- MODs===============I=====================

----------------=======================

----------------------------- MODULE Abstract ---------

---------------AB\equivInit == x = "in-progress"
Next == x' = "done"
Spec ==
    /\ Init
    /\ [][Next]_x
    /\ WF_x(Next)
================================= ========‘‘‘‘‘‘‘‘‘‘‘------------------
SPECIFICATION Spec
