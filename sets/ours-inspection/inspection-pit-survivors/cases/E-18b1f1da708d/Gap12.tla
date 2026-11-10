\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: Generator.java:3632 RemoveConditional_EQUAL_IF: nested ASSUME/PROVE gets a goal, so T!1!2 inside a NEW's scope is accepted instead of rejected
---- MODULE Gap12 ----
CONSTANT S
THEOREM T == ASSUME ASSUME NEW x, x \in S PROVE x \in S PROVE TRUE
B == T!1!2
====
