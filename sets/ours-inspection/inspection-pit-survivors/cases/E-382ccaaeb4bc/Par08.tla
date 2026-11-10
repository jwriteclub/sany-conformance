\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: TLAplusParser.java:329 ConditionalsBoundary: [10] ^# no longer postfix: DEFBREAK lands before ^#, not before B; accepted spec turns into a parse error
---- MODULE Par08 ----
a ^# == a
A == 1 ^#
B == 2
====
