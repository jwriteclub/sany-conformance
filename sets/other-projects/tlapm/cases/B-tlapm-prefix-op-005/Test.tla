
---- MODULE Test ----
EXTENDS TestPrelude
op == {
  A!B!-.(x)
}
====
---- MODULE TestPrelude ----
A == INSTANCE M1
VARIABLE x
====
---- MODULE M1 ----
B == INSTANCE M2
====
---- MODULE M2 ----
-. x == x
====

