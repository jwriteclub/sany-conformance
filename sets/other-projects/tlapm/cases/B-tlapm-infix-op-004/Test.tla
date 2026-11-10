
---- MODULE Test ----
EXTENDS TestPrelude
op == f(##, \otimes)
====
---- MODULE TestPrelude----
CONSTANTS _ ## _, _ \otimes _
f(x(_, _), y(_, _)) == 0
====

