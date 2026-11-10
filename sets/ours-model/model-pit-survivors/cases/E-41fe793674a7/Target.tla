\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: tla2sany/semantic/Generator.java:1258:NullReturnValsMutator (model @cf/deepseek-ai/deepseek-v4.1-flash-1017)
---- MODULE Target ----
VARIABLE f

Foo == [f EXCEPT ![1] = 2]
Bar == Foo!99
====
