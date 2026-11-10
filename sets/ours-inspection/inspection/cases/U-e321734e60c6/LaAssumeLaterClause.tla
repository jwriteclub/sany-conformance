---- MODULE LaAssumeLaterClause ----
\* Aim: TLAplusParser.java:3156, jj_2_15 false: a clause after the comma is none of AssumeProve, NEW, Expression.
\* jj_2_22(3) sees ASSUME f , and commits.
\* JaCoCo cannot show this side: the probe follows the throw at jj_consume_token(-1). See grammar-gaps.md.
THEOREM T == ASSUME f, ) PROVE TRUE
====
