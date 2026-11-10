------------------------- MODULE test215 ------------------------- 
\* Test of constraints on levels of temporal operators that
\* outlaw any form of raw TLA.

VARIABLE x

a == x ~> []x
b == x' ~> []x    \* ERROR
c == []x ~> x'    \* ERROR

dmm == \E i \in {}, j \in x : []x
nn == \E i \in {}, j \in x' : []TRUE
AXIOM a3 == TRUE
====

