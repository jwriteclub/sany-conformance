\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: SpecObj.java:711 RemoveConditional_ORDER_ELSE: getModulesIndirectlyExtending stays {Other10a}; Other10's INSTANCE I is not bound to Other10b's inner I and aborts on a missing file. Helpers: Other10a, Other10b
\* Aim: SpecObj.java:713 RemoveConditional_EQUAL_ELSE: same
\* Aim: SpecObj.java:714 BooleanFalseReturn: same
\* Aim: SpecObj.java:739 NegateConditionals: same
\* Aim: SpecObj.java:739 RemoveConditional_EQUAL_ELSE: same
\* Aim: SpecObj.java:745 NegateConditionals: same
\* Aim: SpecObj.java:745 RemoveConditional_ORDER_ELSE: same
\* Aim: SpecObj.java:750 NegateConditionals: same
\* Aim: SpecObj.java:750 RemoveConditional_EQUAL_ELSE: same
\* Aim: SpecObj.java:754 RemoveConditional_EQUAL_ELSE: same
\* Aim: SpecObj.java:758 VoidMethodCall: same
---- MODULE Other10 ----
EXTENDS Other10a
J == INSTANCE I
v == J!i
====
