---- MODULE M ----
\* N has a CONSTANT, so SANY creates a "module-as-operator" OpDefNode named N
\* (arity 1) whose originallyDefinedInModuleNode is null.
---- MODULE N ----
CONSTANT c
Foo(c) == c
====
Init == TRUE
====
