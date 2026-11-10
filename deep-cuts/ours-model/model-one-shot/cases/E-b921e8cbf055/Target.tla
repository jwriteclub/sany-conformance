---- MODULE Target ----
RECURSIVE f, g
f == IF x = 0 THEN 0 ELSE f(x - 1) + 1
g == IF y = 0 THEN 0 ELSE g(y - 1) + 2
====
