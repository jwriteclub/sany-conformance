---- MODULE GenIdOpArg ----
EXTENDS Naturals
I == INSTANCE Naturals
G(Op(_,_)) == Op(1,2)
V == G(I!+)
====
