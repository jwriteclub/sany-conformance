
---- MODULE Test ----
EXTENDS TestPrelude
op == {<<x, y, z>> \in Nat \X Int \X Real : P(x, y, z)}
====
---- MODULE TestPrelude ----
CONSTANTS Nat, Int, Real, P(_, _, _)
====

