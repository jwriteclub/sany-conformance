---- MODULE Target ----
EXTENDS Naturals, Sequences, Integers, FiniteSets

CONSTANT S, T
W == [x' = x]
X == x \in {1, 2} \ {2}
Y == (A + B) * 2
Z == -A

ASSUME S # {}
Inst == INSTANCE Naturals WITH N <- S
====
