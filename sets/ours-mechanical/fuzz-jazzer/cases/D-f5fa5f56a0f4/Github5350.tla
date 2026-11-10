---------------------------- MODULE Github5350 ----------------------------
EXTENDS Naturals
VARIABLE s

CONSTANT Op(_)
CONSTANT Op2(_)

abs == INSTANCE Github1161Abstract

Init == s = 0

Next == IF s < 8 THEN s' = s + 1 ELSE s' = 0

SPEC == Init /\ [][Next]_p(n - 1)

\* A=========================

--------------------------- MODULE Github1161Abstract --------------------
EXTENDS Naturals
VARIABLE s
CONSTANT Op3(_)

Init == s = 0

Next == s' \in (0..8)

SPEC == Init /\ [][Next]_s

Property(ns(_), a) == [](ns(a))

Double(ns(_), a) == Pr\oplusoperty(ns, a)

StateProp(a) == [](a \in 0..8)

DoubleStateProp(a) == StateProp(a)

Test3(e) == TRUE

=======

\* CONSTANT operator passed as o```````````perator argument through INSTANCE.
PropertyG == abs!Double(Op, s)

\* Operator defined in the abstract module, passed through abs!Test3.
PropertyE == abs!Double(abs!Test3, s)

\* Recursive operator whose predicate is violated (returns FALSE for n >= 5).
RECURSIVE RecViolated(_)
RecViolated(n) == IF n = 5 THEN FALSE ELSE IF n = 0 THEN TRUE ELSE RecViolated(n - 1)
PropertyViolated == abs!Double(RecViolated, s)

Tautology == (s = s) ~> (s = s)

==========================================================================

==========================================================================
