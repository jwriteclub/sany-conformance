\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: Generator.java:3913 VoidMethodCall: bounded quantifier's params stay on the stack, so later label L1 gets 'must contain formal parameter'
\* Aim: Generator.java:3976 VoidMethodCall: same for unbounded quantifier, label L2
\* Aim: Generator.java:3854 VoidMethodCall: same for unbounded CHOOSE, label L3
\* Aim: Generator.java:3885 VoidMethodCall: same for bounded CHOOSE, label L4
\* Aim: Generator.java:4033 VoidMethodCall: same for subset-of, label L5
\* Aim: Generator.java:4056 VoidMethodCall: same for set-of-all, label L6
\* Aim: Generator.java:4622 VoidMethodCall: same for LAMBDA, label L7
---- MODULE Gap01 ----
G(F(_)) == F(1)
A1 == /\ \A x \in {1} : x = 1
      /\ L1 :: TRUE
A2 == /\ \A x : x = 1
      /\ L2 :: TRUE
A3 == /\ (CHOOSE x : x = 1) = 1
      /\ L3 :: TRUE
A4 == /\ (CHOOSE x \in {1} : x = 1) = 1
      /\ L4 :: TRUE
A5 == /\ {x \in {1} : x = 1} = {1}
      /\ L5 :: TRUE
A6 == /\ {x : x \in {1}} = {1}
      /\ L6 :: TRUE
A7 == /\ G(LAMBDA x : x) = 1
      /\ L7 :: TRUE
====
