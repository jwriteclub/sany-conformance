\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: TLAplusParser.java:417 RemoveConditional_EQUAL_IF: [12] postfix body treated as non-infix: UNCHANGED L:: x' loses the removing-label parse error
---- MODULE Par18 ----
VARIABLE x
A == UNCHANGED L:: x'
====
