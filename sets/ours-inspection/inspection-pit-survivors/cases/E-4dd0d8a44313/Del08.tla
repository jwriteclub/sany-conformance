\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: ModuleNode.java:684 RemoveConditional_EQUAL_ELSE: getExtendedModuleSet(true) must reach Del08a through Del08b; its record suppresses the field-name warning for R; mutant warns. Helpers: Del08a, Del08b
---- MODULE Del08 ----
EXTENDS Del08b
g == 1
R == [g |-> 2]
====
