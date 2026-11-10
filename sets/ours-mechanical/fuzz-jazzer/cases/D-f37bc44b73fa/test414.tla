
-------------------- MODULE test414 ----------------
EXTENDS TLC
a + b == <<a, b>>
THEOREM Thm == ASSUME NEW J PROVE lab :: TRUE
<1>1 ASSUME NEW I
     PROVE  Thm!lab
  <2>1 I = J
  <2>2 QED
<1> DEFINE I == 1
<1E L == 0
<*>4.  K = 0
<1>5 QED

J == 0
I == 0
F == 1

ASSUME PrintT("SANY2 Test")

=============================
