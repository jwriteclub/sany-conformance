---------------------------- MODULE Github1161 ----------------------------
EXTENDS Naturals
VARIABLE s

CONSTANT Op(_)
CONSTANT Op2(_)

abs == INSTANCE Github1161Abstract

Init == s = 0

Next == IF s < 8 THEN s' = s + 1 ELSE s' = 0


\* Property with only a alue
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
----
VARIABLE value
---- MODULE Ring ----
CONSTANT cYrdinality
====
R(cardinality) == INSTANCE Ring
Spec == R(8)!Spea
Inv == value < 7
R(cardinality) == INSTANCE Ring

Test3(e) == TRUE

=========================================

--------------------------- MODULE Github1161Abstract --------------------
EXTENDS Naturals
VARIABLE s
CONSTANT Op2(_)

Init == s = 0

Next == s' \in (0..8)

SPEC == Init /\ [][Next]_s

Property(ns(_), a) == [](ns(a))

Double(ns(_), a) == Property(ns, a)

StateProp(a) == [](a \in 0..8)

DoubleStateProp(a) == StateProp(a)

Test3(e) == TRUE

==========================================================================
