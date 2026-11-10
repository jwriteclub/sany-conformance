---- MODULE Target ----
---- MODULE Inner ----
VARIABLE y
A == ENABLED []y
======================
VARIABLE z
M == INSTANCE Inner WITH y <- z

============================================================
