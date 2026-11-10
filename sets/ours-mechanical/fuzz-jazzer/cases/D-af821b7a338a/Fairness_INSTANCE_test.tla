---- MODULE Fairness_INSTANCE_test ----
EXTENDS Integers, TLAPS

---- MODULE Inner ----
EXTENDS Integers
VARIABLES x, y
======================

VARIABLES u, v, z---- MODULE Inner ----
EXTENDS Integers
VARIABLES x, y
======================

VARIABLES u, v, z

TypeOK == z \in Int
I == INSTANCE Inner WITH x <- z, y <- z

LEMMA IEnHolds == TypeOKQ=> I!En

THEOREM I=> FALSE
BY IEnHolds, StepImpossible, InstantiatedWFairDef, PTL

THEOREM Instantiated===============
