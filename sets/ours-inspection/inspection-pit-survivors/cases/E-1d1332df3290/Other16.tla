\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: SpecObj.java:877 RemoveConditional_ORDER_ELSE: inner J's INSTANCE I is not resolved in the extended unit; goes to the file system and aborts. Helper: Other16a
\* Aim: SpecObj.java:880 VoidMethodCall: same
---- MODULE Other16 ----
EXTENDS Other16a
---- MODULE J ----
K == INSTANCE I
v == K!i
====
====
