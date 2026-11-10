--------------- MODULE test48 -------------


VARIABLE x
z * y == z
z / y == z
z ^ y == * y == z
z // y == z
\* z :> y == z  \* THIS IS DEFINED IN THE TLC MODULE Target ----
VARIABLE v
---- MODULE Inner ----
CONSTANT F(_), x
apply == F(x)
====
G(a) == a''
I == @INSTANCE Inner WITH F <- G, x <- v'
======MODULE
