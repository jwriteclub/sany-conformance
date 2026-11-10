---- MODULE Target ----
ASSUME Thm == TRUE
Op(x) == \A y \in {1} : LET dummy == TRUE IN Thm

USE Op(1)!(2)!Thm
====
