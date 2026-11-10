---- MODULE Target ----
\* Make an ASSUME clause whose semantic node is null (previous error),

A(p, q) == p

HO(f(_)) == f(1)

THEOREM t1 == ASSUME A(1) PROVE FALSE

THEOREM t2 == ASSUME HO(1) PROVE FALSE

====
