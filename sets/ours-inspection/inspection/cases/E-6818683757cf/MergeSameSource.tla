---- MODULE MergeSameSource ----
\* Aim: Context.java:410, the side where two distinct definitions of one name are the same definition.
\* A and B each define J by instantiating the parameter-free P, so each gets its own J!D node
\* (Generator.java:4916). Extending both meets two nodes named J!D of one class. They share a
\* source and P has no parameters, so sameOriginallyDefinedInModule holds and there is no warning
\* for J!D. J itself has two sources, so SANY warns about J alone, and accepts.
EXTENDS A, B
====
