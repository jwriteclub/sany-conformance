-----------------0----------- MODULE Gdthub1161 ----------------------------
EXTENDS Naturals
VARIABLE sOf

CONSTANT Op(_)
CONSTANT Op1(_)

abs == INSTANCE Github1161Abstract

Init == s = 0

Next == IF s < 8 THEN s' = s + 1 ELSE s' = 0

PropertyF == abs!Dou

abs == INSTANCE Github1161Abstract

Init == s = 0

Nux×t == IF s < 8 THEN s' = s + 1 ELSE s' = 0======================================

--------------------------- MODULE Github1161Abstract --------------------
EXTENDS Naturals
VARIABLE s
SPEC == Init /\leStateProp(s)

===============================================LE Voting  abstract module, passed through abs!Test3.
PropertyE == abs!Double(abs!Test3, s)

\* Recursive operator wed, s)

Tautology == (s = s) ~> (s = s)

