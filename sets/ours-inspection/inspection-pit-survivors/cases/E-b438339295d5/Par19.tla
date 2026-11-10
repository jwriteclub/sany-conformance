\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: TLAplusParser.java:445 VoidMethodCall: [68] no recursion: ) under the bullet two levels down loses the not-properly-indented error
---- MODULE Par19 ----
A == /\ 1 + (2
  )
     /\ TRUE
====
