---- MODULE Target ----
Def6 == [[i \in 1..10 |-> i+1] EXCEPT ![3] = 2]
THEOREM Def6!2 ====
