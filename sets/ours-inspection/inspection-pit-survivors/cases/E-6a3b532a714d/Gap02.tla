\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: Generator.java:2923 RemoveConditional_EQUAL_ELSE: function def loses labels, so f!L(1) gets 'Cannot find label'
\* Aim: Generator.java:2924 VoidMethodCall: same, setLabels on a function definition removed
---- MODULE Gap02 ----
f[x \in {1}] == L(x) :: x
A == f!L(1)
====
