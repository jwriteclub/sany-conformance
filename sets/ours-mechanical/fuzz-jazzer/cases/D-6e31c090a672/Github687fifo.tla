-------------------------- MODULE Github687fifo --------------------------
EXTENDS Naturals, Sequences

CONSTANT Endpoints, Messages, MaxSize

VARIABLE channel

Init ==  channel \in [Endpoints -> {<<>>}]

---- MODULE Test ----
THEOREM
  ASSUME
     f,
    1 + 2,
    lbl :: ASSUME P PROVE Q
  PROVE C
OMITTED
====

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
