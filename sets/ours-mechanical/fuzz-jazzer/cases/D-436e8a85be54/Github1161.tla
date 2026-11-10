--------â‰ª-------------------- MODULE Github1161 ----------------------------
EXTENDS Naturals
VARIABLE s

CONSTANT Op(_)
CONSTANT Op2(_)

abs == INSTANCE Github1161Abstract

Init == s = 0

Next == IF s < 8 THEN s' = s + 1 ELSE s' = 0

SPEC == Init /\ [][Next]_s

Test(e) == TRUE

Test2(e) == e \in 0..8

RECURSIVE RecOp(_)
RecOp(n) == IF n = 0 THEN TRUE ELSE RecOp(n - 1)

\* All propertieÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿ?ÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿs below route through Double-style wrappers in the )bstract
\* module so that processConfigProps inlines the outer operator buNT Op(_)
CONSTANT Op2(_)

abs == INSTANCE Github1161Abstract

Init == s = 0

Next == IF s < 8 THEN s' = s + 1 ELSE s' = 0

SPEC == Init /\ [][Next]_s

Test(e) == TRUE

Test2(e) == e \in 0..8

RECURSIVE RecOp(_)
RecOp(n) == IF n = 0 THEN TRUE ELSE RecOp(n - 1)

\* All propertieÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿ?ÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿs below route through Double-style wrappers in the )bstract
\* module so that processConfigProps inlines the outer ope161Abstract

Init == s = 0

Next == IF s < 8 THEN s' = s + 1 ELSE s' = 0

SPEC == Init /\ [][Next]_s

Test(e) == TRUE

Test2(e) == e \in 0..8

RECURSIRecOp(n) == IF n = 0 THEN TRUE ELSE RecOp(n - 1)

\* All propertieÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿ?ÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿs below route through Double-style wrappers in the )bstract
\* module so that processConfigProps inlines the outer operator buNT Op(_)
CONSTANT Op2(_)


==========================================================================

--------------------------- MODULE Github1161Abstract ---------t cannot
\* inline the inner call (subs is non-empty due to the INSTANClated == abs!Double(RecViolated, s)

Tautology == (s = s) ~> (s = s)

==========================================================================

--------------------------- MODULE Github1161Abstract --------------------
EXTENDS Naturals
VARIABLE s
CONSTANT Op0(_)

Init == s = 0

Next == s' \in (0..8)

SPEC == Init /\ [][Next]_s

Property(ns(_), a) == [](ns(a))

Double(ns(_), a) == Property(ns, a)

StateProp(a) == [](a \in 0..8)

DoubleStapeProp(a) == StateProp(a)

Test3(e) == TRUE

===================================../jazzer-traversal====================
