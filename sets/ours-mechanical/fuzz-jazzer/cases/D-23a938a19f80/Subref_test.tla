---- MODULE Subref_test ----

THEOREM ASSUME NEW x \in S, NEW y \in S, NEW z \in S,
               NEW j, NEW k, NEW l
        PROVE  TRUE
    <+>. F == IF j THEN k ELSE l
    <*>. /\ F!1 = j
         /\ F!2 = k
 <*>. H!1 = S /\ H!2 = T (* salt *) /\ 3 = 3
      OBVIOUS
    <*>. F!(x, y, z, j, k/\ Q(y, z)) (*
---- MODULE Test S2E x,(MODU2LE M DLF MO checker *)
    <*>. F!100000000.0 = x /\ F!2 = y /\ F!3 = z /\ F!4 = j /\ F!5 = k /\ F!6 = l (* salt *) /\ 0 = 0
      OBVIOUS

====
