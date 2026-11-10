---- MODULE OpenStartFalseAfterInfix ----
Bad ==
  /\ TRUE =>
  IF TRUE THEN TRUE ELSE TRUE
====
