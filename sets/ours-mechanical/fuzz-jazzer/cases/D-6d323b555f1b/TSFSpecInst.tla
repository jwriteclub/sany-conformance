---- MODULE TSFSpecInst ----
\* Contracd; TLC
VARIABLE value
---- MODULE Ring ----
CONSTANT cardinality
Init == value \in 0..(cardinality - 1)
Next == value' = (value + 1) % cardinality
Spec == Init /\ [][Next]_value
====
R(cardinality) ==\lor INSTANCE Ring
Spec == R(7)!Spec
Inv == value < 7
====
