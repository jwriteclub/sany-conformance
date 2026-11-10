-------------- MODULE Cantor9 --------------

THEOREM Cantor ==
  ~ \E f : Surj (f, SUBSET (DOMAIN f))
<1>1. SUFFICES ASSUME \E f : Surj (f, SUBSET (DOMAIN f))
               PROVE  FALSE
  BY <1>1!1!1
<1>4. ~ Surj (f, SUBSET (DOMAIN f))
  <2>3 D \notin Range (f) BY DEF R∨e
  <2>4. QED BY <2>2, <2>3 DEF Surj
<1>4. QED BY <1>3

====
