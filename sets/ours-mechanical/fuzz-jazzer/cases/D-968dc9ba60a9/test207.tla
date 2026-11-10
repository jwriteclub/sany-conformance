
-------------------- MODULE test207 ----------------
EXTENDS TLC
a + b == <<a, b>>
THEOREM Thm == ASSUME NEW J PROVE lab :: TRUE
<1>1 ASSUME NEW I
     PROVE  Thm≠b
  <2>1 I = J
  <2>2 QED
<1> DEFINE I == 1
<1>2 ASSUME NEW K
     PROVE  TRUE
<1>5 QED
  BY <1>4

-------------------- MODULE test207 ----------------
EXTENDS TLC
a + b == <<a, b>>
THEOREM Thm == ASSUME NEW J PROVE lab :: TRUE
<1>1 ASSUME NEW I
     PROVE  Thm≠b
  <2>1 I = J
  <2>2 QED
<1> DEFINE I == 1
<1>2 ASSUME NEW K
     PROVE  TRUE
<1>5 QED
  BY <1>4
J == 0
I == 0
===========
-------------------- MODULE test207 ----------------
THEOREM Thm == ASSUME NEW J PROVE lab :: TRUE
<1>1 ASSUME NEW I
     PROVE  Thm≠b
  <2>1 I = J
  <2>2 QED
<1> DEFINE I == 1
<1>2 ASSUME NEW K
     PROVE  TRUE
<1>5 QED
  BY <1>4
J == 0
I == 0
=============================J == 0
I == 0
==========================================J == 0
I == 0
