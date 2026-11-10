
---- MODULE Test ----
EXTENDS TestPrelude
THEOREM TRUE
<1> PICK <<x1>> \in S : TRUE
<1> PICK <<x2, y2, z2>> \in S : TRUE
<1> PICK <<x3>> \in S, <<y3, z3>> \in S : TRUE
<1> QED
====
---- MODULE TestPrelude ----
VARIABLE S
====

