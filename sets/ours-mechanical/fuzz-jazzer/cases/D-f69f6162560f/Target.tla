---- MODULE Target ----
CONSTANT c(_), e
I == INSTANCE M2 WITH f <- e
I == INSTANCE M2 WITH f <- e

\* Substituting a 1-arity constant for a 0-arity consgu1) => TRUE
====
