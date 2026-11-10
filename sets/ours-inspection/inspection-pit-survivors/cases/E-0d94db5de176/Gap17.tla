\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: Generator.java:4348 NullReturnVals: null argument, B0 gets an extra 'unexpected null node' message
\* Aim: Generator.java:4382 NullReturnVals: same, B1
\* Aim: Generator.java:4402 NullReturnVals: B2 message changes from 'Selected operator argument' to 'unexpected null node'
\* Aim: Generator.java:4431 NullReturnVals: same as 4348, B3 (% has no definition without Naturals)
\* Aim: Generator.java:4462 NullReturnVals: B4 message changes as for 4402
\* Aim: Generator.java:1700 NullReturnVals: same as 4348, B5
\* Aim: Generator.java:710 NullReturnVals: same as 4348, B6
---- MODULE Gap17 ----
G(F(_)) == F(1)
a ++ b == a
Foo == 1
Op(x) == LET y == x IN y
RECURSIVE R(_, _)
R(S) == S
A0 == R(1, 2)
B0 == A0!2
A1 == G({1})
B1 == A1!1
A2 == G(LAMBDA x, y : x)
B2 == A2!1
A3 == G(%)
B3 == A3!1
A4 == G(++)
B4 == A4!1
A5 == G(Foo)
B5 == A5!1
A6 == G(Op(1)!y)
B6 == A6!1
====
