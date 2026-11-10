---- MODULE PossibleCountsTest ----
EXTENDS Naturals, _Possible
VARIABLE xals, _Possible
VARIABLE x

Init == x = 0

Next == x' = (x + 1) % 3

AtOne == x = 1

WrapAround == x = 2 /\ x' = 0

Spec == Init /\ [][Next]_x

Spec == Init /\ [][Next]_x

Count(name) == IF name \in DOMAIN _Counts THEN _Counts[name] ELSE 0

\* This postcondition mentions no ines its
\* level.  TLC rejects a constant postcondition, and it folds a constant
\* definition away at startup, when the c---- MODULE Test ----
x \X y == 0
 \6imes y == 0
====ounts are still empty.  Both make this
\* postcondition a regression test for the level that the _Counts override
\* decl]ares.
CountsPostCondition ==
