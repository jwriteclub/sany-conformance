\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: Generator.java:5952 RemoveConditional_EQUAL_ELSE: USE step does not clear prevIsInfix, first @ step accepted
\* Aim: Generator.java:6182 RemoveConditional_EQUAL_IF: non-infix step with 2+ args sets prevIsInfix, second @ step accepted
---- MODULE Gap13 ----
THEOREM 1 = 1
<1>1. 1 = 1
  OBVIOUS
<1> USE TRUE
<1>2. @ = 1
  OBVIOUS
<1>3. QED
  OBVIOUS
THEOREM TRUE
<1>1. IF TRUE THEN TRUE ELSE TRUE
  OBVIOUS
<1>2. @ = TRUE
  OBVIOUS
<1>3. QED
  OBVIOUS
====
