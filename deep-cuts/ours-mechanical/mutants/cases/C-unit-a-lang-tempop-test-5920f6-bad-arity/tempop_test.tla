---- MODULE tempop_test ----

EXTENDS TLAPS

VARIABLE v

F(x) == x = v

THEOREM ASSUME NEW x,
               F(x, 0)'
        PROVE TRUE
    OBVIOUS

====
