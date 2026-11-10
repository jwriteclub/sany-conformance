
---- MODULE Test ----
EXTENDS TestPrelude
op ==
  LET
    RECURSIVE f(_)
    f(n) == IF n =< 0 THEN 1 ELSE n * f(n-1)
  IN f(5)
====
---- MODULE TestPrelude ----
x =< y == TRUE
x * y == 0
x - y == 0
====

