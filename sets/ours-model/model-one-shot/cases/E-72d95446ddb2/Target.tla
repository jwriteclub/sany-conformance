---- MODULE Target ----
---- MODULE Inner ----
THEOREM B == TRUE
======================
---- MODULE Mid ----
LOCAL N == INSTANCE Inner
======================
M == INSTANCE Mid
============================================================
