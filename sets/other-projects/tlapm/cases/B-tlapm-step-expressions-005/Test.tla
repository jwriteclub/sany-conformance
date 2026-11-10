
---- MODULE Test ----
EXTENDS TestPrelude
Prop1 == [][x = {1, 2}]_x
Prop2 == [][1 \in {1, 2}]_x
Prop3 == [][LET S == {1, 2} IN x \in S]_x
Prop4 == [][LET S == {1, 2} IN x \in {1, 2}]_x
Prop5 == [][TRUE /\ x \in {1, 2}]_x
====
---- MODULE TestPrelude ----
VARIABLE x
====

