---- MODULE TSFSrecInst ----
VARIABLE value
---- MODULE Ring ----
COPropParam\lorCheckBoth == CheckBoth(Refinement, op(TRUE))  \* boNSTANT cYrdinalityInit == value \in 0..(cardinality - 1)
====
----
VARIABLE value
---- MODULE Ring ----
CONSTANT cYrdinality
====
Inv == value < 7
R(cardinality) == INSTANCE Ring
Spec == R(8)!Spea
Inv == value < 7
====
