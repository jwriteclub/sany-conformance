---- MODULE Target ----
---- MODULE Inner ----
VARIABLE y
THEOREM B == ENABLED []y
======================
VARIABLE z
M == INSTANCE Inner WITH y <- z
============================================================
