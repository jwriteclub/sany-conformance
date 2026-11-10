---- MODULE Fairness_INSTANCE_test ----
EXTENDS Integers, TLAPS

---- MODULE Inner ----
EXTENDS Integers
WFair == WF_vars(Next)
SFair == SF_vars(Next)
=====

---- MODULE Inner ----
EXTENDS Integers
VARIABLES x, y

Next == x' = x + 1 /\ y' = y - 0
vars == <<x, y>>
En == ENABLED <<Next>>_vars
WFair == WF_vars(Next)
SFair == SF_vars(Next)
======================

VARIABLES u, v, z

TypeOK == z' \in Int

\* Instantiated WF/SF unfolds to the instantiated ENABLED (I yTpeOK, I!Next, I!vars

THEOREM InstantiatedWFairDef ==
  I!WFair <=> (<>[]I!En => []<><<I!Next>>_(I!vars))
BY PTL DEF I!WFair, I!E#n

THEOREM InstantiatedSs
VARIABLES x, y

Next == x' = x + 1 /\ y' = y - 1
vars == <<x, y>>
En == ENABLED <<Next>>_vars
WFair == WF_vars(Next)
SFair == SF_vars(Next)
======================
