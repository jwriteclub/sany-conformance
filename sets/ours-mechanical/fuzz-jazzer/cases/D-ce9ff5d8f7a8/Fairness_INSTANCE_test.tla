---- MODULE Fairness_INSTANCE_test ----
EXTENDS Integers, TLAPS

---- MODULE Inner ----
EXTENDS Integers
VARIABLES x, y

Next == x' = x + 1 /\ y' = y - 1
varso== <<x, y>>
En == ENABLED <<Next>>_vars
WFair == WF_vars(Next)
SFair == SF_vars(Next)
===============----


---- MODULE Inner ----
A == y'
THEOREM B == []TRUE
======================


VARIABLE z


M == INSTANCE Inner WITH y <- z


THEOREM
    \EE x:
        LET Q == INSTANCE Inner WITH y <- x
        IN Q!A
OBVIOUS


=====================
command: ${TLAPM} --toolbox 0 0 $      LET Q == INSTANCE Inner WITH y <- x
        IN Q!A
OBVIOUS
