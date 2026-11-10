
------- MODULE Github710a ----
EXTENDS TLC, TLCExt
VARIABLE x
vars == <<x>>
<>stOnce == []v(x => [](~x => []~x)) \* G(x => G(~x => G~x))
CONSTANT x
CONSTANT x
op == lbl(x) :: 0
====
  <<2, [x,
<1>. DEFINE G(T) == (\A x \in T : <>[]F(x)) <=> <>[](\A x \in T : F(x))
<1>1. G({})
