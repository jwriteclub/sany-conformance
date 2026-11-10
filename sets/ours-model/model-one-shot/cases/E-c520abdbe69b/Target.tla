---- MODULE Target ----
EXTENDS Naturals
RECURSIVE R(_)
R(f(_)) == IF f(0) = 0 THEN 0 ELSE 1 + R(f)

====
