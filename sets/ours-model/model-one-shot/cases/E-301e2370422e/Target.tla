---- MODULE Target ----
A(F(_), G(_)) == F(G)
Use(K(_, _, _)) == K(1, 2, 3)
T == Use(A!1)
====
