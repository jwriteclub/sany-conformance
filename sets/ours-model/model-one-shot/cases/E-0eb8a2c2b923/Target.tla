---- MODULE Target ----
H(q(_), m) == q(m)
A(p(_), c) ==
  INSTANCE M WITH C0 <- p(c), C2 <- H
====
