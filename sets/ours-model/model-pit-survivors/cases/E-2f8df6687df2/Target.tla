\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: tla2sany/semantic/Generator.java:1499:VoidMethodCallMutator (model @cf/deepseek-ai/deepseek-v4.1-flash-1017)
---- MODULE Target ----
G(Op(_)) == TRUE
F == G(LAMBDA x: x)
A == F!1!(2)
====
