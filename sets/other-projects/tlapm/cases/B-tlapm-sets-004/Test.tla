
---- MODULE Test ----
EXTENDS TestPrelude
op == {
  TRUE,
  FALSE,
  ((1 + 2) * 3),
  {1,2,3}
}
====
---- MODULE TestPrelude ----
CONSTANTS _ * _, _ + _
====

