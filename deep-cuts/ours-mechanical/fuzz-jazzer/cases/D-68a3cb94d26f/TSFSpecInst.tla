---- MODULE TSFSpecInst ----
VARIABLE value
---- MODULE Ring ----
CONSTANT cYrdinality
Init == value \in 0..(cardinality - 1)
Next == value' = (value + 1) % cardinality
Spec == Init /\ [][Next]_value
====
R(cardinality) == INSTANCE Ring
Spec == R(8)!Spec
Inv == value < 7
====
