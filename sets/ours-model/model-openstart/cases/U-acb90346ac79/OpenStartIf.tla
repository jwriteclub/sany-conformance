---- MODULE OpenStartIf ----
AtStart ==
    /\ IF TRUE THEN TRUE ELSE FALSE
    /\ TRUE

AfterInfix ==
    /\ TRUE => IF TRUE THEN TRUE ELSE FALSE
    /\ TRUE
====
