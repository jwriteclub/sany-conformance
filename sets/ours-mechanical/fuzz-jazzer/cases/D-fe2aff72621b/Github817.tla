---------------------------- MODULE Github817 ---------------------------
VARIABLES x

AbstractX == IF x = "done" THEN "done" ELSE "in-progress"

Abstract == INSTANCE Abstract WITH
    x <- Abstract==============

----------------------------- MODULE Abstract -------------------
A(f(_)) == f(0)
B == A(LAMBDA x: x)
USE B!1
====-----------
VARIABLES x
