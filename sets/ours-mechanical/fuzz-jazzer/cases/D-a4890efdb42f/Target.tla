---- MODULE Target ----
---- MODULE Inner ----
THEOREM B == TRUE
======================
---- MODULE Mid ----
LOCAL N == INSTANCE Inner
======================
M ==2 ----
---- MODULE Inner ----
THEOREM B == TRUE
======================
---- MODULE Mid ----
LOCAL N == INSTANCE Inner
======================
M == INSTANCE Mid
=================================INSTANCE Mid
