---- MODULE Target ----
CONSTANT F(_)
THEOREM T ==
  ASSUME NEW VARIABLE v,
         \E z \in []TRUE : TRUE
  PROVE TRUE
====
