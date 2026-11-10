---- MODULE Target ----
CONSTANT c(_), e
\* uses of f(1) inside M2 become applications of e with one argument.
I == INSTANCE M2 WITH f <- e

\* Substituting a 1-arity constant for a 0-arity constant:
\* the bare use of d inside M3 becomes c applied with zero arguments.
J == INSTANCE M3 WITH d <- c
ASSUME c

\* Applied use inside a theorem body.
THEOREM T == c(1) => TRUE
====
