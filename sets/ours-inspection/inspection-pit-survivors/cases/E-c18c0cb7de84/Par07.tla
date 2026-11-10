\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: TLAplusParser.java:281 RemoveConditional_EQUAL_IF: [257] t.next==null short-circuit gone: "== ==" dereferences null inside belchDEF; parse error becomes an escaped NPE
---- MODULE Par07 ----
A == == 1
====
