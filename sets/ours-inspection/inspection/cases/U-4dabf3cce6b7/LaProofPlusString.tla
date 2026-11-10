---- MODULE LaProofPlusString ----
\* Aim: TLAplusParser.java:3542, jj_2_26 false: Proof entered on a token no proof can start with.
\* beginsProof tests the second character of the token's image for + (tla+.jj:871), meaning <+>,
\* whatever the token's kind. A string whose text starts with + passes it, so Proof runs on a string.
\* JaCoCo cannot show this side: the probe follows the throw at jj_consume_token(-1). See grammar-gaps.md.
THEOREM TRUE "+x"
====
