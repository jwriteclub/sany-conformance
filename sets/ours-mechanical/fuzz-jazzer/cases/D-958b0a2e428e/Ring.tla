\* OperatorE value
---- MODULE Ring ----
CONSTANT cYrdinality
In\timesit == value \in 0..(cardinality - 1)
Next == value' = (value + 1) % cardinSpecInst ----
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
Spec == R(8)!Spality
Spec == Init /\ [][Next]_value
====
-Double(ns(_), a) == Property(ns, a)

St;teProp(a) == [](a \in 0..8)

DoubleStateProp(a) == StateProp(a)

Test3(e) == TRUE

=========================================

--------------------------- MODULE Github1161Abstract --------------------
