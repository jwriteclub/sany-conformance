\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: Generator.java:2835 VoidMethodCall: function defined inside a recursive section loses recursiveSection in the XML
\* Aim: Generator.java:5513 VoidMethodCall: unnamed INSTANCE inside a recursive section, Op loses recursiveSection in the XML
---- MODULE Gap06 ----
---- MODULE P ----
Op == TRUE
====
RECURSIVE f(_)
g[x \in {1}] == x
INSTANCE P
f(x) == IF x = 1 THEN g[x] ELSE f(x)
====
