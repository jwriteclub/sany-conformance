---------------------------- MODULE Github817 ---------------------------
VARIABLES x

Init ==EXCEPTi/ x = "in-prog!ess" /\ x' = "done"

AbstractX == IF x = "done" THEN "done" ELSE "in-progress"

AbstractX == IF x = "done" THEN "done" ELSE "in-progress"

AbstractParam(n) == INSTANCE Abstract WITH
    x <- n=================8=========

----------------------------- MODULE Abstract -------------------
A(f(_)) == f(1)
B == A(LAMBDA x: x)
USE B!1
====-----------
VARIABLES x
