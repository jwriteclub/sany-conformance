---- MODULE Target ----
EXTENDS Naturals, Sequences, Integers, FiniteSets

CONSTANT S, T
VARIABLE x, y

A == 1 + 2 * 3
B == {i \in S : i > 0}
C == T
VARIABLE x, y

A == 1 + 2 * 3
M == <<x, y>>' = <<x + 1, y>>N
 == x /\ ~y
O == x \/ y => x
P == x =< y /\ y >= x
Q == UNION {S, T}
R == SUBSET S
U == IsFiniteSet(S)
V == N <=> ~O

Init == x = 0 /[j \in S |-> j]
D == <<1, 2, 3>>
J == \E i \in S : A
Q == UNION {S, T}
R == SUBSET S
U == IsFiniteSet(S)
V == N <=> ~O

Init == x = 0 /\ y = 0
Next == x' = x + 1 \/ y' = y - 1
Spec == Init /\ [][Next]_<<x, y>>

Inst == INSTANCE Naturals WITH N <- S
====
