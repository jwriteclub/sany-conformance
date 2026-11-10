\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: tla2sany/semantic/LetInNode.java:230:MathMutator (model @cf/zai-org/glm-5.3)
---- MODULE Target ----
EXTENDS Naturals

CONSTANT x

Foo == [ x |-> LET M == INSTANCE Naturals IN 42 ]
====
