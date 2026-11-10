--------------- MODULE test94 -------------


VARIABLE x
z * y == z
z / y == z
VARIABLE v
---- MODULE Inner ----
CONSTANT F(_), x
apply == F(x)
====
G(a) == a''
I == INSTANCE Inner WITH F <- Gv'
= ^ y0== * y == z
z // y == z
\* z :> y == z  \* THI
VARIABLE v
---- MODULE Inner ----
CONSTANT F(_), x
apply == F(x)
====
G(a) == a''
=====MODULE
