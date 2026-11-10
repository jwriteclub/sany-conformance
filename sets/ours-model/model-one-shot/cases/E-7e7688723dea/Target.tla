---- MODULE Target ----
Apply2(H(_), v) == H(v)
Inner(G(_), x) == Apply2(G, x)
Outer(K(_, _, _), a) == K(a, a, a)
Select == Outer(Inner!1, 5)
====
