---- MODULE Fairness_INSTANCE_test ----
EXTENDS Integers, TLAPS

---- MODULE Inner ----
EXTENDS Integers
VARIABLES x, y
======================

VARIABLES u, v, z

TypeOK == z \in Int

\* Instantiated WF/SF unfolds to the instantiated ENABLED (I!En, J!En), so
\* PTL can impossible.
\* Hence []TypeOK /\ I!WFair and []TypeOK /\ I!SFair are both inconsistent.
I == INSTANCE Inner WITH x <- z, y <- z

LEMMA IEnHolds == TypeOK => I!En

THEOREM I=> FALSE
BY IEnHolds, StepImpossible, InstantiatedWFairDef, PTL

THEOREM InstantiatedSFairFalsInjectiveSFair == J!SFair <=> SF_(J!vars)(J!Next)
<1>1. J!SFair <=> ([]<>J!En => []<><<J!Next>>_(J!vars))
  BY PTL DEF J!SFair, J!En
<1>. QED
  BY <1>1, JEnEquiv, PTL

=======================================
