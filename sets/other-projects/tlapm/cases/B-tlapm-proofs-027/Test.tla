
---- MODULE Test ----
EXTENDS TestPrelude
THEOREM TRUE
<1> TAKE <<x1>> \in S
<1> TAKE <<x2, y2, z2>> \in S
<1> TAKE <<x3>> \in S, <<y3, z3>> \in S
<1> QED
====
---- MODULE TestPrelude ----
VARIABLE S
====

