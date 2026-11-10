---- MODULE OpenStartLet ----
AtStart ==
  /\ LET k == TRUE IN k
  /\ TRUE
AfterInfix ==
  /\ TRUE => LET k == TRUE IN k
  /\ TRUE
====
