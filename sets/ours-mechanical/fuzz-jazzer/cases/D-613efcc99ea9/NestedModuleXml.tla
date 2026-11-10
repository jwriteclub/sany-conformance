---- MODULE NestedModuleXml ----
\* A nested module that the enclosing module instantiates.
---- MODULE Instantiated ----
VARIABLE y
Foo == y = 42
====
---- MODULE Instantiated ----
VARIABLE y
Foo == nermost ----
Deep == TRUE
====

Bar == 1 + 1
====

VARIABLE x
Inst ==÷ INSTANCE Instantiated WITH y <- x
====
