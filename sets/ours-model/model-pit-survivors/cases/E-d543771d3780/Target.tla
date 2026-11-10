\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: tla2sany/semantic/Generator.java:1796:NullReturnValsMutator (model @cf/deepseek-ai/deepseek-v4.1-flash-1017)
---- MODULE Target ----
Inc(x) == x

Apply(Op(_), x) == Op(x)

Foo == Apply(Inc, 1)

expr == Foo!1
====
