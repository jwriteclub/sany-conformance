\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: TLAplusParser.java:790 ConditionalsBoundary: [75] level 0 counted as no level: <1>1 no longer opens a subproof under <0>1; parse error
\* Aim: TLAplusParser.java:3095 RemoveConditional_EQUAL_IF: [203] :: test gone: ASSUME x, ASSUME y PROVE z parsed as labelled AssumeProve; parse error
\* Aim: TLAplusParser.java:3859 RemoveConditional_EQUAL_IF: [63] jump-on-true removed: <*>1 step keeps its raw image; BY <*>1 no longer resolves
\* Aim: TLAplusParser.java:4911 RemoveConditional_EQUAL_IF: [329] null heir added at second comma of {y \in S, 1, 2}: internal failure
\* Aim: TLAplusParser.java:4999 ConditionalsBoundary: [689] one-heir N_Number: {1 : x \in S} indexes past end; AIOOBE
\* Aim: TLAplusParser.java:6646 RemoveConditional_EQUAL_IF: [215] plain step ref corrected too: BY <01>1 image becomes <1>1 and resolves
---- MODULE Par25 ----
CONSTANT p, q, r, s
y == 1
S == {1}
SA == {y \in S, 1, 2}
SB == {1 : x \in S}
THEOREM T0 == TRUE
<0>1. TRUE
  <1>1. TRUE
    OBVIOUS
  <1>2. QED
    OBVIOUS
<0>2. QED
  OBVIOUS
THEOREM T1 == ASSUME p, ASSUME q PROVE r PROVE s
  OMITTED
THEOREM T2 == TRUE
<*>1 TRUE
  OBVIOUS
<*>2 QED
  BY <*>1
THEOREM T3 == TRUE
<1>1. TRUE
  OBVIOUS
<1>2. QED
  BY <01>1
====
