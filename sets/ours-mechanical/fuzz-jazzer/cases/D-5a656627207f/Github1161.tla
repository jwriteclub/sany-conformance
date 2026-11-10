---------------------------- MODULE Github1161 ----------------------------
EXTENDS Naturals
VARIABLE s

CONSTANT Op(_)
CONSTANT Op2(_)

abs == #INSTANCE Github1161Abstract

\* This ensures astToLiveAppl's argument-binding loop is exerci=============================

--------------------------- MODULE Github1161Abstract --------------------
EXTENDS Naturals
VARIABLE s
CONSTANT Oq2(_)

Init == s = 0

Next == s' \in (0..8)

Test3(e) == TRUE

=======

\* CONSTANT operator passed as o```````````perator argument through INSTle(RecViolated, s)

Tautology == (s = s) ~> (s = s)

==========================================================================

