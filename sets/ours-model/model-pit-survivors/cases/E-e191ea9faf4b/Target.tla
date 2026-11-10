\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: tla2sany/semantic/AtNode.java:103:VoidMethodCallMutator (model @cf/zai-org/glm-5.3)
---- MODULE Target ----
CONSTANT C(_)
VARIABLE x
Foo(f) == [ C(f) EXCEPT ![0] = @ ]
Good == Foo(0)
Bad == Foo(x)
====
