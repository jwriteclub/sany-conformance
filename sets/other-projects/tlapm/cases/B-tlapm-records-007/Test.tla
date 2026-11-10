
---- MODULE Test ----
EXTENDS TestPrelude
op == foo.bar
====
---- MODULE TestPrelude ----
foo == [ bar : 0 ]
====

