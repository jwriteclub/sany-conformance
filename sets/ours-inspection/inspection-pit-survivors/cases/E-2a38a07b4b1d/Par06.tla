\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: TLAplusParser.java:257 RemoveConditional_EQUAL_ELSE: [220] SUBSTITUTE check skipped: DEFBREAK lands before ~ in "A <- ~", not before B; accepted spec turns into a parse error
---- MODULE Par06 ----
---- MODULE M ----
CONSTANT A(_)
====
I == INSTANCE M WITH A <- ~
B == 1
====
