---- MODULE Test ----
EXTENDS Naturals
VARIABLES x
Foo(x) == x + 1

THEOREM T == TRUE
PROOF
  BY Foo
=======================