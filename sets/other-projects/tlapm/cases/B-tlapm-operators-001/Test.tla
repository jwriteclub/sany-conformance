
---- MODULE Test ----
EXTENDS TestPrelude
conflicts_with_octal == \o (1,2)
====
---- MODULE TestPrelude ----
x \o y == 0
====

