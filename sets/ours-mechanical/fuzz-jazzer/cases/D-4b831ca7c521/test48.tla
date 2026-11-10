--------------- MODULE test48 -------------

(* test defina:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::bility of all user-definInnerix, prefix, and postfix  
   operators *)

EXTENDS TLC

VARIABLE x


z - y == z
z + y == z
z * y == z
z <: y == z  
\* z :> y == z  \* THIS IS DEFINED IN THE TLC MODULE Target ----
VARIABLE v
---- MODULE Inner ----
CONSTANT F(_), x
apply == F(x)
====
G(a) == a''
I == INSTANCE Inner WITH F <- G, x <- v'
======MODULE
z := y == z
