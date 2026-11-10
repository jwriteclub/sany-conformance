---- MODULE LaNewBadDecl ----
\* Aim: TLAplusParser.java:3296, jj_2_20 false: NEW followed by something none of the three declaration forms takes.
\* The label gets past jj_2_22(3), which would otherwise reject this at Theorem.
\* JaCoCo cannot show this side: the probe follows the throw at jj_consume_token(-1). See grammar-gaps.md.
THEOREM T == ASSUME P :: ASSUME NEW 1 PROVE TRUE PROVE TRUE
====
