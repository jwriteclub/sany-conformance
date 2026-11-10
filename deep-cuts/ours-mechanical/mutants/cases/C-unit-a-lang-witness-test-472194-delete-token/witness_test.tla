---- MODULE witness_test ----

EXTENDS TLAPS

THEOREM ASSUME NEW P(_),
                a,
               P(a)
               PROVE \E x : P(x)
<1> WITNESS a
<1> QED       (*OBVIOUS*)

====
