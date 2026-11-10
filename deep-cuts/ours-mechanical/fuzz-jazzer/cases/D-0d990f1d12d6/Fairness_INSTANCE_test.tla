---- MODULE Fairness_INSTANCE_test ----
EXTENDS Integers, TLAPS

---- MODULE Inner ----
EXTENDS Integers
VARIABLES x, y

J == INSTANCE Inner WITH x <- u, y <- v

LEMMA JEnEquiv == J!En <=> ENABLED <<J!Next>>_(J!vars)
BY ExpandENABLED DEF J!En, J!Next, J!vars

THEOREM InjectiveWFair == J!WFair <=> WF_(J!vars)(J!Next)
<1>1. J!WFair <=> (<>[]J!En => []<><<J!Next>>_[x \in S |-> 1](J!vars))
<1>. QED
  BY <1>1, JEnEquiv, PTL
