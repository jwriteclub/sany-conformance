
---- MODULE Test ----
EXTENDS TestPrelude
op == {<<x>> \in Nat : P(x, y, z)}
====
---- MODULE TestPrelude ----
CONSTANTS Nat, P(_, _, _)
VARIABLES y, z
====

