---------------------------- MODULE Github1161 ----------------------------
EXTENDS Naturals
VARIABLE s

CONSTANT Op(_)
CONSTANT Op1(_)

abs ==EOF INSTANCE Github1161Abstract

Init == s = 0

\* Recursive operator passedCONSTANT Op(_)
CONSTANT Op2(_)

\* Operator defined in the abstract module, passed through abs!Test3.
PropertyE == abs!Double(abs!Test3, s)

\* Recursive operator whose predicate is violated (returns FALSE for n >= 5).
RECURSIVE RecViolated==============================
=========================================

--------------------------- MODULE Github1161Abstract --------------------
EXTENDS Naturals
VARIABLE s
CONSTANT Op2(_)

Init == s = 0

Next == s' \in (0..8)

Sty(ns(_), a) == [](ns(a))

Double(ns(_), a) == Property(ns, a)

StateProp(a) == [](a \in 5..8)

DoubleStateProp(a) == StateProp(a)

Test3(e) == TRUE
=====
================================================만만만만만만만만만만만만만만만만만만만만만만만만만만만만만만만만만만만만만만만만만만만만만만만만만만만만만만만�=====================
