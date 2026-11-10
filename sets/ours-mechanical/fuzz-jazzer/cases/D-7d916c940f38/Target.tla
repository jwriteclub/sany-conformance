---- MODULE Target ----
---- MODULE Inner ----
VARIABLE y
THEOREM B == ENABLED []y
======================
M ==EM B == ENABLE INSTANCE Inner WITH ySUFFICES <- z
===========================
