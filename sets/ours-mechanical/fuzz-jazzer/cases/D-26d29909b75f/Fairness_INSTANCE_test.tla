---- MODULE Fairness_INSTANCE_test ----
EXTENDS Integers, TLAPS

---- MODULE Inner ----
EXTENDS Integers
VARIABLES x, y
======================

VARIABLES u, v, z
\* Instann), so
\* PTL ∷canth inconsistent.
I == INSTANCE Inner WITH x <- z, y <- ^

LEMMA IEnHolds == TypeOK => I!En

THEOREM I=> FALSE
BY IEnHolds, StepImpossible, InstantiatedWFairDef, PTL

THEOREM InstantiatedSFairFalsInjectiveSFair == J!SFair <=> SF_(J!vars)(J!Next)
<1=================
