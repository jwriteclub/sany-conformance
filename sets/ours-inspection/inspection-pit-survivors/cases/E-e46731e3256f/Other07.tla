\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: ParseUnit.java:450 RemoveConditional_EQUAL_ELSE: earlier peer not in context, EXTENDS A goes to the file system and aborts
\* Aim: ParseUnit.java:535 VoidMethodCall: no context at all for inner modules, same abort
\* Aim: ParseUnit.java:545 RemoveConditional_EQUAL_ELSE: inner modules not registered, same abort
\* Aim: ParseUnit.java:149 NullReturn: no parent, peers not bound, same abort
\* Aim: ParseUnit.java:467 VoidMethodCall: inner C of extended A not bound in B, EXTENDS C aborts
\* Aim: ParseUnit.java:396 RemoveConditional_ORDER_ELSE: same as 467
---- MODULE Other07 ----
---- MODULE A ----
---- MODULE C ----
c == 1
====
a == 1
====
---- MODULE B ----
EXTENDS A, C
b == <<a, c>>
====
====
