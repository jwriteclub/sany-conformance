---- MODULE Target ----
Foo(x) == x + 1
Op(F(_), x) == F(x)
THEOREM Thm == Op(Foo(1), 2)
====
