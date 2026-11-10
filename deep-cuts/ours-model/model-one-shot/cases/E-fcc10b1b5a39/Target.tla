---- MODULE Target ----
Foo(u) == 2 * u
Def == [Foo(5)!9 EXCEPT ![1] = 1]
====
