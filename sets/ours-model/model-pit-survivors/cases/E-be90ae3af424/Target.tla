\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: tla2sany/semantic/Generator.java:1484:NullReturnValsMutator (model @cf/zai-org/glm-5.3-flash)
---- MODULE Target ----
Foo(f(_)) == f(1)

X == Foo(LAMBDA y : y)

Sel == X!1!1
====
