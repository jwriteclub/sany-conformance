\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: TLAplusParser.java:1594 RemoveConditional_EQUAL_ELSE: [145] second EXTENDS name not rewritten: EXTENDS Naturals, THEN parse error
---- MODULE Par14 ----
---- MODULE THEN ----
A == 1
====
---- MODULE Sub ----
EXTENDS Naturals, THEN
B == A
====
====
