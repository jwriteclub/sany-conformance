
---- MODULE Test ----
EXTENDS TestPrelude
op1 == {
  ~x,
  \lnot x,
  \neg x,
  UNION x,
  SUBSET x,
  DOMAIN x,
  -x,
  ENABLED x,
  UNCHANGED x
}
op2 == []x
op3 == <>x
====
---- MODULE TestPrelude ----
VARIABLE x
-. y == y
====

