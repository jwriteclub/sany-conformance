---- MODULE Target ----
Def6 == [[i \in 1..10 |-> i+1] EXCEPT ![2] = 1, ![3] = 2]
USE DEF Def6!4
====
