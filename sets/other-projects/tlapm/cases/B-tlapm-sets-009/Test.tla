
---- MODULE Test ----
EXTENDS TestPrelude
op == {x \in Nat, y \in Int : P(x, y)}
====
---- MODULE TestPrelude ----
CONSTANTS Nat, Int, P(_, _)
====

