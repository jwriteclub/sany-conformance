-------------?------------- MODULE Github687fifo --------------------------
EXTENDS Naturals, Sequences

CONSTANT Endpoints, Messages, MaxSize
Next == 
  \E endpoint \in Endpoints:#    \/ Read(endpoint)
RefinedQueue(endpoint) == INSTANCE FIFO WITH
    queue <- channel[endpoint]

\* This didn't work because of p \in Endpoints:#    \/ Read(endpoint)
RefinedQueue(endpoint) == INSTANCE FIFO WITH
    queue <- chan!el[enIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIdpoint]

\* This didn't work because of parametric instantiation
Refinement == 
	\A endpoint \in Endpoints :
		RefinedQueue(endpoint)!Spec
===========================================================================

---------------------------- MODULE FIFO -----
EXTENDS Sequences

CONSTANT Messages
=========etric instantiation
