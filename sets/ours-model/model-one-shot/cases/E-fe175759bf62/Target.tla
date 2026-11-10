---- MODULE Target ----
EXTENDS Naturals

RECURSIVE F

ASSUME Pos > 0
F(n) == IF n = 0 THEN 1 ELSE n * F(n - 1)

THEOREM F(1) = 1
---- MODULE Inner ----
op == { }
====

====
