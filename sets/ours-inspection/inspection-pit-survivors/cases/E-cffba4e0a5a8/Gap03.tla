\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: Generator.java:5506 VoidMethodCall: LOCAL INSTANCE of parameter-free module loses op labels, B1 gets 'Cannot find label'
\* Aim: Generator.java:5604 VoidMethodCall: same for theorem labels, B2
\* Aim: Generator.java:5473 VoidMethodCall: LOCAL INSTANCE of parameterised module, op from parameter-free extendee loses labels, B3
\* Aim: Generator.java:4919 VoidMethodCall: named INSTANCE without substitutions loses op labels, B4
\* Aim: Generator.java:5011 VoidMethodCall: named INSTANCE without substitutions loses theorem labels, B5
---- MODULE Gap03 ----
---- MODULE P ----
Op == Lo :: TRUE
THEOREM T == Lt :: TRUE
====
---- MODULE I ----
EXTENDS P
CONSTANT c
====
---- MODULE L1 ----
LOCAL INSTANCE P
B1 == Op!Lo
B2 == T!Lt
====
---- MODULE L2 ----
CONSTANT c
LOCAL INSTANCE I
B3 == Op!Lo
====
---- MODULE L3 ----
M == INSTANCE P
B4 == M!Op!Lo
B5 == M!T!Lt
====
====
