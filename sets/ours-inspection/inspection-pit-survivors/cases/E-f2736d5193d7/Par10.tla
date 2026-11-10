\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: TLAplusParser.java:360 ConditionalsBoundary: [34] THEN drops out: R.THEN parse error
\* Aim: TLAplusParser.java:360 ConditionalsBoundary: [38] WITH drops out: R.WITH parse error
\* Aim: TLAplusParser.java:360 RemoveConditional_EQUAL_IF: [44] jump-on-true removed, _ drops out: R._ parse error
---- MODULE Par10 ----
R == [a |-> 1]
A == R.THEN
B == R.WITH
C == R._
====
