
---- MODULE Test ----
EXTENDS TestPrelude
THEOREM TRUE
<1> PICK x1 : TRUE
<1> PICK x2, y2, z2 : TRUE
<1> PICK x3 \in S : TRUE
<1> PICK x4, y4, z4 \in S : TRUE
<1> PICK x5 \in S, y5 \in S, z5 \in S : TRUE
<1> PICK u, v, w \in S, x6, y6 \in S, z6 \in S : TRUE
<1> QED
====
---- MODULE TestPrelude ----
VARIABLE S
====

