\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: tla2sany/semantic/OpApplNode.java:551:ConditionalsBoundaryMutator (model @cf/zai-org/glm-5.3)
---- MODULE Target ----
VARIABLE v
H(h(_)) == h(v)
D(y) == y'
B == H(D)
====
