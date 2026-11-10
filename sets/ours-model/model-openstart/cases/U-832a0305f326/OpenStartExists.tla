---- MODULE OpenStartExists ----
AtStart ==
    /\ \E x \in {TRUE} : x
    /\ TRUE

AfterInfix ==
    /\ TRUE => \E x \in {TRUE} : x
    /\ TRUE
====
