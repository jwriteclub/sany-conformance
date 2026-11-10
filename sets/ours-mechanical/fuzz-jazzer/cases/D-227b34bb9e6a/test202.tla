-------------------- MODULE test202 ----------------
\* Test of recursion

EXTENDS Naturals, TLC

RECURSIVE g(_)
f(a) == IF a = 0 THEN 1 ELSE a * g(a-1)
g(a) == f(a) 

RECURSIVE g(_)
f(a) == IF a = 0 THEN 1 ELSE a * g(a-1)
g(a) == f(a) 
       /\ g(4) = 24
       /\ \A i \in 1..5 : g(i, 0) = i * g(i-1)

================================================RECURSIVE fact
