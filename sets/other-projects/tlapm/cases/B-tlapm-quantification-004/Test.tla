
---- MODULE Test ----
EXTENDS TestPrelude
op1 == \A x, y \in Nat, <<z>> \in Int, a, b, c \in Real : FALSE
op2 == \E <<x, y>> \in Nat, a, b \in Int : TRUE
====
---- MODULE TestPrelude ----
CONSTANTS Nat, Real, Int
====

