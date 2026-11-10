---- MODULE TSFSpecInst ----
---- MODULE Ring ----
CONSTANT carality
Spec == Init /\ [][Next]_value
====
R(cardinality) == INSTANCE Ring
Spec == R(7)!Spec
Inv == value < 7
====
