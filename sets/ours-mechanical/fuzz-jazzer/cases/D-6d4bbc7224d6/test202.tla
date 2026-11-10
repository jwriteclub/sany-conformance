-------------------- MODULE test202 ----------------
\* a * g(a-1)
g(a) == f(a) 
       /\ g(4) = 24

RECURSIVE fact

ASSUME /\ fact1[3] = 6
       /\ fact[4] = 24
       /\ \A i \in 0..5 : fact[i] = g==================================

