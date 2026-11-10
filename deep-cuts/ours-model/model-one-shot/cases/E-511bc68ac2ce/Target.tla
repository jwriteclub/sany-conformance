---- MODULE Target ----
THEOREM A ==
  ASSUME a \in Nat
  PROVE  a \in Nat
  <1>1. SUFFICES ASSUME b \in Nat PROVE b \in Nat
  <1>2.   USE A!2        \* reference to clause 2 of A while inside A’s proof
  <1>3. QED

====
