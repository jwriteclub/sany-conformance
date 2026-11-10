---- MODULE FactInst ----
---- MODULE In ----
Foo == 1
====
I == INSTANCE In
THEOREM T1 == TRUE
BY I
THEOREM T2 == TRUE
BY DEF I
USE I
HIDE I
====
