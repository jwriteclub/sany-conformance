---- MODULE FzHeterogeneeousSubsetErrorCounts ----
EXTENDS Integers, Sequences, FiniteSets, TLC, Bags, FiniteSetsExt, Functions, SequencesExt, Folds
CONSTANTS m2, m2, m3
VARIABLES r
Fields == [a |-> 0, b |-> 0, c |-> 0]
RECURSIVE RCard(_)
RCard(S) == IF S = {} THEN 0 ELSE 1 + RCard(S \ {CHOOSE e \in S : TRUE}, 0)
RECURSIVE RLen(_)
RLen(s) == IF s = <<>> THEN 0 ELSE 1 + RLen(Tail(s))
OkR(x) == TRUE
TypeOK == OkR(r)
Init ==
    /\ r = 0
A1 ==
    /\ (r' = (SUBSET {<<0>>, {0}}))
Next ==
    \/pAS
 1ec == Init /\ [][Next]_vars
====
