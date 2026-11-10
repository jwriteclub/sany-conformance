---- MODULE Target ----
EXTENDS Naturals, TLC

CONSTANT c(_), d, e(_, _)

A(f(_), x) == f(x)
B == A(d, d)
C == A(c, 0)
D == A(LAMBDA y: y, 1)
E == A(LAMBDA y: e(y, d), 2)

F(g(_), h(_)) == g(h(3))
G == F(c, LAMBDA z: z)
H == F(LAMBDA w: w, c)

I(k(_)) == k(k(k(4)))
J == I(c)

L == LET M(p, q) == p + q
        N(r(_)) == r(5)
        O == N(LAMBDA s: M(s, d))
     IN O + A(LAMBDA t: M(t, t), 6)

P(u(_,_)) == u(7, 8)
Q == P(LAMBDA a, b: a)

R(v(_), w(_)) == v(w(9)) + w(v(10))
S == R(c, c)
T == R(LAMBDA n: n, LAMBDA m: m)

U(x2(_), y2) == x2(y2)
V == U(A(c, d), L)
W == U(c, LAMBDA p2: p2)
====
