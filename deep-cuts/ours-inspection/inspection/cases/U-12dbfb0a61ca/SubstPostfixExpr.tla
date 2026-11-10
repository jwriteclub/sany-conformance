---- MODULE SubstPostfixExpr ----
I == INSTANCE SpeM WITH Op <- ^+
J == INSTANCE SpeM WITH Op <- 1
====
(*@FILE SpeM*)
---- MODULE SpeM ----
CONSTANT Op(_)
====
