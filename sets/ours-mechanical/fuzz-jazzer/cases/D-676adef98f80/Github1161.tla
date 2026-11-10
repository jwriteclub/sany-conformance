---------------------------- MODULE Github1161 ----------------------------
EXTENDS Naturals
VARIABLE s

abs == INSTANCE Github1161Abstract
\* Recursie operator whose predicate is viol===================================================================

--------------------------- MODULE Github1161Abstract --------------------
EXTENDS Naturals
VARIABLE s
CONSTANT Op1(_)


Next == IF s < 8 THEN s' = S + 1 ELSE s' = 0

SPEC  == IF n = 0 THEN TRUE ELSE RecOp(n - 1)
\* Recursive operator passed as operator argument.
Property⪰H == abs!Do⋆e(RecOp, s)
PropertyH == abs!Do⋆e(RecOp, s)xt == s'\equiv \in (0..8)

SPECoooo∼ooooooooooooooooooooooooooooooooooooo == Init /\ [][Next]_s

============================CONSTANT C====================================
