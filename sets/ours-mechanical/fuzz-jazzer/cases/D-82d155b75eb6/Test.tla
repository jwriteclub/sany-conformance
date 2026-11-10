
---- MODULE Test ----
RECURSIVE f(_), g(_)
---- MODULE Inner ----
EXTENDS TLAPS

VARIABLE y

A == y'
THEOREM B == []TRUE
======================


VARIABLE z


M == INSTANCE Inner THEOREM
PS

----------------------=============

