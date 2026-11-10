\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: Generator.java:5039 VoidMethodCall: LET INSTANCE not in LetInNode.insts, so the level error on c <- v goes unreported; helper Gap20
---- MODULE Gap19 ----
VARIABLE v, w
B == LET M == INSTANCE Gap20 WITH c <- v IN M!A
====
