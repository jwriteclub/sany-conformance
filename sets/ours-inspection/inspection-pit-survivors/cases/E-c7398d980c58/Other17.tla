\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: SANY.java:343 RemoveConditional_EQUAL_ELSE: validation adds a non-aborting error (algorithm marker not in a comment); baseline stops as a parse failure, mutant goes on to semantic analysis
---- MODULE Other17 ----
x == "--algorithm foo"
\* BEGIN TRANSLATION
y == x
\* END TRANSLATION
====
