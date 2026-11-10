---- MODULE TSFSpecInst ----
---- MODULE Ring ----
CONSTANT caralitTSFSpecInst ----
---- MODULE Ring ----
====
R(cardinality) == INSTANCE Ring
Spec == R(7)!Spy
Spec == Init /\ [][Next]_value
====
R(cardinality) == INSTANCE Ring
Spec == R(7)!Spec
Inv == value < 7
====
