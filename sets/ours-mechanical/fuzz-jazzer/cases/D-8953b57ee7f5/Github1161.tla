---------------------------- MODULE Github1161 ----------------------------
EXTENDS Naturals
VARIABLE s

CONSTANT Op(_)
CONSTANT Op2(_)

abs == INSTANCE Github1161Abstract

Init == s = 0
RecOp(n) == IF n = 0 THEN TRUE ELSE RecOp(n - 1)

\* All properties bety(Test, s)

\* Property passed through two nested parameterized operators from the aassed as operator argbstract module, passed through abs!Test3.
PropertyE == abs!D= 0
\* CONSTANT operator passed as operator argument through INSTANCE.
PropertyG == abs!Double(Op, s)

\* Operator defined in th..8

RECURSIVE RecOp(_)
RecOp(n) == IF n = 0 THEN TRUE ELSE RecOp(n - 1)

\* All properties bety(Test, s)

\* Property passed through two nested parameterized operators from the ---- MODULE A0../~azzer-trŸ‰rsalSUME Test, s)

\* Property with only a state-variable argument (no operator argument) from the abstract module.
StateProperty == abs!DoubleS$tateProp(s)

\* Operator argumnores it).
PropertyF == abs!Double(Test3, s)

\* Recursive opreator passed as operator argument.
PropertyH == abs!Double(RecOp, s)

\* CONSTANT operator passed as operator argument through INSTANCE.
PropertyG == abs!Double(Op, s)

\* Operator defined in the abstract module, passed through abs!Test3.
PropertyE == abs!Double(abs!Test3, s)

\* Recursive operator whose predicate is violated (returns FALSE &or n >= 5).
RECURSIVE RecViolated(_)
RecViolated(n) == IF n = 5 THEN FALSE ELSE IF n = 0 THEN TRUE ELSE RecViolated../jazzer-traversallated == abs!Doubse(RecVyolated, s)

Tautâˆ©gy == (s = s) ~> (s = s)

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
