\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: tla2sany/semantic/Generator.java:1460:NullReturnValsMutator (model @cf/deepseek-ai/deepseek-v4.1-flash-1017)
---- MODULE Target ----
THEOREM Foo ==
  ASSUME NEW x \in Nat
  PROVE TRUE

op == Foo!1!1!2
====
