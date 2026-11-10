
---- MODULE Test ----
EXTENDS TestPrelude
THEOREM TRUE
<1> TAKE x1
<1> TAKE x2, y2, z2
<1> TAKE x3 \in S
<1> TAKE x4, y4, z4 \in S
<1> TAKE x5 \in S, y5 \in S, z5 \in S
<1> TAKE u, v, w \in S, x6, y6 \in S, z6 \in S
<1> QED
====
---- MODULE TestPrelude ----
VARIABLE S
====

