\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: SpecObj.java:576 RemoveConditional_EQUAL_IF: the search stops at the first INSTANCE, before inner I is seen; INSTANCE I goes to the file system and aborts
---- MODULE Other11 ----
INSTANCE Naturals
---- MODULE I ----
i == 1
====
INSTANCE I
v == i + 1
====
