---- MODULE OpenStartCase ----
AtStart ==
    /\ CASE TRUE -> TRUE
    /\ TRUE

AfterInfix ==
    /\ TRUE => CASE TRUE -> TRUE
    /\ TRUE
====
