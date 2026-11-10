
---- MODULE Test ----
EXTENDS TestPrelude
op == {f(x, y, z) : x \in Int, y, z \in Real}
====
---- MODULE TestPrelude ----
CONSTANTS Int, Real, f(_, _, _)
====

