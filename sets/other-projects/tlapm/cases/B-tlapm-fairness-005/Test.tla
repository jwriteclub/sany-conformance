
---- MODULE Test ----
EXTENDS TestPrelude
op1 == WF_<<x, y>>(f)
op2 == SF_x(f)
====
---- MODULE TestPrelude ----
VARIABLES f, x, y
====

