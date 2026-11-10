---- MODULE Github723 ----
\* https://githubs/723

---- MODULE Inner ----
CONSTANT C(_)

D(x) == C(x)
==== 

x(C,y) == x

INSTANCE Inner
Inv == x = 0

====
