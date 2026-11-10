\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: tla2sany/semantic/Generator.java:1210:NullReturnValsMutator (model @cf/deepseek-ai/deepseek-v4.1-flash-1017)
---- MODULE Target ----
Foo == CASE TRUE -> 1 [] OTHER -> 2
Bar == Foo!1
====
