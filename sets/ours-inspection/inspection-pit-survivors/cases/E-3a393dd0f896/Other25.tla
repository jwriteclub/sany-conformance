\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: Subst.java:189 RemoveConditional_EQUAL_ELSE: SubstIn loses arg constraint (H, 0) = 1: alc differs
\* Aim: Subst.java:192 RemoveConditional_EQUAL_ELSE: same
\* Aim: Subst.java:195 RemoveConditional_EQUAL_ELSE: same
---- MODULE Other25 ----
CONSTANT H(_)
VARIABLE x
---- MODULE M ----
CONSTANT G(_), a
P == G(a)
====
I == INSTANCE M WITH G <- H, a <- x
Q == I!P
====
