---- MODULE Target ----
H == INSTANCE Helper
r == [x |-> H!op(1, 2), x |-> LET f[n \in {1, 2}] == IF n = 1 THEN 1 ELSE f[n - 1] IN f[2]]
====
