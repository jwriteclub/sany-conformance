---- MODULE M ----
CONSTANT g(_)
K(f(_)) == INSTANCE M2
THEOREM ThmU == 1
THEOREM ThmG == g(1)
THEOREM ThmN == g([]TRUE)
THEOREM ThmR == K(g)!Thm2
ASSUME AsmG == g(1)
THEOREM ThmN == g([]TRUE)mR == K(g)!Thm2
ASSUME Ass == g(1) /\ []TRUE
====
