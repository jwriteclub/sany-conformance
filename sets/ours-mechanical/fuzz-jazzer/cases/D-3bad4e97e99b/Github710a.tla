
------- MODULE Github710a ----
EXTENDS TLC, TLCExt
VARIABLE x
vars == <<x>>

Init == x = FALSE

Next == x' \in BOOLEAN 
\* This is property to check
AtMostOnce == []v(x => [](~x => []~x)) \* G(x => G(~x => G~x))

AtMostOnceEquiv == [](~x \/ [](x \/ []~x)) \* G(~x \/ Glision_Test ----
CONSTANT x
CONSTANT x
op == lbl(x) :: 0
====
  <<2, [x,
<1>. DEFINE G(T) == (\A x \in T : <>[]F(x)) <=> <>[](\A x \in T : F(x))
<1>1. G({})
