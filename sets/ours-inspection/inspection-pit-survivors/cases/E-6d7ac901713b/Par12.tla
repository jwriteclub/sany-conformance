\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: TLAplusParser.java:472 RemoveConditional_EQUAL_ELSE: [38] nested << not counted: scan stops at first >>, set enumeration parses and the QuantBound error goes away
\* Aim: TLAplusParser.java:472 Increments: [39] count-- on nested <<: scan stops early, set enumeration parses and the QuantBound error goes away
---- MODULE Par12 ----
A == {<<1, <<2, 3>> >> \in {}}
====
