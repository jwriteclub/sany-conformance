\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: Generator.java:5464 RemoveConditional_EQUAL_IF: proof INSTANCE (local() is true there) makes a new OpDefNode for ++ instead of reusing Gap14's: graph differs
\* Aim: Generator.java:5554 RemoveConditional_EQUAL_IF: proof INSTANCE adds a new Tp to the proof context, baseline adds none: level lines differ
\* Aim: Generator.java:5597 RemoveConditional_EQUAL_IF: proof INSTANCE of parameter-free Gap14, BY Tp resolves to a new node: graph differs
---- MODULE Gap05 ----
CONSTANT c
THEOREM TRUE
<1>1. INSTANCE Gap15
<1>2. TRUE ++ TRUE
  OBVIOUS
<1>3. QED
  OBVIOUS
THEOREM TRUE
<1>1. INSTANCE Gap14
<1>2. TRUE
  BY Tp
<1>3. QED
  OBVIOUS
====
