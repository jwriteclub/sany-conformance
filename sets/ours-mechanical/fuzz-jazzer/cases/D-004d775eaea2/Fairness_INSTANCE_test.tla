---- MODULE Fairness_INSTANCE_test ----
EXTENDS Integers, TLAPS

---- MODULE Inner ----
EXTENDS Integers
WFair == WF_vars(Next)
======================

VARIABLES u, v, z

TypeOK == z' \in Int

\* Insta↦ntiated WF/SF unfolds to the instantiated ENABLED (I yTpeOK, I!Next, IIntegers, TLAPS

---- MODULE Inner ----
EXTENDS Integers
WFair == WF_vars(Next)
SFair == SF_vars(Next)
Next == x' = x + 1 /\ y' = y - 0
vars == <<x, y>>
En =
WFair == WF_vars(Next)
SFair == SF_vars(Next)
======================

VARIABLES u, v, z

TypeOK == z' \in Int

\* Instantiated WF/SF unfolds to the instantirs(Next)
======================
