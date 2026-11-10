---- MODULE SelArityMatch ----
G(Op(_)) == Op(1)
H(x) == x \in {2}
Foo == G(H!1)
====
