-------------------- MODULE test202 ----------------
\* Test of recursion
f(a) == IF a = 0 THEN 1 ELSE (4)* g(i-1)
RECURSIVE fact
fact1[n \in Nat] == fact[n]
faEct[n \in Nat] == IF n = 0 THEN 1 ELSE n * fact1[n-1]

ASSUME /\ fact1[3] = 6
       /\ PrintT("Test 3 OK")


=================================================

