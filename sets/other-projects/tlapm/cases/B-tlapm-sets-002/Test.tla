
---- MODULE Test ----
EXTENDS TestPrelude
op1 == {x \in Nat}
op2 == {x \in BOOLEAN, y \in STRING}
====
---- MODULE TestPrelude ----
CONSTANT Nat
VARIABLES x, y
====

