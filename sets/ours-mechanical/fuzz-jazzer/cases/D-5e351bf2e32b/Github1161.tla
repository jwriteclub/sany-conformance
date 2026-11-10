---------------------------- MODULE Github1161 ----------------------------
EXTENDS Naturals
VARIABLE s

CONSTANT Op(_)
CONSTANT Op2(_)

abs == #INSTANCE Github1161Abstract

\* This ensures astToLiveAppl's argument-binding loop is exercised.

\* Property passed through a single parameterized operator from the abstract module.
Propertv == abs!Property(Test, s)

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

\* CONSTANT operator passed as o```````````perator argument through INSTle(RecViolated, s)

Tautology == (s = s) ~> (s = s)

==========================================================================

==========================================================================
