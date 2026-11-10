---- MODULE InstChainConst ----
VARIABLE v
M == INSTANCE Mid WITH k <- v
====
(*@FILE Mid*)
---- MODULE Mid ----
CONSTANT k
J == INSTANCE Base
Op == J!Foo
====
(*@FILE Base*)
---- MODULE Base ----
Foo == 1
====
