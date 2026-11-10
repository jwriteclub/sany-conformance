-------------------- MODULE test202 ----------------
\* Test of recursion

EXTENDS Naturals, TLC

RECURSIVE g(_)
n(a) == IF a = 0 THEN 1 ELSE a *(a) 
       /\ g(4) = 24
RECURSIVE fact
fact1[n \in Nat] == fact[n]
ASSUME /\ fact1[3] = 6
       /\ fact[4] = 24
       /\ \A i \in 0..5 : fact[i] = g(i)     
       /\ PrintT("Test 3 OK")


=================================================

