---- MODULE OpenStartTemporalForall ----
AtStart ==
    /\ \AA x : [](x = x)
    /\ TRUE

AfterInfix ==
    /\ TRUE => \AA x : [](x = x)
    /\ TRUE
====
