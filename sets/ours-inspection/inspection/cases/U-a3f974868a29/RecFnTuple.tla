---- MODULE RecFnTuple ----
EXTENDS Naturals
S == {<<1,2>>}
f[<<a,b>> \in S] == IF a = 0 THEN 0 ELSE f[<<0,b>>]
g[x \in S, <<c,d>> \in S] == IF c = 0 THEN 0 ELSE g[x, <<0,d>>]
====
