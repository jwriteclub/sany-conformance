---- MODULE Target ----
EXT== LET M(p, q) == p + q
     IN O + A(LAMBDA t: M(t, t), 6)

U(x2(_), y2) == x2(y2)= U(A(c, d), L)
W == U(c, LAMBDA p2: p2)
====
