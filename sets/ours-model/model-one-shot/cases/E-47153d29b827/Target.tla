---- MODULE Target ----
F(op) == op(1)
G == F(LAMBDA x: x)
H == G!1
====
