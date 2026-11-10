---- MODULE DiamondInst ----
EXTENDS A, B
X == Foo
====
(*@FILE A*)
---- MODULE A ----
INSTANCE Base
====
(*@FILE B*)
---- MODULE B ----
INSTANCE Base
====
(*@FILE Base*)
---- MODULE Base ----
Foo == 1
====
