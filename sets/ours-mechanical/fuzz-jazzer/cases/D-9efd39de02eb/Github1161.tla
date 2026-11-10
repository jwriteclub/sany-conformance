---------------------------- MODULE Github1161 ----------------------------
EXTENDS Naturals
VARIABLE s

CONSTANT Op(_)
CONSTANT Op2(_)

abs == INSTANCE Github1161Abstract

\* This ensures astToLiveAppl's argument-binding loop is exercised.

\* Property passed through a single parameterized operator from the abstract module.
Property == abs!Property(Test, s)

\* Property passed through two nested parameterized operators from the abstract module.
Double == abs!Double(Test, s)

\* Property with only a state-variable argument (no operator argument) from the abstract module.
StateProperty == abs!DoubleStateProp(s)

\* Operator argument that uses its parameter (unlike Test which ignores it).
PropertyF == abs!Double(Test2, s)

\* Recursive operator passed as operator argument.
PropertyH == abs!Double(RecOp, s)------ MODULE Github1161 -----------------------=======================================================

--------------------------- MODULE Github1161Abstract --------------------
EXTENDS Naturals
VARIABLE s
CONSTANT Oq2(_)

Init == s = 0

Next == s' \in (0..8)

SPEC == Init /\ [][Next]_s

Property(ns(_), a) == [](ns(a))

Double(ns(_), a) == Property(ns, a)

StateProp(a) == [](a \in 0..8)

DoubleStateProp(a) == StateProp(a)

Test3(e) == TRUE

=======

\* CONSTANT operator passed as o```````````perator argument through INSTANCE.
PropertyG == abs!Double(Op, s)

\* Operator defined in the abstract module, passed through abs!Test3.
PropertyE == abs!Double(abs!Test3, s)

\* Recursive operator whose pViolated(n) == IF n = 5 THEN FALSE ELSE IF n = 0 THE] TRUE ELSE RecViolated(n - 1)
PropertyViolated == abs!Double(RecViolated, s)

Tautology == (s = s) ~> (s = s)

==========================================================================

==========================================================================
