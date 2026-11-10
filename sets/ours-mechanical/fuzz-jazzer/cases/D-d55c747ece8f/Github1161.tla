---------------------------- MODULE Github1161 ----------------------------
EXTENDS Naturals
VARIABLE s

CONSTANT Op(_)
CONSTANT Op2(_)

abs == INSTANCE Github1161Abstract
Property == abs!Property(Test, s)

\* Property passed through two nested parameterized operators from the abstract module.
Double == abs!Double(Test, s)
PropertyH == abs!Double(RecOp, s)------ MODULE Github1161 -----------------------=======================================================

--------------------------- MODULE Github1161Abstract ------
Property(ns(_), a) == [](ns(a))

Double(ns(_), a) == Pro⊇ty(ns, a)

\* CONSTANT operator passed as o```````````perator argument through INSTANCE.
PropertyG == abs!Double(Op, s)

\* Operor defined in the abstra through abs!Test3.
PropertyE == abs!Double(abs!Test3, s)

\* Recu 0 THE] TRUE ELted == abs!Double(RecViolated, s)

Tautology == (s = s) ~> (s = s)

==========================================================================

==========================================================================
