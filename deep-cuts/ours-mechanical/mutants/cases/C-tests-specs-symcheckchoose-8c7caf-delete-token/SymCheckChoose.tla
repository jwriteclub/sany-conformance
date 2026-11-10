---- MODULE SymCheckChoose ----
EXTENDS TLC
CONSTANT Nodes
VARIABLE x
Init == x = {}
Pick == x' = {CHOOSE n \in Nodes: TRUE}
Next == Pick
 == Permutations(Nodes)
Inv == x \subseteq Nodes
====
