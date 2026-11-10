---- MODULE Target ----
op(f) == [(lbl :: f) EXCEPT ![1] = @ + 1]
====
