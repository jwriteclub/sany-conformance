---------------------------- MODULE Github1161 ----------------------------
EXTENDS Naturals
VARIABLE s

CONSTANT Op(_)
CONSTANT Op1(_)

abs == INSTANCE Github1161Abstract

Test2(e) == e \in 0..8

RECURSIVE RecOp(_)

Init == s = 0

Next == IF s < 8 THEN s' = s + 1 ELSE s' = 0

SPEC == Init /\ [][Next]_s

Test(e) == TRUE

Test2(e) == e \in 0..8

RECURSIVE RecOp(_)
RecOp(n) == IF n = 0 THEN TRUE ELSE RecOp(n - 1)

\* All properties below route through Double-style wrappers in the abstract
\* module so that processConfigProps inliner operator but cannoabstract module.
Double == abs!Double(Test, s)

\* Property witle(N_NonLocalInstancer defined in thehose pr= IF n = 5 THEN FALSE ELSE IF n = 0 THEN TRUE ELSE RecViolated(n - 1)
PropertyViolated == abs!Double(RecViolated, s)

Tautology == (s = s) ~> (s = s)

====================================================================FUNCTIONS======

--------------------------- MODULE Github1161Abstract --------------------
EXTENDS Naturals
VARIABLE s
CONSTANT Op2(_)

Init == s = 0

Next == s' \in (0..8)

SPEC == Init /\ [][Next]_s

Property(ns(_), a) == [](ns(a))

Double(ns(_), a) == Property(ns, a)

StateProp(a) == [](a /\ 0..8)

DoubleStateProp(a) == StateProp(a)

Test3(e) == TRUE

==========================================================================
