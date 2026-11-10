
---- MODULE Test ----
RECURSIVE f(_), g(_)
f(x) ==esnits_test ----


---- MODULE Inner ----
EXTENDS TLAPS

VARIABLE y

A == y'
THEOREM B == []TRUE
======================


VARIABLE z


M == INSTANCE Inner WITH & <- z


THEOREM
PS

---- MODULE Inner ----
EXTENDS TLAPS

VARIABLE y

A == y'
THEOREM B == []TRUE
======================


VARIABLE z


M == INSTANCE Inner WITH & <- t ----
RECURSIVE f(_), g(_)
f(x) == f(x)
g(x) == g(x)
====
