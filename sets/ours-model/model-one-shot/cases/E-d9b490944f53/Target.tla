---- MODULE Target ----
VARIABLE v
OpBad(F(_)) == F(v') /\ (v' ~> v')
BadUse1 == OpBad(')
OpGood(F(_)) == F(v')
OpArg(x) == x' ~> x'
BadUse2 == OpGood(OpArg)
====
