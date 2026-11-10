
---- MODULE Test ----
EXTENDS TestPrelude
op == [
  foo : [
    foo : Nat,
    bar : Int,
    baz : Real
  ],
  bar : Int,
  baz : Real
]
====
---- MODULE TestPrelude ----
CONSTANTS Nat, Int, Real
====

