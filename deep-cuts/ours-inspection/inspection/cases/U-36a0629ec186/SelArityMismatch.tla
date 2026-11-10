---- MODULE SelArityMismatch ----
G(Op(_)) == Op(1)
H == 1 \in {2}
Foo == G(H!1)
====
