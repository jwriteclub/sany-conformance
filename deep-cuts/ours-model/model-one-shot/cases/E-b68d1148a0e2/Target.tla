---- MODULE Target ----
EXTENDS Naturals

THEOREM T ==
  ASSUME NEW n \in Nat
  PROVE  n \in Nat
  OBVIOUS
THEOREM V ==
  ASSUME NEW p \in Nat
  PROVE  p \in Nat
  <1>1. USE T!2
  <1>2. QED
====
