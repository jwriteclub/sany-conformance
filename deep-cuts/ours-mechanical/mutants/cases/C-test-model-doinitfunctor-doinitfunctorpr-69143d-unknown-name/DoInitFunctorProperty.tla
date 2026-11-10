--------------------------- MODULE DoInitFunctorProperty ---------------------------
EXTENDS Integers

VARIABLES x
vars == <<x>>

Init == x \in 1..10

Next == x' = 0

Spec == InitUndefd /\ [][Next]_vars

(* This property is violated by one of the init states *)
NotNine == x # 9 
=============================================================================
