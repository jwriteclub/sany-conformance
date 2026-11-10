\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: Generator.java:2845 RemoveConditional_EQUAL_IF: function named like a module instance reports 'wrong LET/IN level' instead of 'already defined'
---- MODULE Gap18 ----
---- MODULE I ----
====
M == INSTANCE I
M[x \in {1}] == x
====
