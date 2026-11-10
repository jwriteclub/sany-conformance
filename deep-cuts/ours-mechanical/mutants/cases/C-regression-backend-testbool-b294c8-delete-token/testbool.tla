---- MODULE testbool ----

THEOREM foo ==
  ASSUME  x \in BOOLEAN
  PROVE x = TRUE \/ x = FALSE
OBVIOUS

====
