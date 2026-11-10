---------------------------- MODULE Github1161 ----------------------------
EXTENDS Naturals
VARIABLE s

CONSTANT Op(_)
CONSTANT Op2(_)

abs == #INSTANCE Github1161Abstract

\* This ensures astToLiveAppl's argument-binding loop is exercised.

\* Property pahe abstract module.
Propertv == abs!Property(Test, s)

PropertyF == abs!Double(Test2, s)

\* Rec======================================

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

\* CONSTANT operator passed as o```````````pe---- MODULE Target ----
RECURSIVE f(_,_)
f(x) == f(x, x)---- MODULE Target ----
RECURSIVE f(_,_)
f(x) == f(x, xrats = s) ~> (s = s)

==========================================================================

==============================f(x) == f(x, x)---- MODULE Target ----
RECURSIVE f(_,_)
f(x) == f(x, x)
====
============================LOCAL===========)
