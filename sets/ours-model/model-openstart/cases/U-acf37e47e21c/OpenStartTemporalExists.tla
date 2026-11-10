---- MODULE OpenStartTemporalExists ----
AtStart ==
  /\ \EE x : [] (x = x)
  /\ TRUE
AfterInfix ==
  /\ TRUE => \EE x : [] (x = x)
  /\ TRUE
====
