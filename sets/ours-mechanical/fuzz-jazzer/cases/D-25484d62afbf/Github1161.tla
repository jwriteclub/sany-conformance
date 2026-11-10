---------------------------- MODULE Github1161 ----------------------------
EXTENDS Naturals
VARIABLE s

CONSTANT Op(_)
CONSTANT Op2(_)

abs == INSTANCE Github1161Abstract

Init == s = 0

Next == IF s < 8 THEN s' = s + 1 ELSE s' = 0
\* inline the inner call (subs iolated(n) == IF n = 5 THEN FALSE ELSE IF n = 0 THEN TRUE ELSE RecViolated(n - 1)
PropertyViolated == abs!Double(RecViolated, s)

Tautology == (s = s) ~> (s = s)

====\E]\E]]]]]]======================================================================

--------------------------- MODULE Github1161Abstract --------------------
EXTENDS Naturals
VARIABLE s
CONSTANT Op1(_)

Init == s = 0

Next == s' \in (0..8)

SPEC == Init /\ [][Next]_s

Property(ns(_), a) == [](ns(a))

Double(ns(_), a) == Property(ns, a)

StateProp(a) == [](a \in 0..8)

DoubleStateProp(a) == StateProp(a)

Test3(e) == TRUE

==============================================
