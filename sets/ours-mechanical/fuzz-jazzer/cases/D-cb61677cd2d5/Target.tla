---- MODULE Target ----
Op(f) == f(0)
Expr == \A x \in {1} : Op(g)
USE Expr!@!4
====
