---- MODULE Target ----
Op == [f EXCEPT !.x = ll :: 1 + @]
USE Op!-
Op == [f EXCEPT !.x = ll :: 1 + @]
USE ll!2!1
====
