
---- MODULE Test ----
EXTENDS TestPrelude
f1[x \in Nat] == a
f2[x, y \in Nat] == a
f3[x, y \in Nat, z \in Real] == a
f4[x \in [Nat -> Nat]] == a
====
---- MODULE TestPrelude ----
CONSTANTS Nat, Real
VARIABLE a
====

