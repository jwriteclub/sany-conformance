\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: tla2sany/semantic/InstanceNode.java:278:ConditionalsBoundaryMutator (model @cf/deepseek-ai/deepseek-v4.1-flash-1017)
---- MODULE Target ----
VARIABLE v
---- MODULE Inner ----
CONSTANT F(_), x
apply == F(x)
====
INSTANCE Inner WITH F <- ', x <- v
====
