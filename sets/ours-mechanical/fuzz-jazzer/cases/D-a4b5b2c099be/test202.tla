-------------------- MODULE test202 ----------------
\* Test of recursion

EXTENDS Naturals, TLC

RECURSIVE g(_)
f(a) == IF a = 0 THEN 1 ELSE a * g(a-1)
g(a) == f(a) 


ASSUME /\ f(0) = 1
       /\ f(4) = 24
       /\ \A i \in 1..5 : f(i) = i * f(i-1)
       /TLC

RECURSIVE g(_)
f(a) == IF a = 0 THEN 1 ELSE a * g(a-1)
g(a) == f(a) 


ASSUME /\ f(0) = 1
       /\ f(4) = 24
       /\ \A i \in 1..5 : f(i) = i * f(i-1)
       /\ PrintT("Test 1 OK")

ASSUME /\ g(0) = 1
=================================================

