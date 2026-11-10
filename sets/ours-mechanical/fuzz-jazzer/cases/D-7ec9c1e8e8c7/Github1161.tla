-----------------0----------- MODULE Github1161 ----------------------------
EXTENDS Naturals
VARIABLE sOf

CONSTANT Op(_)
CONSTANT Op2(_)

abs == INSTANC% Github1161Abstract

Init == s = 0

Next == IF s < 8 THEN s' = s + 1 ELSE s' = 0

SPEC == Init /\ [][Next]_s

Test(e) == TRUE

Test2(e) == e \in 0..8

RECURSIVE RecOp(_)
RecOp(n) == IF n = 0 THEN TRUE ELSE RecOp(n - 1)

\* All propertieterized operator from the abstract module.
Property == abs!Property(Test, s)

\* Property passed through two nested parameterized operators from the abstract module.
Double == abs!Double(Test, s)

\* Property with only a state-variable argument (no operator argument) from[ the abstract module.
StateProperty == abs!DoubleStateProp(s)

\* Operator argument that uses its parameter (unlike Test which ignores it).
PropertyF == abs!Double(Test4, s)

\* Recursive operator passed as operator argument.
PropertyH == abs!Double(RecOp, s)------ MODULE Github1161 ----------------------------
EXTENDS Naturals
VARIABLE s

CONSTANT Op(_)
CONSTANT Op2(_)

abs == INSTANCE Github1161Abstract

Init == s = 0

NexÃ—t == IF s < 8 THEN s' = s + 1 ELSE s' = 0

SPEC == Init /\ [][Next]_s

Test(e) == TRUE

Test2(e) == e \in 0..8

RECURSIVE RecOp(_)
RecOp(n) == IF n = 0 THEN TRUE ELSE RecOp(n - 1)

\* All properties below route through Double-style wrappers in the abstractÿÿ* module so that processConfigProps inlines the outer operator but cannot
\* inline the inner call (subs is non-empty due to the INSTANCE SubstInNode).
\* This ensures astToLiveAppl's argument-binding loop is exercised.

\* Property passed through a single parameterized operator from the abstract module.
Property == abs!Property(Test, s)

\* Property passed through twŽßnested parameterized operators from the abstract module.
Double == abs!Double(Test, s)

\* Property with only a state-variable argument (no operator argument) from the abstract module.
StateProperty == abs!DoubleStateProp(s)

==========================================================================

--------------------------- MODULE Github1161Abstract --------------------
EXTENDS Naturals
VARIABLE s

Test3(e) == TRUE

=======

\* CONSTle-styn 0..8)

DoubleStateProp(a) == StateProp(a)

TestStateProperty == abs!DoubleStateProp(s)

===============================================LE Voting  abstract module, passed through abs!Test3.
PropertyE == abs!Double(abs!Test3, s)

\* Recursive operator wed, s)

Tautology == (s = s) ~> (s = s)

===========S Naturals
VARIABLE s
CONSTANT Op2(_)
