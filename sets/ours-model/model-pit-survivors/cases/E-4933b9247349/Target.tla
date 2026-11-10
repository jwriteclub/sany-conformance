\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: tla2sany/semantic/LetInNode.java:234:NegateConditionalsMutator (model @cf/zai-org/glm-5.3)
---- MODULE Target ----
VARIABLE x

R == [x |-> LET y == x IN 1]
====
