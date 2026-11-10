\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: Linter.java:192 BooleanFalseReturn: the walk does not stop after a parameter is found, follows F's recursive call for ever: StackOverflowError
---- MODULE Other20 ----
RECURSIVE F(_)
F(n) == IF n = 0 THEN 0 ELSE F(n)
foo == 1
R == [foo |-> F(1)]
====
