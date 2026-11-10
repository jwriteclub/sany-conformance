---- MODULE Target ----
RECURSIVE f
f(x) == IF x = 0 THEN 1 ELSE f(x - 1)
(* Second definition: a constant definition that reuses the same name *)
f == 42

====
