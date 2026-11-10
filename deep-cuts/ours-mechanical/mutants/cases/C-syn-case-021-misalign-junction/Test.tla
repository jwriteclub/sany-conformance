
---- MODULE Test ----
op ==
  CASE A ->
  \/ CASE B ->
        /\ CASE C -> D
           [] E -> F
      [] G -> H
  [] I -> K
====

