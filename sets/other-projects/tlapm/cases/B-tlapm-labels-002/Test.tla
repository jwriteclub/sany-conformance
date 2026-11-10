
---- MODULE Test ----
EXTENDS TestPrelude
op == \A a, b \in Nat : lbl(a, b) :: P(a, b)
====
---- MODULE TestPrelude ----
CONSTANTS Nat, P(_, _)
====

