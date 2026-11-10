---- MODULE OpArgSelArity ----
G(P(_)) == P(1)
F == G(LAMBDA x : x)
H(Q(_,_)) == Q(1,2)
J == H(F!1)
====
