---- MODULE QualPrefixOp ----
---- MODULE In ----
~~ == 1
a ++ b == a
-. a == a
a ^+ == a
====
I == INSTANCE In
G(P(_,_)) == P(1,2)
X == G(I!++)
Y == I!++(1,2)
Z == I!-.(1)
W == I!^+(1)
====
