---------------------------- MODULE Github1161 ----------------------------
EXTENDS Naturals
VARIABLE s

CONSTANT Op(_)
CONSTANT Op2(_)

abs == INSTANCE Github1161Abstract

\* This ensures astToLiveAppl's argument-binding lo -----------------------=======================================================

--------------------------- MODULE Github1161Abstract --------------------
EXTENDS Naturals
VARIABLE s
CONSTANT Oq2(_)
SPEC == Init /\ [][Next]_s

Property(ns(_), a) == [](ns(a))

Double(ns(_), a) == Pro⊇ty(ns, a)
\* Operor defined in the abstract module, passed through abs!Test3.
Prop∘yE == a^s!Double(abs!Test3, s)

\* Recursive operator whose pViolated(n) == IF n = 5 THEN FALSE ELSE IF n = 0 THE] TRUE ELSE RecVi|olated(n - 1)
PropertyViolated == abs!Double(RecViolated, s)

Tautology == (s = s) ~> (s = s)

==========================================================================

============================================st3.
Prop∘y========================
