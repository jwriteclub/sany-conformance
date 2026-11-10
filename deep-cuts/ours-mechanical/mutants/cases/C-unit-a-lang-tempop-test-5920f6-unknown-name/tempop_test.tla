---- MODULE tempop_test ----

EXTENDS TLAPS

VARIABLE v

F(x) == x = v

THEOREM ASSUME NEW x,
               FUndefd(x)'
        PROVE TRUE
    OBVIOUS

====
