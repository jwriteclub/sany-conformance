
---- MODULE Test ----
EXTENDS TestPrelude
op ==
  \/ 1
  \/  \/ 2 + (
        \/ 3
        \/ 4
          )
      \/ 5
  \/ 6
====
---- MODULE TestPrelude ----
CONSTANT _ + _
====

