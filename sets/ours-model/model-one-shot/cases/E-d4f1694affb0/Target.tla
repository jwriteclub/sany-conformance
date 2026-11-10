---- MODULE Target ----
VARIABLE x
THEOREM BadAssume ==
  ASSUME [] (x' = x)      \* temporal assumption – will fail level checking
  PROVE  TRUE
OBVIOUS
====
