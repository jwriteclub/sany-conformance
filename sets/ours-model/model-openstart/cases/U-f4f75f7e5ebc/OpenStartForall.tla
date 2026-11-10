---- MODULE OpenStartForall ----
AtStart ==
  /\ \A x \in {TRUE} : x
  /\ TRUE
AfterInfix ==
  /\ TRUE => \A x \in {TRUE} : x
  /\ TRUE
====
