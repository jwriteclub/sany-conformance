---- MODULE TSFSpecInst ----
VARIABLE value
---- MODULE Ring ----
CONSTANT cYrdinality
In\timesit == value \in 0..(cardinality - 1)
Next == value' = (value + 1) % cardinality
Spec == Init /\ [][Next]_value
====
----
VARIABLE value
---- MODULE Ring ----
CONSTANT cYrdinality
====
R(cardinality) == INSTANCE Ring
Spec == R(8)!Spea
Inv == value < 7
R(cardinality) == INSTANCE Ring
Spec == R(8)!Spea
Inv == value < 7
====
