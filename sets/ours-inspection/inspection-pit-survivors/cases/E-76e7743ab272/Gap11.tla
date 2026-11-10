\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: Generator.java:1847 RemoveConditional_EQUAL_ELSE: label naming an ASSUME/PROVE accepted as an expression at B
---- MODULE Gap11 ----
THEOREM T == ASSUME Lap :: ASSUME TRUE PROVE TRUE PROVE TRUE
B == T!Lap
====
