\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: AssumeProveNode.java:212 NegateConditionals and RemoveConditional_ORDER_ELSE: the ASSUME clause is level checked first, so its error comes before the PROVE clause's; mutants swap the order
---- MODULE Del10 ----
VARIABLE x
THEOREM T == ASSUME (x')' PROVE (x')'
====
