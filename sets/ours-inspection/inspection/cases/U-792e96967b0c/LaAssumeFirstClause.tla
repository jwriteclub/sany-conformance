---- MODULE LaAssumeFirstClause ----
\* Aim: TLAplusParser.java:3113, jj_2_14 false: the first clause of an inner ASSUME is none of AssumeProve, NEW, Expression.
\* Theorem's jj_2_22(3) only sees ASSUME P ::, so it commits to AssumeProve; the label takes the inner ASSUME, where ) fails.
\* JaCoCo cannot show this side: the probe follows the throw at jj_consume_token(-1). See grammar-gaps.md.
THEOREM T == ASSUME P :: ASSUME ) PROVE TRUE PROVE TRUE
====
