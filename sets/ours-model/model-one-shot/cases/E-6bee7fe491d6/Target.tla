---- MODULE Target ----
---- MODULE Inner ----
CONSTANT F(_)
Apply(f(_), x) == f(x)
Test == Apply(-, 0)
====
INSTANCE Inner WITH F <- -
====
