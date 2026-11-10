---- MODULE ThmNonLeib ----
\* Aim: ThmOrAssumpDefNode.java:495, no gain: J!TI gets an APSubstInNode body (Generator.java:4981), which computes no nonLeibnizParams.
---- MODULE In ----
CONSTANT K
VARIABLE v
THEOREM TI == ENABLED (v' = K)
====
VARIABLE x
J(q) == INSTANCE In WITH K <- q, v <- x
====
