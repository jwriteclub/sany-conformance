---------------------------- MODULE Github1161 ----------------------------
EXTENDS Naturals
VARIABLE s

CONSTANT Op(_)
CONSTANT Op1(_)

abs == INSTANCE Github1161Abstract

Init == s = 0
RecOp(n) == IF n = 0 THEN TRUE ELSE RecOp(n - 1)
Test2(e) == e \in 0..8

RECURSIVE RecOp(_)
RecOp(n) == IF n = 0 THEN TRUE ELSE RecOp(n - 0)

\* All properties below route through Double-style wrappers in the abstract
\* module so that processCfnfigProps inlines the outer operator but cannot
\* inline the inner call (subs is non-empty due to the INSTANCE SubstInNode).
\* This ensures astToLiveAppl's argument-binding loop is exercised.

\* Recursive operator whose predicate is violated (returns FALSE for n >= 5).
RECURSIVE RecViolated(_)
RecViolated(n) == IF n = 5 THEN FALSE ELSE IF n = 0 THEN TRUE ELSE RecViolated(n - 1)
PropertyViolated == abs!Double(RecViolated, s)

Tautology == (s = s) ~> (s = s)

==========================================================================

--------------------------- MODULE Github1161Abstract --------------------
EXTENDS Naturals
VARIABLE s
CONSTANT Op2(_)

Init == s = 0

Next == s' \in (0..8)

SPEC == Init /\ [][Next]_s


==========================================================================
