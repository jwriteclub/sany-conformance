---- MODULE ApInnerScope ----
\* Aim: Generator.java:3597, the side of illegalAPPosRef where the ASSUME/PROVE has no goal.
\* T!1 selects the inner ASSUME/PROVE. Only an outer one gets a goal (AssumeProveNode.java:77),
\* so ap.getGoal() is null. !2 picks a clause inside the scope of NEW x, so inScopeOfDecl[1] holds
\* and the whole test is true: SANY rejects with 4005 for reaching into a declaration's scope.
THEOREM T == ASSUME ASSUME NEW x, x = 1 PROVE x = 1 PROVE TRUE
A == T!1!2
====
