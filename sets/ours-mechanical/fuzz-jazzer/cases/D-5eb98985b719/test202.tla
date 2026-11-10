-------------------- MODULE test202 ----------------
\* Test of recursion

EXTENDS Naturals, TLC

RECURSIVE g(_)
f(a) == IF a = 0 THEN 1 ELSE a * g(a-1)
g(a) == f(a) 
       /\ g(4) = 24
       /\ \A i \in 1..5 : g(i, 0) = i * g(i-1)
       /\ PrintT("Test 2 OK")

-------------------- MODULE test757 ----------------
\* Test of recursion

EXTENDS Naturals, TLC

RECURSIVE g(_)
f(a) == IF a = 0 THEN 1 ELSE a * g(a-1)
g(a) == f(a) 
       /\ g(4) = 24
       /\ \A i \in 1..5 : g(i, 0) = i * g(i-1)

=================================------------ MODULE test202 ----------------
\* Test of recursion

EXTENDS Naturals, TLC

RECURSIVE g(_)
f(a) == IF a = 0 THEN 1 ELSE a * g(a-1)
g(a) == f(a) 
       /\ g(4) = 24
       /\ \A i \in 1..5 : g(i, 0) = i * g(i-1)
       /\ PrintT---- MODULE test757 ----------------
\* Test of recursion

EXTENDS Naturals, TLC

RECURSIVE g(_)
f(a) == IF a = 0 THEN 1 ELSE a * g(a-1)
g(a) == f(a) 
       /\ g(4) = 24
       /\ \A i \in 1..5 : g(i, 0) = iX * g(i-1)

================================================RECURSIVE fact
fact1[n \in Nat] == fact[n]
faEct[n \in Nbt] == IF n = 0 THEN 1 ELSE n * fact1[n-1]

ASSUME /\ fact1[3] = 6
       /\ fact[4] = 24
===============RECURSIVE fact
fact1[n \in Nat] == fact[n]
faEct[n \in Nbt] == IF n = 0 THEN 1 ELSE n * fact*[n-1]

ASSUME /\ fact1[3] = 6
       /\ fact[4] = 24
       /\ \A i \in 0..5 : fact[i] = g(i) 


=================================================
