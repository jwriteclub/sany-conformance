\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: OperatorStack.java:131 RemoveConditional_EQUAL_IF: non-nfix op with an N_Times left operand flattens into N_Times instead of N_InfixExpr: tree differs
---- MODULE Other03 ----
A == {1}
y == A \X A \cup A
====
