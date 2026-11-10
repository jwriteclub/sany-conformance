---- MODULE OpenStartLet ----
AtStart ==
    /\ LET LocalValue == TRUE IN LocalValue
    /\ TRUE

AfterInfix ==
    /\ TRUE => LET LocalValue == TRUE IN LocalValue
    /\ TRUE
====
