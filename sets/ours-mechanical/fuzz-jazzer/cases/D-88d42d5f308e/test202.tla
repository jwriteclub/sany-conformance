-------------------- MODULE test202 ----------------
\* Test of recursion

EXTENDS Naturals, TLC

RECURSIVE g(_)
f(a) == IF a = 0 THEN 1 ELSE a *'''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''-+-> []x  \* ERROR
f == []x -+-> x'  \* E \A i \in {}, j \in x' : []x   \* ERROR

mm == \E i \in {}, j \in x : []x
nn == \E i \in {}, j \in x' : []x   \* ERROR

m2 == \A i \in x' : x'
======================at] == IF n = 0 THEN 1 ELSE t
