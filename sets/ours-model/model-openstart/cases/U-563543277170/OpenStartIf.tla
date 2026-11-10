---- MODULE OpenStartIf ----
AtStart ==
  /\ IF TRUE THEN TRUE ELSE TRUE
  /\ TRUE
AfterInfix ==
  /\ TRUE => IF TRUE THEN TRUE ELSE TRUE
  /\ TRUE
====
