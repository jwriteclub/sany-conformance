\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: Context.java:400 RemoveConditional_EQUAL_IF: Both extends Left and Right over a base that declares CONSTANT and VARIABLE; the same symbols arrive twice and must not warn; mutant warns
\* Aim: Context.java:342 and :344 RemoveConditional_EQUAL_ELSE, ModuleNode.java:920 ORDER_ELSE: inner modules Base..Join must be found (XML differs: level data of their definitions)
---- MODULE Del05 ----
EXTENDS Naturals
---- MODULE Base ----
CONSTANT p
VARIABLE w
Foo == 1
====
---- MODULE Left ----
EXTENDS Base
====
---- MODULE Right ----
EXTENDS Base
====
---- MODULE Both ----
EXTENDS Left, Right
====
---- MODULE Plain ----
Bar == 2
====
---- MODULE Mid ----
INSTANCE Plain
====
---- MODULE Join ----
EXTENDS Plain, Mid
====
INSTANCE Naturals
====
