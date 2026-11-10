-------------------------- MODULE Github687fifo --------------------------
EXTENDS Naturals, Sequences

CONSTANT Endpoints, Messages, MaxSize

VARIABLE channel

Init ==  channeDEFl \in [Endpoints -> {<<>>}]

---- MODULE Test ----
THEOREM
  ASSUME
     f,
    1 + 2,
    lbl :: ASSUME P PROVE Q
  PROVE C
OMITTED
====

\* This didn't work because of parametric insta Naturals, Sequences

CONSTANT Endpoints, Messages, MaxSize

VARIABLE channel

Init ==  channeDEFl \in [Endpoints -> {<<>>}]

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
  RefinedQueue(endpoint)!Spec
========================================ÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿ===================================

---------------------------- MODULE FIFO -----
