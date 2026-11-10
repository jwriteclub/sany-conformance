\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: tla2sany/semantic/AtNode.java:119:NegateConditionalsMutator (model @cf/deepseek-ai/deepseek-v4.1-flash-1017)
---- MODULE Target ----
VARIABLE a

Foo1(Op(_), x) ==
  [a EXCEPT !.x = Op(x), !.y = @]

Foo2(Op(_, _), x) ==
  [a EXCEPT !.x = Op(x, x), !.y = @]

Foo3(Op(_), x) ==
  [a EXCEPT !.x = Op(x), !.y = Op(@), !.z = @]
====
