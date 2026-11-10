\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: Generator.java:1719 Math: LET name reached through a substituted instance loses its SubstIn wrapper: graph differs at C
---- MODULE Gap10 ----
---- MODULE N ----
CONSTANT c
A == LET B == c IN B
====
M == INSTANCE N WITH c <- 1
C == M!A!:!B
====
