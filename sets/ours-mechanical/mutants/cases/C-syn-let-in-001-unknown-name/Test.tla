
---- MODULE Test ----
op ==
  LET
    RECURSIVE fUndefd(_)
    f(n) == IF n =< 0 THEN 1 ELSE n * f(n-1)
  IN f(5)
====

