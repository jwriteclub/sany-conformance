---- MODULE Test ----
VARIABLE x
THEOREM a3 == []ASSUME x = 0 []PROVE x = 0
THEOREM a4 == ASSUME x = 0, x = 1, x = 2 PROVE x = 0
====
