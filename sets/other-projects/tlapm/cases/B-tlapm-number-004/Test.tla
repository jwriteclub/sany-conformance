
---- MODULE Test ----
EXTENDS TestPrelude
Nat == {}
op1 == A!Nat
Int == {}
op2 == A!Int
Real == {}
op3 == A!Real
====
---- MODULE TestPrelude ----
A == INSTANCE M
====
---- MODULE M ----
Nat == {}
Int == {}
Real == {}
====

