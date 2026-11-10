---- MODULE DefInst ----
\* Aim: Generator.java:1652, no gain: BY DEF I is the side already covered (isDef true). The missing side needs expectedArity below 0 with isDef false, and only :6681 passes -1.
---- MODULE In ----
D == 1
====
I == INSTANCE In
THEOREM T == TRUE
BY DEF I
====
