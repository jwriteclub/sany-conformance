---- MODULE Outer ----
Op(b) == INSTANCE Inner WITH C <- b
THEOREM (Op(1)!Sum(2,3)) = 5
====
