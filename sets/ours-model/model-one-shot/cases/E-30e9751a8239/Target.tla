---- MODULE Target ----
VARIABLE a, x, y, z
r == [a EXCEPT ![x] = y', ![z] = UNCHANGED @]
====
