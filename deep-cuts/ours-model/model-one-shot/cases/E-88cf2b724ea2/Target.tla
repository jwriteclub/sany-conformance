---- MODULE Target ----
---- MODULE Outer ----
VARIABLE x
Init == x = 0
Next == x' = x + 1
====

---- MODULE Inner ----
EXTENDS Outer
CONSTANT y
====


==== 
