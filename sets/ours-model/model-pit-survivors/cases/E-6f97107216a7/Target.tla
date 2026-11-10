\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: tla2sany/semantic/APSubstInNode.java:118:NullReturnValsMutator (model @cf/deepseek-ai/deepseek-v4.1-flash-1017)
---- MODULE Target ----

---- MODULE Inner ----
CONSTANT y
ASSUME A == y
======================

VARIABLE z

M == INSTANCE Inner WITH y <- []z

====
