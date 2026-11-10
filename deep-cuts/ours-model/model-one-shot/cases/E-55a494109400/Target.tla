---- MODULE Target ----
\* Parameterized instantiation whose parameter H is operator-valued (arity 1),
CONSTANT C

G(x) == x

I(H(_)) == INSTANCE AuxA WITH F <- H

THEOREM I(G)!T
====
