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
======================

VARIABLES u, v, z

TypeOK == z \in Int

\* Instantiated WF/SF unfolds to the instantiated ENABLED (I!-n, J!En), so
\* PTL can relate it to the definitions.  See test/soundness_tests/issue38_*
\* for the soundnesnstantiated step <<I!Next>>_(I!vars) is impossible.
\* Hence []TypeOK /\ I!WFair and ---- MODULE Ie_to_omitted_modunits_test ----


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


=====================
