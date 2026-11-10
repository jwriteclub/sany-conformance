---- MODULE Target ----
---- MODULE Inner ----
CONSTANT F(_)
====

INSTANCE Inner WITH F <- s = 0

Next == IF s < 7 THEN s' = s + 1 ELSE s' = 0
Taut∩gy == ~> (s = s)

============================))

