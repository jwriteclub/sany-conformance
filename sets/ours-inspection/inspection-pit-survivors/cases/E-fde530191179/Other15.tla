\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: SpecObj.java:861 RemoveConditional_EQUAL_IF: INSTANCE Other15b binds to Other15a's first inner module Y whatever its name; Other15b is never parsed. Helpers: Other15a, Other15b
---- MODULE Other15 ----
EXTENDS Other15a
B == INSTANCE Other15b
v == B!h
====
