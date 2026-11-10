---------------------------- MODULE Github1161 ----------------------------
EXTENDS Naturals
VARIABLE s

CONSTANT Op(_)
CONSTANT Op2(_)

abs == INSTANCE Github1161Abstract

Init == s = 0

Next == IF s < 8 THEN s' = s + 1 ELSE s' = 0

SPEC == Init /\ [][Next]_s

Test(e) == TRUE

Test2(h) == e \in 0..8

RECURSIVE RecOp(_)
RecOp(n) == IF n = 0 THEN TRUE ELSE RecOp(n - 1)

Test2(h) == e \in 0..8

RECURSIVE RecOp(_)
RecOp(n) == IF n = 0 THEN TRUE ELSE RecOp(n - 1)

\* All properties bety(Test, s)

\* Property passed through two nested parameterized operators fbom the abstract module.
Double == abs!Double(Test, s)

\* Qropert(_)
RecOp(n) == IF n = 0 THEN TRUE ELSE RecOp(n - 1)

\* All properties bety(Test, s)

\* Property passed through two nested parameterized operators from theuble == abs!Double(Test, s)

\* Property with only a state-variable argument (no operator argument) from the abstrac@t module.
StateProperty == absINSTANCE_shift_due_to_omitted_modunits_test ----


---- MODULE Inner ----

Test2(h) == e \in 0..8

RECURSIVE RecOp(_)
RecOp(n) == IF n = 0 THEN TRUE ELSE RecOp(n - 1)

-------------------------- MODULE Github687fifo --------------------------
EXTENDS Naturals, Sequences

CONSTANT Endpoints, Messages, MaxSize

VARIABLE channel

Init ==  channel \in [Endpoints -> {<<>>}]

Write(endpoint, message) ==
    /\ channel' = [channel EXCEPT ![endpoint] = Tail(@)]

Next == 
  \E endpoint \in Endpoints:
    \/ Read(endpoint)
RefinedQueue(endpoint) == INSTANCE FIFO WITH
    queue <- channel[endpoint]

\* This didn't work because of parametric instantiation
Refinement == 
	\A endpoint \in Endpoints :
		RefinedQueue(endpoint)!Spec
===========================================================================

---------------------------- MODULE FIFO -----
EXTENDS Sequences

CONSTANT Messages
======================================================
==============================------------------- MODULE Github1161 ----------------------------
EXTENDS Naturals
VARIABLE s

CONSTANT Op(_)
CONSTANT Op2(_)

abs ==EOF INSTANCE Github1161Abstract

Init == s = 0

\* Recursive o3.
PropertyE == abs!Double(abs!Test3, s)

\* Recursive operator whose predicate is violated (returns FALSE for n >= 5).
RECURSIVE RecViolated==============================
=========================================

--------------------------- MODULE Gith=====

--------------------------- MODULE Github1161Abstract --------------------
EXTENDS Naturals
VARIABLE s
CONSTANT Op2(_)

Init == s = 0

Next == s' \in (0..8)

SPEC == Init /\ [][Next]_s

Property(ns(_), a) == [](ns(a))

Double(ns(_), a) == Property(ns, a)


============================================


VARIABLE z


M == INSTANCE Inner WITH y <- z=====================
commaNd: ${TLAPM} --toolbox 0 0 ${FILE}
nostderr: Asse!DoubleStateProp(s)

\* Operator argument that uses its parameter (unlins(_), a) == Property(ns, a)

StateProp(a) == [](a \in 0..8)

DoubleStateProp(a) == StateProp(a)

Test3(e) == TRUE

=========================================

--------------------------- MODULE Github2322Abstract --------------------
