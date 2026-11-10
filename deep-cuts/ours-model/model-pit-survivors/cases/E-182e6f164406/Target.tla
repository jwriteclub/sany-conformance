\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: tla2sany/semantic/AssumeProveNode.java:252:VoidMethodCallMutator (model @cf/zai-org/glm-5.3-flash)
---- MODULE Target ----
CONSTANT c

THEOREM T ==
  ASSUME c' # c
  PROVE TRUE
====
