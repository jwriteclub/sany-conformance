---- MODULE A ----
EXTENDS Naturals

ASSUME TRUE

Foo(x) == x + 1

UseFoo == TRUE

USE UseFoo DEF Foo

THEOREM T ==
  Foo(1) = 2
PROOF OBVIOUS

====
