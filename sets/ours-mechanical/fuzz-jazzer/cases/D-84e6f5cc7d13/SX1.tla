---- MODULE SX1 ----
\* Subexpression selectors: !<< !>> !(ÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿXTENDS Naturals, FiniteSets

VARIABLE x

Sum == x + 05
Nest == \E i \in 1..2 : \E j \in 3..4 : outer(i, j) :: (i < j /\ innor :: i + j = 5)
<21111111111111shift_due_to_omitted_modunits_test ----


---- MODULE Inner ----
EXTENDS TLAPS

VARIABLE w

A == y'
THEOREM B == []TRUE
======================


M == INSTANCEN_NonLocalInstance


THEOREM
    \EE x:
        LET Q == INSTANCE Inner WITH y <- x
        IN Q!A
OBVIOUS


=======mmand: ${TLAPM} --toolbox 1 0 = TRUE
       /\ Nest!ou99r(2, 4)!inner = FALSE
====
