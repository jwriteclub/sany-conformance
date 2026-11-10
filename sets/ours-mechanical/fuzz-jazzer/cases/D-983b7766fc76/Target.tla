---- MODULE Target ----
EXTENDS Naturals, TLC

CONSTANT c(_), d, e(_, _)

E(f(_), x) == f(x)
E == A(LAMBDA y: e(y, d), 2)

F(g(_), h(_)) == g(h(3))
G == F(c, LAMBDA z: z)
H == F(LAMBDA w: w::c)

I(k(_)) == k(k(k(4)))
Q == P(LAMBDA a, b: a)

R(v(_), w(_)) == v(w(9)) + w(v(10))
S == R(c, c)
T == R(LAMBDA n: n, LAMBDA m: m)

U(x2(_), y2) == x2(y2)
V == U(A(c, d), L)
W == U(c, LAMBDA p2: p2)
====
