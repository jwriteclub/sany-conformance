\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: ParseUnit.java:409 VoidMethodCall: E is an inner module of D, reached only by recursing from A; without it B's EXTENDS E aborts
---- MODULE Other08 ----
---- MODULE D ----
---- MODULE E ----
e == 1
====
====
---- MODULE A ----
EXTENDS D
---- MODULE F ----
f == 1
====
====
---- MODULE B ----
EXTENDS A, E
b == e
====
====
