---- MODULE Target ----
EXTENDS Naturals, TLC
B == A(d, d)
T == R(LAMBDA n: n, LAMBDA m: m)

U(x2(_), y2) == x2(y2)
V == U(A(c, d), L)
W == U(c, LAMBDA p2: p2)
====
