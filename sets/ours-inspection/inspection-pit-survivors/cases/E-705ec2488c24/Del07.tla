\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: ModuleNode.java:648 RemoveConditional_ORDER_ELSE and :650 VoidMethodCall: theorem TN of N is copied into M; M is therefore not constant and INSTANCE M WITH c <- x is an error; mutants accept it
---- MODULE Del07 ----
VARIABLE x
---- MODULE N ----
THEOREM TN == []TRUE
====
---- MODULE M ----
EXTENDS N
CONSTANT c
====
J == INSTANCE M WITH c <- x
====
