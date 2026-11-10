\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: TLAplusParser.java:4631 RemoveConditional_EQUAL_IF: [166] column check forced true: LAMBDA left of the bullet is consumed; error changes to the indentation message
---- MODULE Par23 ----
F(a, f(_)) == f(a)
A ==   /\ F(1,
  LAMBDA x : x)
====
