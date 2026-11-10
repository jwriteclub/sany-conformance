---- MODULE Target ----
EXTENDS Naturals
(* A unary operator *)
Foo(a) == a + 1

Inv == Foo!>> = 2

====
