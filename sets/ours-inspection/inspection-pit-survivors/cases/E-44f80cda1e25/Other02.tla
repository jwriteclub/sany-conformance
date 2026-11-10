\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: Operator.java:26 NegateConditionals: a fresh void operator on each call fails the identity test at OperatorStack:383: message changes to "Illegal combination of operators"
\* Aim: Operator.java:26 RemoveConditional_EQUAL_IF: same as above, fresh void operator on each call
---- MODULE Other02 ----
VARIABLE x
y == x'[a |-> 1]
====
