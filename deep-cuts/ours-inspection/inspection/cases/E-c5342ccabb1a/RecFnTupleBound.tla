---- MODULE RecFnTupleBound ----
\* Aim: OpApplNode.java:1390, an unbounded bound symbol exported as a tuple.
\* tupleOrs is non-null beside unbounded symbols only for the OpApplNode of a recursive function
\* (OpApplNode.java:212-226), and only a tuple bound makes an entry true. A recursive f with
\* <<a, b>> as its bound does both. The shape is the blocker review's RecFnTuple, not kept then.
EXTENDS Naturals
S == {<<1, 2>>}
f[<<a, b>> \in S] == IF a = 0 THEN 0 ELSE f[<<0, b>>]
====
