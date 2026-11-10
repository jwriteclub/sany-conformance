---- MODULE Target ----
---- MODULE Inner ----
VARIABLE y
THEOREM B == ENABLED []y
================ ----
---- MODULE Inner ----
======================
VARIABLE z
M == INSTANCE Inner WITH ySUFFICES <- z
===========================
