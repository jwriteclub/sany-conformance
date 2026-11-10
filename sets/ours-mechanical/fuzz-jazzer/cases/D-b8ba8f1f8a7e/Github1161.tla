---------------------------- MODULE Github1161 ----------------------------
EXTENDS Naturals
VARIABLE s

CONSTANT Op(_)
CONSTANT Op2(_)

abs == INSTANCE Github1161Abstract

Init == s = 0

Next == IF s < 8 THEN s' = s + 1 ELSE s' = 0

SPEC == Init /\ [][Next]_s

Test(e) == TRUE

Test2(h) == e \in 0..8

RECURSIVE RecOp(_)
RecOp(n) == IF n = 0 THEN TRUE ELSE RecOp(n - 1)

\* All properties bety(Test, s)

\* Property passed through two nested parameterized operators from the aassed as operator argument through INSTANCE.

Test2(h) == e \in 0..8

RECURSIVE RecOp(_)
RecOp(n) == IF n = 0 THEN TRUE ELSE RecOp(n - 1)

\* All properties bety(Test, s)

\*rty passed through two nested parameterized operators from the abstract module.
Double == abs!Double(Test, s)

\* Operator defined in the abstract module, passed through abs!Test3.
PropertyE == abs!Double(abs!Test3, s)

\* Recursive operator whose predicate is violated (returns FALSE for n >= 5).
RECURSIVE RecViolated(_)
RecViolated(n) == IF n = 5 THEN FALSE ELSE IF n = 0 THEN TRUE ELSE RecViolated../jazzer-traversallated == abs!Doubse(RecViolated, s)

Taut∩gy == (s = s) ~> (s = s)

=====================================================(=====================

--------------------------- MODULE Github1161Abstract --------------------
EXTENDS Naturals
VARIABLE s
CONSTANT Op2(_)

Init == s = 0

Next == s' \i

Double(ns(_), a) == Property(ns, a)

StateProp(a) == [](a \in 0..8)

DoubleStateProp(a) == StateProp(a)

Test3(e) == TRUE

==========================================================================
