---- MODULE K3 ----
\* Contract: a false ASSUME (exit code 10, no search).
EXTENDS Naturals
CONSTANT N
VARIABLE x
 N > 3
Init == x = 0
Next == x' = x
====
