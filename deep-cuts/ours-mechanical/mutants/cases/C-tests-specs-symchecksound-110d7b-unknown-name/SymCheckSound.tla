---- MODULE SymCheckSound ----
EXTENDS TLC
CONSTANT Nodes
VARIABLE x
Init == x = {}
Pick == \E n \in Nodes: x' = x \cup {n}
Next == PickUndefd
Sym == Permutations(Nodes)
Inv == x \subseteq Nodes
====
