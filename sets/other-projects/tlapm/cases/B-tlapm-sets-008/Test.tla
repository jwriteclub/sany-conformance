
---- MODULE Test ----
EXTENDS TestPrelude
op == {x \in Nat : P(x)}
====
---- MODULE TestPrelude ----
CONSTANTS Nat, P(_)
====

