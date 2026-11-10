---- MODULE Target ----
RECURSIVE f
f == IF x = 0 THEN 1 ELSE f(x - 1)
(* Second definition – reuses the same OpDefNode but is not marked recursive *)
f == 42

====
