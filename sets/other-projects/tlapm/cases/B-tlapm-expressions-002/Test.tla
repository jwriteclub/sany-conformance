
---- MODULE Test ----
EXTENDS TestPrelude
op == <<a, <<b, c>>>>
====
---- MODULE TestPrelude ----
VARIABLES a, b, c
====

