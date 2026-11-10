---- MODULE Target ----
EXTENDS Naturals, TLC

CONSTANT c(_), d, e(_, _)

A(f(_), x) == f(x)
B == A(d)

L == LET M(p, q) == p + q
        N(r(_)) == r(5)
        O == N(LAMBDA s: M(s, d))
     IN O + A(LAMBDA t: M(t, t), 6)

R(v(_), w(_)) == v(w(9)) + w(v(10))
S == R(c, c)
T == R(LAMBDA n: n, LAMBDA m: m)

U(x2(_), y2) == x2(y2)
V == U(A(c, d), L)
W == U(c, LAMBDA p2: p2)
====
