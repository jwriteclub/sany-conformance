\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: tla2sany/semantic/AtNode.java:111:NegateConditionalsMutator (model @cf/deepseek-ai/deepseek-v4.1-flash-1017)
---- MODULE Target ----
VARIABLE a, x

Foo(Op(_)) ==
  [a EXCEPT !.p = Op(x), !.q = @]

Bar(y) == 1

r == Foo(Bar)
====
