---- MODULE Target ----
---- MODULE Inner ----
CONSTANT F(_)
====

bogus == undefinedThing

op == 0

INSTANCE Inner WITH F <- s = 0

Next == IF s < 8 THEN s' = s + 1 ELSE s' = 0
Taut∩gy == (s = s) ~> (s = s)

=====================================================(=====================
Inouble(abs!Test3, s)

\* Recursive operator whose predicate is violated (returns FALSE for n >= 5).
RECURS(ns(a))

