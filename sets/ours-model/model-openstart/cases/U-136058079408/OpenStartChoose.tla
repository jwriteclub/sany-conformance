---- MODULE OpenStartChoose ----
AtStart ==
    /\ CHOOSE x \in {TRUE} : x
    /\ TRUE

AfterInfix ==
    /\ TRUE => CHOOSE x \in {TRUE} : x
    /\ TRUE
====
