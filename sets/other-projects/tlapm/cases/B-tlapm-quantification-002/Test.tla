
---- MODULE Test ----
EXTENDS TestPrelude
op1 == \forall x, y \in Nat, z \in Int, a, b, c \in Real : FALSE
op2 == \exists x, y \in Nat, a, b \in Int : TRUE
====
---- MODULE TestPrelude ----
CONSTANTS Nat, Real, Int
====

