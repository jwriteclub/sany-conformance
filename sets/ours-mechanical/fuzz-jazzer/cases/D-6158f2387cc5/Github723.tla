---- MODULE Github723 ----
\* https://githubs/723

---- MODULE Inner ----
CONSTANT C(_)

D(x) == C(x)
==== 

x(C,y) == x

INSTANCE Inner

VARIABLE x
Spec == x = 0 /\ [][x' = D(x)]_x
Inv == x = 0

====
