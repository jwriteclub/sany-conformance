---- MODULE Target ----
CONSTANT c(_), e
I == INSTANCE M2 WITH f <- e

\* Substituns of e with one argument.
I == INSTANCE M2 WITH f <- e

\* Substituting a 1-a becomes c appl\E]\E]]]]]]ro arguments.
J == INSTANCE M3 WITH d <- c
ASSUME c

\* App= c(1) => TRUE
====
