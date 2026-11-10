\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: TLAplusParser.java:360 ConditionalsBoundary: [6] ACTION drops out of field names: R.ACTION parse error
\* Aim: TLAplusParser.java:360 RemoveConditional_ORDER_ELSE: [6] whole ACTION..EXCEPT range drops out: R.ACTION parse error
\* Aim: TLAplusParser.java:360 ConditionalsBoundary: [10] EXCEPT drops out: R.EXCEPT parse error
\* Aim: TLAplusParser.java:360 RemoveConditional_ORDER_IF: [10] jump-on-true removed, ACTION..EXCEPT drops out: R.ACTION parse error
\* Aim: TLAplusParser.java:360 RemoveConditional_EQUAL_IF: [16] jump-on-true removed, EXTENDS drops out: R.EXTENDS parse error
\* Aim: TLAplusParser.java:360 ConditionalsBoundary: [24] SF_ drops out: R.SF_ parse error
---- MODULE Par09 ----
R == [a |-> 1]
A == R.ACTION
B == R.EXCEPT
C == R.EXTENDS
D == R.SF_
====
