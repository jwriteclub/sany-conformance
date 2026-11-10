\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: TLAplusParser.java:483 RemoveConditional_EQUAL_ELSE: [108] non-identifier after comma not rejected: {x, 2 \in S} read as a bound; accepted spec turns into a parse error
---- MODULE Par13 ----
x == 1
A == {x, 2 \in {}}
====
