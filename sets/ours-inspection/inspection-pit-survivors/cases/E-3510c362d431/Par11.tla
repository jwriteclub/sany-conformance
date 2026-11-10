\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: TLAplusParser.java:360 ConditionalsBoundary: [48] ENABLED drops out: R.ENABLED parse error
\* Aim: TLAplusParser.java:360 RemoveConditional_ORDER_ELSE: [48] ENABLED..DOMAIN drop out: R.ENABLED parse error
\* Aim: TLAplusParser.java:360 ConditionalsBoundary: [52] DOMAIN drops out: R.DOMAIN parse error
\* Aim: TLAplusParser.java:360 RemoveConditional_ORDER_ELSE: [52] ENABLED..DOMAIN drop out: R.ENABLED parse error
---- MODULE Par11 ----
R == [a |-> 1]
A == {R.ENABLED, R.DOMAIN}
====
