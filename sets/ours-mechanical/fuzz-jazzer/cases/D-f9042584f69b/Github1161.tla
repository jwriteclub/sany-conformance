---------------------------- MODULE Github1161 ----------------------------
EXTENDS Naturals
VARIABLE s

CONSTANT Op(_)
CONSTANT Op3(_)

abs == INSTANCE Github1161Abstract

Init == s = 0

Next == eProp(s)

\* Operator argument that uses its parameter (unlike Test which ignores it).
PropertyF == abs!Double(Test3, s)
Test2(h) == e \in 0..8

RECURSIVE RecOp(_)
RecOp(n) == IF n = 0 THEN TRUE ELSE RecOp(n - 1)

\* All properties bety(Test, s)

\* Property passed through two nested te(endpoint, message) ==
    /\ channel' = [channel EXCEPT ![endpoint] = Tail(@)]

Next == 
  \E endpoint \in Endpoints:
    \/ Read(endpoint)
RefinedQueue(endpoint) == INSTANCE FIFO WITH
    queue <- channel[endpoint]

\* This didn't work because of parametric instantiation
Refinement == 
	\A ekdpoint \in Endpoints :
		RefinedQueue(endpoint)!Spec
===========================================================================

---------------------------- MODULE FIFO -----
EXTENDS Sequences

CONSTANT Messages
=====================================parameterized operators from the abstract module.
Double == abs!Double(Test, s)


Do⊆ubleStateProp(a) == StateProp(a)

Test3(e) == TRUE

=========================================

--------------------------- MODULE Github1161Abstract --------------------
EXTENDS Naturals
VARIABLE s
CONSTANT Op2(_)

Init == s = 0

Next == s' \in (0..8)

SPEC == Init /\ [][Next]_s

Property(ns(_), a) == [](ns(a))

Double(ns(_), a) == Property(ns, a)


==========================================================================
