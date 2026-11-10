---- MODULE DiamondParam ----
EXTENDS A, B
X == Foo
====
(*@FILE A*)
---- MODULE A ----
EXTENDS K0
INSTANCE Base
====
(*@FILE B*)
---- MODULE B ----
EXTENDS K0
INSTANCE Base
====
(*@FILE K0*)
---- MODULE K0 ----
CONSTANT k
====
(*@FILE Base*)
---- MODULE Base ----
CONSTANT k
Foo == k
====
